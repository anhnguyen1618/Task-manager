package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/anhnguyen300795/Task-manager/src/interfaces"
	"github.com/anhnguyen300795/Task-manager/src/models"
	"github.com/anhnguyen300795/Task-manager/src/utils"
	"github.com/gorilla/mux"
)

func (controller *Controllers) CommentController(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	taskId, err := strconv.Atoi(vars["id"])

	utils.CheckErrors(w, err, http.StatusBadRequest)

	if r.Method == "POST" {
		body, err := ioutil.ReadAll(r.Body)

		utils.CheckErrors(w, err, http.StatusInternalServerError)

		var comment interfaces.Comment
		json.Unmarshal(body, &comment)

		Comments := &models.Comments{controller.DB}

		commentID, err := Comments.Add(&comment, taskId)

		utils.CheckErrors(w, err, http.StatusInternalServerError)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Comment "+strconv.FormatInt(commentID, 10)+" added!")
		return
	}

}

func (controller *Controllers) UpdateCommentController(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	commentID, err := strconv.Atoi(vars["commentID"])

	utils.CheckErrors(w, err, http.StatusBadRequest)

	user := utils.ExtractContext(r)

	Comments := &models.Comments{controller.DB}
	currentComment := Comments.GetByID(commentID)

	if currentComment == nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Comment "+vars["commentID"]+" not found!")
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
		updatedComment.Id = commentID

		if &updatedComment == nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		err = Comments.Update(&updatedComment)

		utils.CheckErrors(w, err, http.StatusInternalServerError)

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "comment "+vars["commentID"]+" updated!")
		return
	}

	if r.Method == "DELETE" {
		err := Comments.Delete(commentID)
		utils.CheckErrors(w, err, http.StatusInternalServerError)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "comment "+vars["commentID"]+" removed!")
		return
	}
}
