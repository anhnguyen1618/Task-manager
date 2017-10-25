package utils

import (
	"net/http"
)

func CheckErrors(w http.ResponseWriter, err error, status int) {
	if err != nil {
		http.Error(w, err.Error(), status)
		panic(err)
	}
}
