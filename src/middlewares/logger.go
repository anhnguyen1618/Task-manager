package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func (env *MiddleWares) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		defer func() {
			end := time.Now()
			duration := end.Sub(start)
			fmt.Printf("[%s] %q %v\n", r.Method, r.URL.String(), duration)
		}()

		next.ServeHTTP(w, r)
	}
}
