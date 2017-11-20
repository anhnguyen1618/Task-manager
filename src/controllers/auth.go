package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/anhnguyen300795/Task-manager/src/config"
	"github.com/anhnguyen300795/Task-manager/src/interfaces"
	"github.com/anhnguyen300795/Task-manager/src/models"
	"github.com/anhnguyen300795/Task-manager/src/utils"
)

func (controller *Controllers) LoginController(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)

		utils.CheckErrors(w, err, http.StatusInternalServerError)

		var user interfaces.UserInfo
		json.Unmarshal(body, &user)

		Users := models.Users{controller.DB}

		realUser := Users.CheckCredential(&user)

		if realUser != nil {
			token := utils.GenerateToken(realUser)

			result, err := json.Marshal(realUser)
			utils.CheckErrors(w, err, http.StatusInternalServerError)

			w.Header().Set("Authorization", token)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(result)
			return
		}

		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Login failed!")

	}
}

func (controller *Controllers) SignUpController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)

		utils.CheckErrors(w, err, http.StatusInternalServerError)

		var user interfaces.UserInfo
		json.Unmarshal(body, &user)

		Users := models.Users{controller.DB}

		createdUser := Users.AddOne(&user)

		if createdUser == nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		resPayload, err := json.Marshal(createdUser)
		utils.CheckErrors(w, err, http.StatusInternalServerError)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resPayload)
	}
}

func (controller *Controllers) SignOutController(w http.ResponseWriter, r *http.Request) {

	var rawToken string
	if len(r.Header["Authorization"]) > 0 {
		rawToken = r.Header["Authorization"][0]
	}
	controller.RedisDB.SAdd(config.INVALID_TOKENS, rawToken)
}

func (controller *Controllers) CurrentUserController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		user := utils.ExtractContext(r)
		filteredData := interfaces.UserInfo{user.UserName, "", user.Email, user.Role}
		result, err := json.Marshal(filteredData)
		utils.CheckErrors(w, err, http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
	}
}
