package services

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	r *mux.Router
)

//create a mux router
func CreateRouter() {
	r = mux.NewRouter()
}

//start the server
func ServerStart() {
	fmt.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080",
		handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}))(r))
	if err != nil {
		log.Fatal(err)
	}
}

//initialize all routes
func InitializeRoute() {

	// View
	view_route()

	// API
	auth_route()

	bowling_route()

	file_route()

	people_route()

	user_route()
}
