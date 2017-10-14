package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../config"
	"../interfaces"
	"../models"
	"../utils"
)

func LoginController(env *interfaces.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {
			body, err := ioutil.ReadAll(r.Body)

			utils.CheckErrors(w, err, http.StatusInternalServerError)

			var user interfaces.UserInfo
			json.Unmarshal(body, &user)

			Users := models.Users{env.DB}

			realUser := Users.CheckCredential(&user)

			if realUser != nil {
				token := utils.GenerateToken(realUser)

				w.Header().Set("Authorization", token)
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, "Login successfully!")
				return
			}

			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Login failed!")

		}
	}
}

func SignUpController(env *interfaces.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			body, err := ioutil.ReadAll(r.Body)

			utils.CheckErrors(w, err, http.StatusInternalServerError)

			var user interfaces.UserInfo
			json.Unmarshal(body, &user)

			Users := models.Users{env.DB}

			status, err := Users.AddOne(&user)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}

			fmt.Fprintf(w, status)
		}
	}
}

func SignOutController(env *interfaces.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var rawToken string
		if len(req.Header["Authorization"]) > 0 {
			rawToken = req.Header["Authorization"][0]
		}

		env.RedisDB.SAdd(config.INVALID_TOKENS, rawToken)

		fmt.Println(env.RedisDB.SMembers(config.INVALID_TOKENS))
	}
}
