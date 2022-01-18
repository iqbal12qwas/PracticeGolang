package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"practice/controllers"
	"practice/db"
	"practice/entity"

	jwt "github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	router    *mux.Router
	secretkey string = "secretkeyjwt"
)

//create a mux router
func CreateRouter() {
	router = mux.NewRouter()
}

func SetError(err entity.Error, message string) entity.Error {
	err.IsError = true
	err.Message = message
	return err
}

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Header["Token"] == nil {
			var err entity.Error
			err = SetError(err, "No Token Found")
			json.NewEncoder(w).Encode(err)
			return
		}

		var mySigningKey = []byte(secretkey)

		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing token.")
			}
			return mySigningKey, nil
		})

		if token != nil && err == nil {
			fmt.Println("token verified")
			handler.ServeHTTP(w, r)
			return
		} else {
			var reserr entity.Error
			reserr = SetError(reserr, "Not Authorized.")
			json.NewEncoder(w).Encode(err)
		}

		if err != nil {
			var err entity.Error
			err = SetError(err, "Your Token has been expired.")
			json.NewEncoder(w).Encode(err)
			return
		}
	}
}

//initialize all routes
func InitializeRoute() {
	// User Example Query Builder
	router.HandleFunc("/user/{id}", IsAuthorized(controllers.GetUrlParamUserByID)).Methods("GET")
	router.HandleFunc("/user", IsAuthorized(controllers.GetReqParamUserByID)).Methods("GET")
	router.HandleFunc("/update_user/{id}", IsAuthorized(controllers.UpdatedUser)).Methods("PUT")

	// Bowling
	router.HandleFunc("/bowling", IsAuthorized(controllers.ChooseContainer)).Methods("GET")

	// People Example Eloquent
	router.HandleFunc("/post_people", IsAuthorized(controllers.CreatePerson)).Methods("POST")
	router.HandleFunc("/get_people", IsAuthorized(controllers.GetAllPerson)).Methods("GET")
	router.HandleFunc("/get_people/{id}", IsAuthorized(controllers.GetPersonByID)).Methods("GET")
	router.HandleFunc("/update_people/{id}", IsAuthorized(controllers.UpdatePersonByID)).Methods("PUT")
	router.HandleFunc("/delete_people/{id}", IsAuthorized(controllers.DeletPersonByID)).Methods("DELETE")

	// Upload File
	router.HandleFunc("/upload", IsAuthorized(controllers.UploadImage)).Methods("POST")
	router.HandleFunc("/delete_file/{id}", IsAuthorized(controllers.DeleteFileByID)).Methods("DELETE")
	router.HandleFunc("/update_file/{id}", IsAuthorized(controllers.UpdateFileByID)).Methods("PUT")

	// Auth
	router.HandleFunc("/signin", controllers.SignIn).Methods("POST")
	router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
	})
}

//start the server
func ServerStart() {
	fmt.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router))
	if err != nil {
		log.Fatal(err)
	}
}

// Init DB
func initDB() {
	config :=
		db.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "",
			DB:         "user_db",
		}

	connectionString := db.GetConnectionString(config)
	err := db.Connect(connectionString)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	initDB()
	// db.MigratePeople(&entity.People{})
	// db.MigrateFile(&entity.File{})
	CreateRouter()
	InitializeRoute()
	ServerStart()
}
