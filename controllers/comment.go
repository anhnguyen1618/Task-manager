package controllers

import (
	"../interfaces"
	Comments "../models/comments"
	"../utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func CommentController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["id"])

	utils.CheckErrors(w, err, http.StatusBadRequest)

	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)

		utils.CheckErrors(w, err, http.StatusInternalServerError)

		var comment interfaces.Comment
		json.Unmarshal(body, &comment)

		commentID, err := Comments.Add(&comment, taskId)

		utils.CheckErrors(w, err, http.StatusInternalServerError)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Comment "+strconv.FormatInt(commentID, 10)+" added!")
		return
	}

}

func UpdateCommentController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	commentId, err := strconv.Atoi(vars["commentId"])

	utils.CheckErrors(w, err, http.StatusBadRequest)

	if r.Method == "PUT" {
		body, err := ioutil.ReadAll(r.Body)

		utils.CheckErrors(w, err, http.StatusBadRequest)

		var comment interfaces.Comment
		json.Unmarshal(body, &comment)
		comment.Id = commentId

		if &comment == nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = Comments.Update(&comment)

		utils.CheckErrors(w, err, http.StatusInternalServerError)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "comment "+vars["commentId"]+" updated!")
		return
	}

	if r.Method == "DELETE" {
		err := Comments.Delete(commentId)
		utils.CheckErrors(w, err, http.StatusInternalServerError)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "comment "+vars["commentId"]+" removed!")
		return
	}

}
