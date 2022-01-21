package services

import (
	"practice/controllers"
	"practice/security"
)

func book_route() {
	// Book Example Eloquent
	r.HandleFunc("/post_book", security.IsAuthorized(controllers.CreateBook)).Methods("POST")
	r.HandleFunc("/get_book", security.IsAuthorized(controllers.GetAllBook)).Methods("GET")
	r.HandleFunc("/update_book/{id}", security.IsAuthorized(controllers.UpdateBookByID)).Methods("PUT")
	r.HandleFunc("/delete_book/{id}", security.IsAuthorized(controllers.DeletBookByID)).Methods("DELETE")

}
