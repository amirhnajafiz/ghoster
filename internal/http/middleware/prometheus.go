package middleware

import (
	"net/http"

	"github.com/amirhnajafiz/ghoster/internal/metrics"
)

func Metrics(metrics metrics.Metrics) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.RequestURI != "/healthz" {
				metrics.AddRequest(r.RequestURI, r.Method)
			}

			next.ServeHTTP(w, r)
		})
	}
}
