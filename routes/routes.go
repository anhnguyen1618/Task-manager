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

	// Pass db connections to controllers
	Controllers := &controllers.Controllers{env}

	r.Handle("/public", fs)
	r.HandleFunc("/", authMW(Controllers.LandingController))
	r.HandleFunc("/login", Controllers.LoginController)
	r.HandleFunc("/signUp", Controllers.SignUpController)
	r.HandleFunc("/signout", authMW(Controllers.SignOutController))

	r.HandleFunc("/tasks", authMW(Controllers.AllTaskController))
	r.HandleFunc("/tasks/{id}", authMW(Controllers.UpdateTaskController))

	r.HandleFunc("/tasks/{id}/comments", authMW(Controllers.CommentController))
	r.HandleFunc("/tasks/{id}/comments/{commentId}", authMW(Controllers.UpdateCommentController))

	http.HandleFunc("/", loggerMW(errorMW(r)))

}
