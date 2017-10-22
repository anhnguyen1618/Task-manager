package main

import (
	"fmt"
	"net/http"

	"github.com/anhnguyen300795/Task-manager/database"
	"github.com/anhnguyen300795/Task-manager/interfaces"
	"github.com/anhnguyen300795/Task-manager/routes"
)

func main() {
	db := database.Initialize()
	redisDB := database.InitializeRedis()

	env := &interfaces.Env{db, redisDB}

	routes.InititalizeRoutes(env)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Server is listening port: ", 8080)
}
