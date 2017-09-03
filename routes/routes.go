package routes

import (
	"../controllers"
	"../middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

func InititalizeRoutes() {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("/public"))

	// Declare shortHand middlewares
	authMW := middlewares.Authenticate

	r.Handle("/public", fs)
	r.HandleFunc("/", authMW(controllers.LandingController))
	r.HandleFunc("/login", controllers.LoginController)
	r.HandleFunc("/signUp", controllers.SignUpController)
	r.HandleFunc("/signout", authMW(controllers.SignOutController))

	r.HandleFunc("/tasks", authMW(controllers.AllTaskController))
	r.HandleFunc("/tasks/{id}", authMW(controllers.UpdateTaskController))

	r.HandleFunc("/tasks/{id}/comments", authMW(controllers.CommentController))
	r.HandleFunc("/tasks/{id}/comments/{commentId}", authMW(controllers.UpdateCommentController))

	http.HandleFunc("/", middlewares.Logger(middlewares.MuxErrorHandler(r)))

}
