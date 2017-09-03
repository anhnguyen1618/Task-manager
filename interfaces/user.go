package interfaces

import (
	"github.com/dgrijalva/jwt-go"
)

type UserInfo struct {
	Id       int
	UserName string `json:"userName"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type Claims struct {
	Id       int    `json:"id"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	jwt.StandardClaims
}
