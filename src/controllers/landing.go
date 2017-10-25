package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func (controller *Controllers) LandingController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		data, err := ioutil.ReadFile("public/index.html")
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, string(data))
	}

}
