package routes

import (
	"net/http"

	"github.com/anhnguyen300795/Task-manager/controllers"
	"github.com/anhnguyen300795/Task-manager/interfaces"
	"github.com/anhnguyen300795/Task-manager/middlewares"
	"github.com/gorilla/mux"
)

func InititalizeRoutes(env *interfaces.Env) {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("/public"))

	// Declare shortHand middlewares
	middlewareEnv := &middlewares.MiddleWares{env}
	authMW := middlewareEnv.Authenticate
	loggerMW := middlewareEnv.Logger
	errorMW := middlewareEnv.MuxErrorHandler

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
	r.HandleFunc("/tasks/{id}/comments/{commentID}", authMW(Controllers.UpdateCommentController))

	http.HandleFunc("/", loggerMW(errorMW(r)))

}
