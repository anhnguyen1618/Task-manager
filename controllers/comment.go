package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"../interfaces"
	"../models"
	"../utils"
	"github.com/gorilla/mux"
)

func CommentController(env *interfaces.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		taskId, err := strconv.Atoi(vars["id"])

		utils.CheckErrors(w, err, http.StatusBadRequest)

		if r.Method == "POST" {
			body, err := ioutil.ReadAll(r.Body)

			utils.CheckErrors(w, err, http.StatusInternalServerError)

			var comment interfaces.Comment
			json.Unmarshal(body, &comment)

			Comments := &models.Comments{env.DB}

			commentID, err := Comments.Add(&comment, taskId)

			utils.CheckErrors(w, err, http.StatusInternalServerError)

			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Comment "+strconv.FormatInt(commentID, 10)+" added!")
			return
		}
	}

}

func UpdateCommentController(env *interfaces.Env) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		commentId, err := strconv.Atoi(vars["commentId"])

		utils.CheckErrors(w, err, http.StatusBadRequest)

		user := utils.ExtractContext(r)

		Comments := &models.Comments{env.DB}
		currentComment := Comments.GetById(commentId)

		if currentComment == nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "Comment "+vars["commentId"]+" not found!")
			return
		}

		if user.Role != "ADMIN" && user.UserName != currentComment.Author {
			w.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(w, "You do not have permission for this action!")
			return
		}

		if r.Method == "PUT" {

			if user.UserName != currentComment.Author {
				w.WriteHeader(http.StatusForbidden)
				fmt.Fprintf(w, "You do not have permission for this action!")
				return
			}

			body, err := ioutil.ReadAll(r.Body)

			utils.CheckErrors(w, err, http.StatusBadRequest)

			var updatedComment interfaces.Comment
			json.Unmarshal(body, &updatedComment)
			updatedComment.Id = commentId

			if &updatedComment == nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			err = Comments.Update(&updatedComment)

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
}
