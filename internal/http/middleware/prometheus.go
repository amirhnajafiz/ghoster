package middleware

import (
	"net/http"

	"github.com/amirhnajafiz/ghoster/internal/metrics"

	"github.com/prometheus/client_golang/prometheus"
)

func Metrics(metrics metrics.Metrics) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.RequestURI != "/healthz" {
				metrics.Requests.With(prometheus.Labels{"endpoint": r.RequestURI, "method": r.Method}).Add(1)
			}

			next.ServeHTTP(w, r)
		})
	}
}
