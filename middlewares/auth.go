package middlewares

import (
	"../config"
	"../utils"
	"context"
	"fmt"
	"net/http"
)

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		rawToken := utils.ExtractToken(req)

		if rawToken != "" && !utils.CheckValidToken(rawToken) {
			res.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(res, "Unauthorized access!")
			return
		}

		claims := utils.ExtractUserData(rawToken)

		if claims == nil {
			res.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(res, "Unauthorized access!")
			return
		}

		ctx := context.WithValue(req.Context(), config.USER_DATA_CONTEXT_ADDRESS, claims)
		next(res, req.WithContext(ctx))
		return

	}
}
