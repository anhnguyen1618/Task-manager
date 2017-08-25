package main

import (
	"./database"
	"./routes"
	"fmt"
	"net/http"
)

func main() {
	database.Initialize()
	database.InitializeRedis()
	defer database.Close()

	routes.InititalizeRoutes()
	http.ListenAndServe(":8080", nil)
	fmt.Println("Server is listening port: ", 8080)
}
