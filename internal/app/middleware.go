package app

import (
	"net/http"

	"go.uber.org/zap"
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

		// TODO: authentication here

		next.ServeHTTP(w, r)
	})
}
