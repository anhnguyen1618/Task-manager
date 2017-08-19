package controllers

import (
	Users "../models/users"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func LoginController(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var user Users.UserInfo
		json.Unmarshal(body, &user)

		realUser := Users.CheckCredential(&user)

		if realUser != nil {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Login successfully!")
			return
		}

		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintf(w, "Login failed!")

	}
}

func SignUpController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var user Users.UserInfo
		json.Unmarshal(body, &user)

		status, err := Users.AddOne(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		fmt.Fprintf(w, status)
	}
}
