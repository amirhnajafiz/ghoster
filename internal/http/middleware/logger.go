package middleware

import (
	"log"
	"net/http"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != "/healthz" {
			log.Println(r.RequestURI)
		}

		next.ServeHTTP(w, r)
	})
}
