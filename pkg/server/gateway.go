package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/tonghia/go-challenge-transaction-app/pkg/server/gateway"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

const (
	maxMsgSize = 50 * 1024 * 1024
)

// gatewayServer wraps gRPC gateway server setup process.
type gatewayServer struct {
	server *http.Server
	config *gatewayConfig
}

type gatewayConfig struct {
	addr              Listen
	muxOptions        []runtime.ServeMuxOption
	serverConfig      *gateway.HTTPServerConfig
	serverMiddlewares []gateway.HTTPServerMiddleware
	serverHandlers    []gateway.HTTPServerHandler
}

func createDefaultGatewayConfig() *gatewayConfig {
	config := &gatewayConfig{
		addr: Listen{
			Host: "0.0.0.0",
			Port: 9000,
		},
		muxOptions: []runtime.ServeMuxOption{
			gateway.DefaultMarshaler(),
			runtime.WithErrorHandler(func(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, writer http.ResponseWriter, req *http.Request, err error) {
				//creating a new HTTTPStatusError with a custom status, and passing error
				s := status.Convert(err)
				httpStatus := runtime.HTTPStatusFromCode(s.Code())
				if httpStatus == http.StatusInternalServerError {
					httpStatus = http.StatusBadRequest
				}

				if hs, ok := err.(interface {
					HttpStatus() int
				}); ok {
					httpStatus = hs.HttpStatus()
				}
				newError := runtime.HTTPStatusError{
					HTTPStatus: httpStatus,
					Err:        err,
				}
				// using default handler to do the rest of heavy lifting of marshaling error and adding headers
				runtime.DefaultHTTPErrorHandler(ctx, mux, marshaler, writer, req, &newError)
			}),
		},
		serverHandlers: []gateway.HTTPServerHandler{
			gateway.PrometheusHandler,
			gateway.PprofHandler,
		},
	}

	return config
}

func newGatewayServer(c *gatewayConfig, grpcAddr string, servers []ServiceServer) (*gatewayServer, error) {
	var unaryIntercept []grpc.UnaryClientInterceptor
	var streamIntercept []grpc.StreamClientInterceptor

	conn, err := grpc.NewClient(
		grpcAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(maxMsgSize)),
		grpc.WithChainUnaryInterceptor(unaryIntercept...),
		grpc.WithChainStreamInterceptor(streamIntercept...),
	)
	if err != nil {
		return nil, fmt.Errorf("fail to dial gRPC server. %w", err)
	}

	// init mux
	mux := runtime.NewServeMux(c.muxOptions...)
	// allowCors := cors.New(cors.Options{
	// 	AllowedOrigins: []string{"*"},
	// 	AllowedMethods: []string{
	// 		http.MethodHead,
	// 		http.MethodGet,
	// 		http.MethodPost,
	// 		http.MethodPut,
	// 		http.MethodPatch,
	// 		http.MethodDelete,
	// 		http.MethodOptions,
	// 	},
	// 	AllowedHeaders:   []string{"*"},
	// 	AllowCredentials: false,
	// })

	for _, svr := range servers {
		err := svr.RegisterWithMuxServer(context.Background(), mux, conn)
		if err != nil {
			return nil, fmt.Errorf("failed to register handler. %w", err)
		}
	}

	var handler http.Handler = mux
	for i := len(c.serverMiddlewares) - 1; i >= 0; i-- {
		handler = c.serverMiddlewares[i](handler)
	}

	httpMux := http.NewServeMux()
	for _, h := range c.serverHandlers {
		h(httpMux)
	}

	httpMux.Handle("/", handler)

	svr := &http.Server{
		Addr:    c.addr.String(),
		Handler: httpMux,
	}
	// svr.Handler = allowCors.Handler(httpMux)

	if cfg := c.serverConfig; cfg != nil {
		cfg.ApplyTo(svr)
	}

	return &gatewayServer{
		server: svr,
		config: c,
	}, nil
}

func (s *gatewayServer) serve() error {
	log.Println("http server starting at", s.config.addr.String())
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Println("error starting http server", err)
		return err
	}

	return nil
}

func (s *gatewayServer) shutdown(ctx context.Context) {
	err := s.server.Shutdown(ctx)
	log.Println("all http(s) requests finished")
	if err != nil {
		log.Println("failed to shutdown grpc-gateway server", err)
	}
}
