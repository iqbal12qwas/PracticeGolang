package controllers

import (
	"encoding/json"
	"net/http"
	"practice/entity"
)

func GetValueSession(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-name")
	// log.Println("Session id : ", session.Values["id"])
	// log.Println("Session name : ", session.Values["name"])
	// log.Println("Session email : ", session.Values["email"])
	// log.Println("Session username : ", session.Values["username"])

	var data = map[string]interface{}{
		"id":       session.Values["id"],
		"name":     session.Values["name"],
		"email":    session.Values["email"],
		"username": session.Values["username"],
	}

	w.Header().Set("Content-Type", "application/json")
	resp := entity.Response{Http: "GET/200", Message: "Success", Data: data}
	json.NewEncoder(w).Encode(resp)
}
