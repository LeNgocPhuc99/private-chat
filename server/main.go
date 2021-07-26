package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/LeNgocPhuc99/private-chat/server/database"
	"github.com/LeNgocPhuc99/private-chat/server/routes"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	fmt.Printf("Server will start at http://%s:%s\n", os.Getenv("HOST"), os.Getenv("PORT"))
	database.ConnectDatabase()

	route := routes.NewRoutes()

	log.Fatal(http.ListenAndServe(":8080", route))
}
