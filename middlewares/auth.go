package middlewares

import (
	"../config"
	"../utils"
	"context"
	"net/http"
)

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		rawToken := utils.ExtractToken(req)

		if rawToken != "" && !utils.CheckValidToken(rawToken) {
			http.NotFound(res, req)
			return
		}

		claims := utils.ExtractUserData(rawToken)

		if claims == nil {
			http.NotFound(res, req)
			return
		}

		ctx := context.WithValue(req.Context(), config.USER_DATA_CONTEXT_ADDRESS, claims)
		next(res, req.WithContext(ctx))
		return

	}
}
