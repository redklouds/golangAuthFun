package main

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

//our jwt secret DO NOT GIVE THIS AWAYY
var jwtKey = []byte("My_Secret_Man!")

//create a credential stuct to read the user name and password from the request body
type Crendentials struct {
	Password string `json:"pasword"`
	Username string `json:"username"`
}

//create a struct that will be encoded to a JWT.
//we add jwt.Standdard Claims as an embedded type to provide fields like expire time
//below we are telling go how to serialize this field so jsonify the field username as username

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func main() {

	//our routes
	http.HandleFunc("/signin", SignIn)
	http.HandleFunc("/Welcome", Welcome)
	http.HandleFunc("/refresh", Refresh)
	http.HandleFunc("/ping", ping)

	//our entry
	log.Fatal(http.ListenAndServe(":8080", nil))
}
