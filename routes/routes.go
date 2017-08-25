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
	http.HandleFunc("/", middlewares.ApplyMiddleware(controllers.LandingController, middlewares.Logger, middlewares.Authenticate))
	http.HandleFunc("/login", controllers.LoginController)
	http.HandleFunc("/signUp", controllers.SignUpController)
	http.HandleFunc("/signout", controllers.SignOutController)

}
