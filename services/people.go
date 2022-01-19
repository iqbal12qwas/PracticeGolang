package services

import (
	"practice/controllers"
	"practice/security"
)

func people_route() {
	// People Example Eloquent
	r.HandleFunc("/post_people", security.IsAuthorized(controllers.CreatePerson)).Methods("POST")
	r.HandleFunc("/get_people", security.IsAuthorized(controllers.GetAllPerson)).Methods("GET")
	r.HandleFunc("/get_people/{id}", security.IsAuthorized(controllers.GetPersonByID)).Methods("GET")
	r.HandleFunc("/update_people/{id}", security.IsAuthorized(controllers.UpdatePersonByID)).Methods("PUT")
	r.HandleFunc("/delete_people/{id}", security.IsAuthorized(controllers.DeletPersonByID)).Methods("DELETE")

}
