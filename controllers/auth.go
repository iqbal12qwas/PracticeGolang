package controllers

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"practice/db"
	"practice/entity"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var (
	secretkey string = "secretkeyjwt"
)

var store = sessions.NewCookieStore(securecookie.GenerateRandomKey(32))

func SetError(err entity.Error, message string) entity.Error {
	err.IsError = true
	err.Message = message
	return err
}

//take password as input and generate new hash password from it
func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//compare plain password with hash password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	// Connection
	connection := db.Connector

	var user entity.Authentication

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		var err entity.Error
		err = SetError(err, "Error in reading payload.")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	var authUser entity.User
	connection.Raw("SELECT * FROM tb_account WHERE email = ?", user.Email).Scan(&authUser)

	if authUser.Email == "" {
		var err entity.Error
		err = SetError(err, "Username is incorrect")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	} else {
		hasher := sha256.New()
		hasher.Write([]byte(user.Password))
		sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
		if sha != authUser.Password {
			var err entity.Error
			err = SetError(err, "Password is incorrect")
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(err)
			return
		}

	}

	validToken, err := GenerateJWT(authUser.Email)
	if err != nil {
		var err entity.Error
		err = SetError(err, "Failed to generate token")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	// Session Start
	sessions.NewSession(store, "session-name")
	// gob.Register(&entity.Tb_account{})
	session, _ := store.Get(r, "session-name")
	session.Values["id"] = authUser.Id
	session.Values["name"] = authUser.Name
	session.Values["email"] = authUser.Email
	session.Values["username"] = authUser.Username
	session.Save(r, w)

	var token entity.Token
	token.Email = authUser.Email
	token.TokenString = validToken
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}

//Generate JWT token
func GenerateJWT(email string) (string, error) {
	var mySigningKey = []byte(secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Println("Something went Wrong: " + err.Error())
		return "", err
	}

	return tokenString, nil
}
