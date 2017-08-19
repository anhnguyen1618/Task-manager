package routes

import (
	"../controllers"
	"../middlewares"
	"net/http"
)

func InititalizeRoutes() {
	fs := http.FileServer(http.Dir("/public"))

	// http.Handle("/", httpInterceptor(router))
	http.Handle("/public", fs)
	http.Handle("/", middlewares.ApplyMiddleware(controllers.LandingController, middlewares.Logger))
	http.HandleFunc("/login", controllers.LoginController)
	http.HandleFunc("/signUp", controllers.SignUpController)

}
