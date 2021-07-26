package routes

import (
	"log"
	"net/http"

	myHandler "github.com/LeNgocPhuc99/private-chat/server/handlers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func NewRoutes() *mux.Router {
	log.Println("Loading Routes...")
	route := mux.NewRouter()
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
		handlers.AllowCredentials(),
	)
	route.Use(cors)

	// route.HandleFunc("/isUsernameAvailable/{username}", myHandler.IsUsernameAvailable)
	route.HandleFunc("/login", myHandler.Login).Methods(http.MethodPost, http.MethodOptions)
	route.HandleFunc("/registration", myHandler.Registration).Methods(http.MethodPost, http.MethodOptions)

	log.Println("Routes are loaded")
	return route
}
