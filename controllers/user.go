package controllers

import (
	"encoding/json"
	"net/http"
	"practice/db"
	"practice/entity"

	"github.com/gorilla/mux"
)

// Url Parameter
func GetUrlParamUserByID(w http.ResponseWriter, r *http.Request) {
	// Get Params By ID User
	vars := mux.Vars(r)
	key := vars["id"]

	// Con
	connection := db.Connector

	// Data statis Get
	var users entity.Tb_account

	connection.Raw("SELECT * FROM tb_account WHERE  id = ?", key).Scan(&users)

	// Show Response & Parse To Json
	w.Header().Set("Content-Type", "application/json")
	resp := entity.Response{Http: "GET/200", Message: "Success", Data: users}
	json.NewEncoder(w).Encode(resp)
}

// Request Parameter
func GetReqParamUserByID(w http.ResponseWriter, r *http.Request) {
	// Get Params By ID User
	query := r.URL.Query()
	key := query.Get("id")

	// Con
	connection := db.Connector

	// Data statis Get
	var user entity.Tb_account
	var users []entity.Tb_account
	// Show Response & Parse To Json
	w.Header().Set("Content-Type", "application/json")
	if key != "" { //If key/id = 2
		connection.Raw("SELECT * FROM tb_account WHERE  id = ?", key).Scan(&user)
		resp := entity.Response{Http: "GET/200", Message: "Get Data Success", Data: user}
		json.NewEncoder(w).Encode(resp)
	} else { //If key/id != 2
		connection.Raw("SELECT * FROM tb_account ORDER BY id ASC").Scan(&users)
		resp := entity.Response{Http: "GET/200", Message: "Get Data Success", Data: users}
		json.NewEncoder(w).Encode(resp)
	}

}

// Updated Data
func UpdatedUser(w http.ResponseWriter, r *http.Request) {
	var user entity.Tb_account

	vars := mux.Vars(r)
	key := vars["id"]

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		var err entity.Error
		err = SetError(err, "Error in reading payload.")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	// Con
	connection := db.Connector
	var authUser entity.User
	connection.Raw("SELECT * FROM tb_account WHERE email = ? ", user.Email).Scan(&authUser)
	connection.Raw("SELECT * FROM tb_account WHERE username = ? ", user.Username).Scan(&authUser)

	if authUser.Email != "" {
		var err entity.Error
		err = SetError(err, "Email has Exist")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	} else if authUser.Username != "" {
		var err entity.Error
		err = SetError(err, "Username has Exist")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	} else {
		sqlStatement := "UPDATE tb_account SET name = ?, email = ?, username = ? WHERE id = ?;"
		connection.Exec(sqlStatement, user.Name, user.Email, user.Username, key)
		resp := entity.Response{Http: "GET/200", Message: "Update Success", Data: user}
		json.NewEncoder(w).Encode(resp)
		return
	}

}
