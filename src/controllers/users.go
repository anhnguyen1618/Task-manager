package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/anhnguyen300795/Task-manager/src/models"
	"github.com/anhnguyen300795/Task-manager/src/utils"
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
