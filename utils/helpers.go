package utils

import (
	"../config"
	"../database"
	"../interfaces"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

func ExtractContext(r *http.Request) *interfaces.Claims {
	claims, ok := r.Context().Value(config.USER_DATA_CONTEXT_ADDRESS).(interfaces.Claims)
	if ok {
		return &claims
	}
	return nil
}

func ExtractToken(r *http.Request) string {
	var rawToken string

	if len(r.Header["Authorization"]) > 0 {
		rawToken = r.Header["Authorization"][0]
	}

	return rawToken
}

func CheckValidToken(token string) bool {
	isExist := database.RedisConn.SIsMember(config.INVALID_TOKENS, token).Val()
	return !isExist
}

func GenerateToken(user *interfaces.UserInfo) string {
	expireToken := time.Now().Add(time.Hour * 1).Unix()
	claims := interfaces.Claims{
		user.Id,
		user.UserName,
		user.Email,
		jwt.StandardClaims{
			ExpiresAt: expireToken,
			Issuer:    "localhost:8080",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, _ := token.SignedString([]byte(config.APP_SECRET))

	return signedToken
}

func ExtractUserData(rawToken string) *interfaces.Claims {
	token, err := jwt.ParseWithClaims(rawToken, &interfaces.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(config.APP_SECRET), nil
	})

	if err != nil {
		return nil
	}

	claims, ok := token.Claims.(*interfaces.Claims)

	if ok && token.Valid {
		return claims
	}

	return nil
}
