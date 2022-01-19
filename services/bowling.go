package services

import (
	"practice/controllers"
	"practice/security"
)

func bowling_route() {
	// Bowling
	r.HandleFunc("/bowling", security.IsAuthorized(controllers.ChooseContainer)).Methods("GET")

}
