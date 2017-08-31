package controllers

import (
	"../utils"
	"fmt"
	"io/ioutil"
	"net/http"
)

func LandingController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		claims := utils.ExtractContext(r)
		fmt.Println(claims)
		data, err := ioutil.ReadFile("public/index.html")
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, string(data))
	}
}