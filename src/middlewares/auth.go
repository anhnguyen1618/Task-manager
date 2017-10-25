package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/anhnguyen300795/Task-manager/src/config"
	"github.com/anhnguyen300795/Task-manager/src/utils"
)

func (env *MiddleWares) Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		rawToken := utils.ExtractToken(req)

		if rawToken != "" && !utils.CheckValidToken(env.RedisDB, rawToken) {
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
