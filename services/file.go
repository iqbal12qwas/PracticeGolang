package services

import (
	"practice/controllers"
	"practice/security"
)

func file_route() {
	// Upload File
	r.HandleFunc("/upload", security.IsAuthorized(controllers.UploadImage)).Methods("POST")
	r.HandleFunc("/delete_file/{id}", security.IsAuthorized(controllers.DeleteFileByID)).Methods("DELETE")
	r.HandleFunc("/update_file/{id}", security.IsAuthorized(controllers.UpdateFileByID)).Methods("PUT")

}
