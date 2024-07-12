package gateway

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type HTTPServerConfig struct {
	TLSConfig         *tls.Config
	ReadTimeout       time.Duration
	ReadHeaderTimeout time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	MaxHeaderBytes    int
	TLSNextProto      map[string]func(*http.Server, *tls.Conn, http.Handler)
	ConnState         func(net.Conn, http.ConnState)
}

func (c *HTTPServerConfig) ApplyTo(s *http.Server) {
	s.TLSConfig = c.TLSConfig
	s.ReadTimeout = c.ReadTimeout
	s.ReadHeaderTimeout = c.ReadHeaderTimeout
	s.WriteTimeout = c.WriteTimeout
	s.IdleTimeout = c.IdleTimeout
	s.MaxHeaderBytes = c.MaxHeaderBytes
	s.TLSNextProto = c.TLSNextProto
	s.ConnState = c.ConnState
}

// DefaultMarshaler return default grpc-gateway marshaler with additional support for emit empty field
func DefaultMarshaler() runtime.ServeMuxOption {
	return runtime.WithMarshalerOption(runtime.MIMEWildcard,
		&runtime.JSONPb{
			MarshalOptions:   protojson.MarshalOptions{UseEnumNumbers: true, EmitUnpopulated: true, UseProtoNames: true},
			UnmarshalOptions: protojson.UnmarshalOptions{DiscardUnknown: true},
		})
}

// ProtoJSONMarshaler return the marshaler option with support serialization data with json_name specific
func ProtoJSONMarshaler() runtime.ServeMuxOption {
	return runtime.WithMarshalerOption(runtime.MIMEWildcard,
		&runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{EmitUnpopulated: true},
		})
}

// Redirect support redirect endpoint for GRPC gateway by using server metadata
// Example:
//
//	redirect := metadata.Pairs(
//		"Redirect", "https://example.com",
//		"Redirect-Code", "308",
//	)
//	grpc.SendHeader(ctx, redirect)
//
// This will redirect endpoint to https://example.com with code 308, by default redirect code is 301.
func Redirect() runtime.ServeMuxOption {
	fn := func(ctx context.Context, w http.ResponseWriter, _ proto.Message) error {
		md, ok := runtime.ServerMetadataFromContext(ctx)
		if !ok {
			return nil
		}

		redirect := md.HeaderMD.Get("redirect")
		redirectCode := md.HeaderMD.Get("redirect-code")
		if len(redirect) > 0 {
			w.Header().Set("Location", redirect[0])
			if len(redirectCode) > 0 {
				code, err := strconv.Atoi(redirectCode[0])
				if err != nil {
					return err
				}
				if code > 300 && code < 400 {
					w.WriteHeader(code)
				}
			}

			w.WriteHeader(301)
		}
		return nil
	}
	return runtime.WithForwardResponseOption(fn)
}

// FormURLEncodedMarshaler is custom Marshaler that supports reading x-www-form-urlencoded mine type
func FormURLEncodedMarshaler() runtime.ServeMuxOption {
	return runtime.WithMarshalerOption("application/x-www-form-urlencoded", &formMarshaler{
		&runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{EmitUnpopulated: true},
		},
	})
}

type formMarshaler struct {
	*runtime.JSONPb
}

func (j *formMarshaler) NewDecoder(r io.Reader) runtime.Decoder {
	return runtime.DecoderFunc(func(v interface{}) error { return formDecoderFunc(r, v) })
}

func formDecoderFunc(d io.Reader, v interface{}) error {
	msg, ok := v.(proto.Message)
	if !ok {
		return fmt.Errorf("not proto message")
	}

	formData, err := ioutil.ReadAll(d)
	if err != nil {
		return err
	}

	values, err := url.ParseQuery(string(formData))
	if err != nil {
		return err
	}

	filter := &utilities.DoubleArray{}
	return runtime.PopulateQueryParameters(msg, values, filter)
}
