package main

//https://www.sohamkamani.com/blog/golang/2019-01-01-jwt-authentication/
import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func ping(res http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(res, "COOOL")

}

//Sign In Handler
func SignIn(res http.ResponseWriter, req *http.Request) {

	var creds Crendentials

	//get the JSON body and decode into credentials

	//this is where the tilda comes into play
	//ONLY words if the expected body serializes into the credential object meainig
	//there exist those exact fields, username and password
	err := json.NewDecoder(req.Body).Decode(&creds)
	if err != nil {
		//if the structure of the body, the request is wrong then throw an error
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	//we have a valid object

	//in Go this is how we retrive from a map as well as
	//check if the key exist,
	//accessing the map will return an error 'ok' as well as the value
	//for that key

	expectedPassword, ok := users[creds.Username]

	//if password exist for the given user: OK IS NOT NIL
	//AND if it is the SAME AS WE EXPECTED it,
	//if its not the same throw UNAUTHORIZED

	if !ok || expectedPassword != users[creds.Username] {
		res.WriteHeader(http.StatusUnauthorized)
		return
	}

	//we have a valid username and password proceed to moving forward

	//declare the expiration time for this token

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			//in JWT, the expiry time is expressed as unix miliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}

	//declare the token with the algorithm used for signing, and the claim

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//create the JWT String

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Finally, we set the client cookie for 'token' as the JWT we generated
	// we also set an expiry time which is  the same as  the token itself

	http.SetCookie(res, &http.Cookie{
		Name:    "Token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func Refresh(res http.ResponseWriter, req *http.Request) {

}
func Welcome(res http.ResponseWriter, req *http.Request) {

}
