package routes

import (
	"log"
	"net/http"

	myHandler "github.com/LeNgocPhuc99/private-chat/server/handlers"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
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

	route.HandleFunc("/userLoginCheck/{userID}", myHandler.UserLoginCheck).Methods(http.MethodGet, http.MethodOptions)

	route.HandleFunc("/ws/{userID}", func(rw http.ResponseWriter, r *http.Request) {
		var upgrader = websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin:     func(r *http.Request) bool { return true },
		}

		_ = mux.Vars(r)["userID"]
		_, err := upgrader.Upgrade(rw, r, nil)
		if err != nil {
			log.Println(err)
			return
		}
	})

	log.Println("Routes are loaded")
	return route
}
