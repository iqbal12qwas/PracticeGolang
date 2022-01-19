package services

import (
	"practice/controllers"
	"practice/security"
)

func user_route() {
	// User Example Query Builder
	r.HandleFunc("/user/{id}", security.IsAuthorized(controllers.GetUrlParamUserByID)).Methods("GET")
	r.HandleFunc("/user", security.IsAuthorized(controllers.GetReqParamUserByID)).Methods("GET")
	r.HandleFunc("/update_user/{id}", security.IsAuthorized(controllers.UpdatedUser)).Methods("PUT")
}
