package auth

import "fmt"

// To expose fuction outside the package it should start with upper case variable
func LoginWithCredentials(username, password string) {

	fmt.Println("Logging user with", username, password)

}
