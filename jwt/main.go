package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtKey = []byte("grace_liu_yay")

var users = map[string]string {
	"Apple Pie": "sweet",
	"Avocado Milkshake": "yummy",
}

// will be encoded to a jwt
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type Creds struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Signin(w http.ResponseWriter, r *http.Request) {
	var credential Creds
	// Decode request body into struct credential
	err := json.NewDecoder(r.Body).Decode(&credential)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// get expected password
	expectedPW, ok := users[credential.Username]
	if expectedPW != credential.Password || !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// expiration time of token
	expTime := time.Now().Add(5*time.Minute)
	claims := &Claims {
		Username: credential.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime,
		},
	}
	// sign claim using HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err !=nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// set client cookie for token as jwt
	http.SetCookie(w, &http.Cookie{
		Name:       "token",
		Value:      tokenString,
		Expires:    expTime
	})
}

func main() {
	http.HandleFunc("/signin", Signin)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
