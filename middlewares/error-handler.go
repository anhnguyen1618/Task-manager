package middlewares

import (
	"fmt"
	"net/http"
)

func MuxErrorHandler(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()

		next.ServeHTTP(w, r)
	}
}
