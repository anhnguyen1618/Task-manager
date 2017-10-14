package main

import (
	"fmt"
	"net/http"

	"./database"
	"./interfaces"
	"./routes"
)

func main() {
	db := database.Initialize()
	redisDB := database.InitializeRedis()

	env := &interfaces.Env{db, redisDB}

	routes.InititalizeRoutes(env)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Server is listening port: ", 8080)
}
