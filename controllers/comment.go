package controllers

import (
	"../interfaces"
	Comments "../models/comments"
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

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var comment interfaces.Comment
		json.Unmarshal(body, &comment)

		commentID, err := Comments.Add(&comment, taskId)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Comment "+strconv.FormatInt(commentID, 10)+" added!")
		return
	}

}

func UpdateCommentController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	commentId, err := strconv.Atoi(vars["commentId"])

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

		var comment interfaces.Comment
		json.Unmarshal(body, &comment)
		comment.Id = commentId

		if &comment == nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = Comments.Update(&comment)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "comment "+vars["commentId"]+" updated!")
		return
	}

	if r.Method == "DELETE" {
		err := Comments.Delete(commentId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "comment "+vars["commentId"]+" removed!")
		return
	}

}
