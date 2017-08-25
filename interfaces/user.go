package interfaces

import (
	"github.com/dgrijalva/jwt-go"
)

type UserInfo struct {
	Id       int
	UserName string `json:"userName"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Claims struct {
	Id       int    `json:"id"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	jwt.StandardClaims
}
