package controllers

import (
	"../interfaces"
	Tasks "../models/tasks"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func AllTaskController(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		tasks := Tasks.GetAll()
		result, err := json.Marshal(tasks)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
		return
	} else if r.Method == "POST" {

		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var task interfaces.Task
		json.Unmarshal(body, &task)

		if &task == nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := Tasks.Add(&task)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "task "+strconv.FormatInt(id, 10)+" created!")
		return
	}
}

func UpdateTaskController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if r.Method == "PUT" {
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var task interfaces.Task
		json.Unmarshal(body, &task)

		fmt.Println(task)

		if &task == nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		task.Id = id

		err = Tasks.Update(&task)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "task "+vars["id"]+" updated!")
		return
	}

	if r.Method == "DELETE" {
		err := Tasks.Delete(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "task "+vars["id"]+" removed!")
		return
	}

}
