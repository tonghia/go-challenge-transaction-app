package gateway

import (
	"net/http"
	"net/http/pprof"

	_ "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type HTTPServerHandler func(*http.ServeMux)

func PrometheusHandler(httpMux *http.ServeMux) {
	// Register prometheus handlers
	httpMux.Handle("/metrics", promhttp.Handler())
}

func PprofHandler(httpMux *http.ServeMux) {
	// Register pprof handlers
	httpMux.HandleFunc("/debug/pprof/", pprof.Index)
	httpMux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	httpMux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	httpMux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	httpMux.HandleFunc("/debug/pprof/trace", pprof.Trace)
}
