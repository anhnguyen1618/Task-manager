package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/anhnguyen300795/Task-manager/interfaces"
	"github.com/anhnguyen300795/Task-manager/models"
	"github.com/anhnguyen300795/Task-manager/utils"
	"github.com/gorilla/mux"
)

func (controller *Controllers) AllTaskController(w http.ResponseWriter, r *http.Request) {
	Tasks := &models.Tasks{controller.DB}
	if r.Method == "GET" {
		tasks := Tasks.GetAll()
		result, err := json.Marshal(tasks)
		utils.CheckErrors(w, err, http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
		return

	} else if r.Method == "POST" {

		body, err := ioutil.ReadAll(r.Body)

		utils.CheckErrors(w, err, http.StatusInternalServerError)

		var task interfaces.Task
		json.Unmarshal(body, &task)

		if &task == nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := Tasks.Add(&task)

		utils.CheckErrors(w, err, http.StatusInternalServerError)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "task "+strconv.FormatInt(id, 10)+" created!")
		return
	}
}

func (controller *Controllers) UpdateTaskController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	utils.CheckErrors(w, err, http.StatusBadRequest)

	Tasks := &models.Tasks{controller.DB}

	if r.Method == "PUT" {
		body, err := ioutil.ReadAll(r.Body)

		utils.CheckErrors(w, err, http.StatusInternalServerError)

		var task interfaces.Task
		json.Unmarshal(body, &task)

		if &task == nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		task.Id = id

		err = Tasks.Update(&task)

		utils.CheckErrors(w, err, http.StatusInternalServerError)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "task "+vars["id"]+" updated!")
		return
	}

	if r.Method == "DELETE" {
		user := utils.ExtractContext(r)
		task := Tasks.GetOne(id)

		if task == nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "task "+vars["id"]+" not found!")
			return
		}

		if user.Role != "ADMIN" && user.UserName != task.Assignor {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "You do not have permission for this action!")
			return
		}

		err := Tasks.Delete(id)
		utils.CheckErrors(w, err, http.StatusInternalServerError)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "task "+vars["id"]+" removed!")
		return
	}

}
