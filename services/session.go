package services

import (
	"practice/controllers"
	"practice/security"
)

func session_route() {
	r.HandleFunc("/get_data_session", security.IsAuthorized(controllers.GetValueSession)).Methods("GET")
}
