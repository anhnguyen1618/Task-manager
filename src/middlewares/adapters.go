package middlewares

import (
	"net/http"
)

type Adapter func(http.HandlerFunc) http.HandlerFunc

func ApplyMiddleware(middleWares ...Adapter) Adapter {
	return func(rawHandler http.HandlerFunc) http.HandlerFunc {
		// Convert rawHandler from type http.HandlerFunc to type http.Handler
		convertedHandler := rawHandler

		/*
		* Iterate through list of middleware [a,b,c]
		* Compose middlewares functions to one new function: convertedHandler = (rw, r) => a(b(c))(rw, r)
		 */
		for i := len(middleWares) - 1; i >= 0; i-- {
			convertedHandler = middleWares[i](convertedHandler)
		}
		return convertedHandler

	}

}
