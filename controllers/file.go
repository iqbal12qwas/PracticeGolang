package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"practice/db"
	"practice/entity"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// CreateImage : POST /api/images
func UploadImage(w http.ResponseWriter, r *http.Request) {
	var efile entity.File

	// Maximum upload of 10 MB files
	r.ParseMultipartForm(10 << 20)
	if r == nil {
		fmt.Fprintf(w, "No files can be selected\n")
	}

	// formdata := r.MultipartForm

	//get the *fileheaders
	// fil := formdata.File["files"]
	file, handler, err := r.FormFile("files")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()

	// fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	// fmt.Printf("File Size: %+v\n", handler.Size)
	// fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create file
	t := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	name_file := t + "-" + handler.Filename
	dst, err := os.Create(filepath.Join("file", filepath.Base(name_file)))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// return that we have successfully uploaded our file!

	insForm, err := db.Connector.CommonDB().Prepare("INSERT INTO files(name, path) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	} else {
		log.Println("data insert successfully . . .")
	}

	res_query, err := insForm.Exec(name_file, "file\\"+name_file)
	if err != nil {
		log.Println(err)
	}

	last_id, err := res_query.LastInsertId()
	if err != nil {
		log.Println(err)
	}

	efile.Id = last_id
	efile.Name = name_file
	efile.Path = "file\\" + name_file
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(efile)
}

//delete's file with specific ID
func DeleteFileByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var file_existing entity.File
	db.Connector.Raw("SELECT * FROM files WHERE  id = ?", key).Scan(&file_existing)

	resp := entity.Response{Http: "GET/200", Message: "Delete Success", Data: file_existing}

	// Removing file from the directory
	// Using Remove() function
	path := file_existing.Path
	e := os.Remove(path)
	if e != nil {
		log.Fatal(e)
	}

	var file entity.File
	id, _ := strconv.ParseInt(key, 10, 64)
	db.Connector.Where("id = ?", id).Delete(&file)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func UpdateFileByID(w http.ResponseWriter, r *http.Request) {

	// Delete File From Directory ===========
	vars := mux.Vars(r)
	key := vars["id"]

	var file_existing entity.File
	db.Connector.Raw("SELECT * FROM files WHERE  id = ?", key).Scan(&file_existing)

	path := file_existing.Path
	e := os.Remove(path)
	if e != nil {
		log.Fatal(e)
	}

	// Update Data File ==============
	var efile entity.File

	// Maximum upload of 10 MB files
	r.ParseMultipartForm(10 << 20)
	if r == nil {
		fmt.Fprintf(w, "No files can be selected\n")
	}

	// formdata := r.MultipartForm

	//get the *fileheaders
	// fil := formdata.File["files"]
	file, handler, err := r.FormFile("files")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	defer file.Close()

	// fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	// fmt.Printf("File Size: %+v\n", handler.Size)
	// fmt.Printf("MIME Header: %+v\n", handler.Header)

	// Create file
	t := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	name_file := t + "-" + handler.Filename
	dst, err := os.Create(filepath.Join("file", filepath.Base(name_file)))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Copy the uploaded file to the created file on the filesystem
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	upForm, err := db.Connector.CommonDB().Prepare("UPDATE files SET name = ?, path = ? WHERE id = ?;")
	if err != nil {
		panic(err.Error())
	} else {
		log.Println("data update successfully . . .")
	}

	upForm.Exec(name_file, "file\\"+name_file, key)

	efile.Name = name_file
	efile.Path = "file\\" + name_file
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(efile)
}
