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

//GetAllBook get all book data
func GetAllBook(w http.ResponseWriter, r *http.Request) {
	// Get Params By ID Book
	query := r.URL.Query()
	key := query.Get("id")

	// Data statis Get
	var book []entity.Book
	// Show Response & Parse To Json
	w.Header().Set("Content-Type", "application/json")
	if key != "" { //If key/id != NULL
		db.Connector2.Find(&book, key)
		resp := entity.Response{Http: "GET/200", Message: "Get Detail Data Success", Data: book}
		json.NewEncoder(w).Encode(resp)
	} else { //If key/id = NULL
		db.Connector2.Find(&book)
		resp := entity.Response{Http: "GET/200", Message: "Get All Data Success", Data: book}
		json.NewEncoder(w).Encode(resp)
	}
}

//CreatePerson creates person
func CreateBook(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var book entity.Book

	validate = validator.New()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Validation
	err := validate.Struct(book)
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

	json.Unmarshal(requestBody, &book)

	db.Connector2.Save(&book)
	json.NewEncoder(w).Encode(book)
}

//UpdateBookByID updates book with respective ID
func UpdateBookByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	vars := mux.Vars(r)
	key := vars["id"]
	var book entity.Book
	var update_book entity.Book

	validate = validator.New()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Validation
	err := validate.Struct(update_book)
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

	json.Unmarshal(requestBody, &update_book)
	db.Connector2.First(&book, key)
	db.Connector2.Model(&book).Update(&update_book)

	json.NewEncoder(w).Encode(book)
}

//DeletBookByID delete's book with specific ID
func DeletBookByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var book_existing entity.Book
	db.Connector2.Raw("SELECT * FROM books WHERE  id = ?", key).Scan(&book_existing)

	resp := entity.Response{Http: "GET/200", Message: "Delete Success", Data: book_existing}

	var book entity.Book
	id, _ := strconv.ParseInt(key, 10, 64)
	db.Connector2.Where("id = ?", id).Delete(&book)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
