package routes

import (
	"net/http"

	"../controllers"
	"../interfaces"
	"../middlewares"
	"github.com/gorilla/mux"
)

func InititalizeRoutes(env *interfaces.Env) {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("/public"))

	// Declare shortHand middlewares
	authMW := middlewares.Authenticate
	loggerMW := middlewares.Logger
	errorMW := middlewares.MuxErrorHandler

	r.Handle("/public", fs)
	r.HandleFunc("/", authMW(controllers.LandingController(env)))
	r.HandleFunc("/login", controllers.LoginController(env))
	r.HandleFunc("/signUp", controllers.SignUpController(env))
	r.HandleFunc("/signout", authMW(controllers.SignOutController(env)))

	r.HandleFunc("/tasks", authMW(controllers.AllTaskController(env)))
	r.HandleFunc("/tasks/{id}", authMW(controllers.UpdateTaskController(env)))

	r.HandleFunc("/tasks/{id}/comments", authMW(controllers.CommentController(env)))
	r.HandleFunc("/tasks/{id}/comments/{commentId}", authMW(controllers.UpdateCommentController(env)))

	http.HandleFunc("/", loggerMW(errorMW(r)))

}
