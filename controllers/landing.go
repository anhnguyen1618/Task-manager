package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"../interfaces"
	"../utils"
)

func LandingController(env *interfaces.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

}
