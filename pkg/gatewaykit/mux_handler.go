package gatewaykit

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/status"
)

const (
	RedirectKey = "url_redirect"
)

type httpGetHandlerFn func(r *http.Request, pathParams map[string]string) (interface{}, error)
type httpHandlerFn func(r *http.Request) (interface{}, error)

/*
How to use:
	func (s *service) RegisterWithMuxServer(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
		mux.HandlePath(http.MethodGet, "/path", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
			gatewaykit.WrapHttpGetHandler(w, r, pathParams, false, httpGetHandlerFn)
		})

		pb.RegisterGrpcHandler(ctx, mux, conn)
		return nil
	}
*/

func encodeJson(w http.ResponseWriter, data []byte, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}

func encodeHtml(w http.ResponseWriter, data []byte, statusCode int) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(statusCode)
	w.Write(data)
}

func encodeStream(w http.ResponseWriter, data []byte, statusCode int) {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.WriteHeader(statusCode)
	w.Write(data)
}

func WrapHttpGetHandler(w http.ResponseWriter, r *http.Request, pathParams map[string]string, allowRedirect bool, h httpGetHandlerFn) {
	resp, err := h(r, pathParams)
	httpHandler(w, r, resp, allowRedirect, err)
}

func WrapHttpHandler(w http.ResponseWriter, r *http.Request, allowRedirect bool, h httpHandlerFn) {
	resp, err := h(r)
	httpHandler(w, r, resp, allowRedirect, err)
}

func WrapHtmlGetHandler(w http.ResponseWriter, r *http.Request, pathParams map[string]string, allowRedirect bool, h httpGetHandlerFn) {
	resp, err := h(r, pathParams)
	if err != nil {
		errHandler(w, err)
		return
	}

	encodeHtml(w, []byte(resp.(string)), http.StatusOK)
}

func WrapHtmlHandler(w http.ResponseWriter, r *http.Request, allowRedirect bool, h httpHandlerFn) {
	resp, err := h(r)
	if err != nil {
		errHandler(w, err)
		return
	}

	encodeHtml(w, []byte(resp.(string)), http.StatusOK)
}

func httpHandler(w http.ResponseWriter, r *http.Request, resp interface{}, allowRedirect bool, err error) {
	if err != nil {
		errHandler(w, err)
		return
	}

	switch v := resp.(type) {
	case string:
		// redirect if response is string
		if allowRedirect {
			http.Redirect(w, r, v, http.StatusFound)
		} else {
			encodeJson(w, []byte(v), http.StatusOK)
		}
		return

	case map[string]string:
		if urlRedirect, ok := v[RedirectKey]; ok && allowRedirect {
			for k, vl := range v {
				if k != RedirectKey {
					w.Header().Add(k, vl)
				}
			}

			http.Redirect(w, r, urlRedirect, http.StatusTemporaryRedirect)
			return
		}

	case []byte:
		encodeStream(w, v, http.StatusOK)
		return
	}

	b, err := json.Marshal(resp)
	if err != nil {
		errHandler(w, errors.New("can not serialize data"))
		return
	}

	encodeJson(w, b, http.StatusOK)
}

func errHandler(w http.ResponseWriter, err error) {
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

	b := bytes.Buffer{}
	b.WriteString("{\"message\": \"")
	b.WriteString(err.Error())
	b.WriteString("\", \"code\": ")
	b.WriteString(strconv.Itoa(int(s.Code())))
	b.WriteString("}")

	encodeJson(w, b.Bytes(), httpStatus)
}
