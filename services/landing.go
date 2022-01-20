package services

import (
	"net/http"
	"practice/controllers"

	"github.com/gorilla/mux"
)

func view_route() {
	setStaticFolder(r)
	r.HandleFunc("/", controllers.Index).Methods("GET")
}

func setStaticFolder(route *mux.Router) {
	fs := http.FileServer(http.Dir("./public/"))
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))
}
