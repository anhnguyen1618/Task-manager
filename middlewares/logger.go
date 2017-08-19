package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		defer func() {
			end := time.Now()
			duration := end.Sub(start)
			fmt.Printf("[%s] %q %v\n", r.Method, r.URL.String(), duration)
		}()

		next.ServeHTTP(w, r)
	})
}
