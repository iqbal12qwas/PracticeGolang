package security

import (
	"encoding/json"
	"fmt"
	"net/http"
	"practice/entity"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	secretkey string = "secretkeyjwt"
)

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
