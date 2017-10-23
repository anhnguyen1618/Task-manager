package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/anhnguyen300795/Task-manager/database"
	"github.com/anhnguyen300795/Task-manager/interfaces"
	"github.com/anhnguyen300795/Task-manager/routes"
)

func main() {
	db := database.Initialize()
	redisDB := database.InitializeRedis()

	env := &interfaces.Env{db, redisDB}

	routes.InititalizeRoutes(env)

	fmt.Println("Server is listening port: ", 8080)
	port := determineListenAddress()
	http.ListenAndServe(port, nil)
}

func determineListenAddress() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}
