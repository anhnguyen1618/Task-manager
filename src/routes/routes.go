package routes

import (
	"net/http"

	"github.com/anhnguyen300795/Task-manager/src/controllers"
	"github.com/anhnguyen300795/Task-manager/src/interfaces"
	"github.com/anhnguyen300795/Task-manager/src/middlewares"
	"github.com/gorilla/mux"
)

func InititalizeRoutes(env *interfaces.Env) {
	r := mux.NewRouter()
	fs := http.FileServer(http.Dir("public"))

	// Declare shortHand middlewares
	middlewareEnv := &middlewares.MiddleWares{env}
	authMW := middlewareEnv.Authenticate
	loggerMW := middlewareEnv.Logger
	errorMW := middlewareEnv.MuxErrorHandler

	// Pass db connections to controllers
	Controllers := &controllers.Controllers{env}

	http.Handle("/public/", http.StripPrefix("/public/", fs))

	apiHandler := r.PathPrefix("/api").Subrouter()

	apiHandler.HandleFunc("/login", Controllers.LoginController)
	apiHandler.HandleFunc("/signUp", Controllers.SignUpController)
	apiHandler.HandleFunc("/logout", authMW(Controllers.SignOutController))
	apiHandler.HandleFunc("/currentUser", authMW(Controllers.CurrentUserController))

	apiHandler.HandleFunc("/users", authMW(Controllers.UsersController))
	apiHandler.HandleFunc("/users/{userName}", authMW(Controllers.UpdateUserController))

	apiHandler.HandleFunc("/tasks", authMW(Controllers.AllTaskController))
	apiHandler.HandleFunc("/tasks/{id}", authMW(Controllers.UpdateTaskController))

	apiHandler.HandleFunc("/tasks/{id}/comments", authMW(Controllers.CommentController))
	apiHandler.HandleFunc("/tasks/{id}/comments/{commentID}", authMW(Controllers.UpdateCommentController))

	r.PathPrefix("/").HandlerFunc(Controllers.LandingController)

	http.HandleFunc("/", loggerMW(errorMW(r)))

}
