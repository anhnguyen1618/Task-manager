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

	// http.Handle("/", httpInterceptor(router))
	r.Handle("/public", fs)
	r.HandleFunc("/", middlewares.ApplyMiddleware(controllers.LandingController, middlewares.Logger, middlewares.Authenticate))
	r.HandleFunc("/login", controllers.LoginController)
	r.HandleFunc("/signUp", controllers.SignUpController)
	r.HandleFunc("/signout", controllers.SignOutController)

	r.HandleFunc("/tasks", controllers.AllTaskController)
	r.HandleFunc("/tasks/{id}", controllers.UpdateTaskController)

	r.HandleFunc("/tasks/{id}/comments", controllers.CommentController)
	r.HandleFunc("/tasks/{id}/comments/{commentId}", controllers.UpdateCommentController)

	http.Handle("/", r)

}
