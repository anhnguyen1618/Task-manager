package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/anhnguyen300795/Task-manager/src/interfaces"

	"github.com/anhnguyen300795/Task-manager/src/models"
	"github.com/anhnguyen300795/Task-manager/src/utils"
	"github.com/gorilla/mux"
)

func (controller *Controllers) UsersController(w http.ResponseWriter, r *http.Request) {

	Users := &models.Users{controller.DB}

	if r.Method == "GET" {
		users := Users.GetAll()
		result, err := json.Marshal(users)
		utils.CheckErrors(w, err, http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
		return
	}
}

func (controller *Controllers) UpdateUserController(w http.ResponseWriter, r *http.Request) {

	Users := &models.Users{controller.DB}

	vars := mux.Vars(r)
	userName := vars["userName"]

	if r.Method == "PUT" {
		body, err := ioutil.ReadAll(r.Body)
		utils.CheckErrors(w, err, http.StatusInternalServerError)

		var user interfaces.UserInfo
		json.Unmarshal(body, &user)

		user.UserName = userName
		updatedUser := Users.UpdateOne(&user)

		if updatedUser == nil {
			http.Error(w, err.Error(), http.StatusNotFound)
		}

		resPayload, _ := json.Marshal(updatedUser)

		w.Header().Set("Content-Type", "application/json")
		w.Write(resPayload)
		return
	} else if r.Method == "DELETE" {
		err := Users.DeleteOne(userName)
		utils.CheckErrors(w, err, http.StatusInternalServerError)
	}

}
