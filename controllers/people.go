package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"practice/db"
	"practice/entity"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

var validate *validator.Validate

//GetAllPerson get all person data
func GetAllPerson(w http.ResponseWriter, r *http.Request) {
	var peoples []entity.People
	db.Connector.Find(&peoples)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(peoples)
}

//GetPersonByID returns person with specific ID
func GetPersonByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var people entity.People
	db.Connector.First(&people, key)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(people)
}

//CreatePerson creates person
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var people entity.People

	validate = validator.New()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Validation
	err := validate.Struct(people)
	if err != nil {

		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		errs := err.(validator.ValidationErrors)
		resp := entity.Error{IsError: true, Message: errs.Error()}
		json.NewEncoder(w).Encode(resp)
		return
	}

	json.Unmarshal(requestBody, &people)

	db.Connector.Save(&people)
	json.NewEncoder(w).Encode(people)
}

//UpdatePersonByID updates person with respective ID
func UpdatePersonByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	vars := mux.Vars(r)
	key := vars["id"]
	var people entity.People
	var update_people entity.People

	validate = validator.New()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Validation
	err := validate.Struct(update_people)
	if err != nil {

		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		errs := err.(validator.ValidationErrors)
		resp := entity.Error{IsError: true, Message: errs.Error()}
		json.NewEncoder(w).Encode(resp)
		return
	}

	json.Unmarshal(requestBody, &update_people)
	db.Connector.First(&people, key)
	db.Connector.Model(&people).Update(&update_people)

	json.NewEncoder(w).Encode(people)
}

//DeletPersonByID delete's person with specific ID
func DeletPersonByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var people entity.People
	id, _ := strconv.ParseInt(key, 10, 64)
	db.Connector.Where("id = ?", id).Delete(&people)
	w.WriteHeader(http.StatusNoContent)
}
