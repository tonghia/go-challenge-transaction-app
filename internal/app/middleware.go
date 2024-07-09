package app

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type middleware struct {
	log *zap.SugaredLogger
}

func NewCustomMiddleware(
	log *zap.SugaredLogger,
) *middleware {
	return &middleware{
		log: log,
	}
}

func (m *middleware) MiddlewareHandleFunc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		// r.Header.Set(HeaderUserID, "")
		// r.Header.Set(HeaderUserPhone, "")

		if strings.HasPrefix(path, "/user") {
			header := r.Header

			authorizeHeader := header.Get("Authorization")

			if authorizeHeader == "" {
				m.log.Warn("Empty Authorization header")
				authenError(w, r)
				return
			}

			tokenStrings := strings.Split(authorizeHeader, "Bearer ")
			if len(tokenStrings) != 2 {
				m.log.Warn("Invalid Authorization header", zap.Any("header", header))
				authenError(w, r)
				return
			}

			// userInfo, err := m.authenUserRequest(tokenStrings[1])
			// if err != nil {
			// 	m.log.Error("Invalid authen user", l.Error(err), l.String("path", path), l.String("token", tokenStrings[1]), l.Any("header", header))
			// 	authenError(w, r)
			// 	return
			// }

			// if !userInfo.GetActive() {
			// 	authenError(w, r)
			// 	return
			// }

			// r.Header.Set(gwutils.ToMetadataHeader(gwutils.UserIDKey), userInfo.GetSub())
			// r.Header.Set(gwutils.ToMetadataHeader(gwutils.UserPhoneKey), userInfo.GetPhone())
		}

		next.ServeHTTP(w, r)
	})
}

type errorBody struct {
	Code    uint32 `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func customHTTPError(w http.ResponseWriter, r *http.Request, err error) {
	const fallback = `{"message": "failed to marshal error message"}`

	w.Header().Set("Content-type", "application/json")
	e := errorBody{
		Code:    uint32(status.Code(err)),
		Message: status.Convert(err).Message(),
	}

	s := status.Convert(err)
	st := runtime.HTTPStatusFromCode(s.Code())
	if st == http.StatusInternalServerError {
		e.Message = "Something went wrong. Please try again later."
	}
	w.WriteHeader(st)
	jErr := json.NewEncoder(w).Encode(e)

	if jErr != nil {
		w.Write([]byte(fallback))
	}
}

func authenError(w http.ResponseWriter, r *http.Request) {
	customHTTPError(w, r, status.Error(codes.PermissionDenied, "Thông tin xác thực không hợp lệ."))
}
