package main

import "fmt"

// variables and const can also be declared outside the functions
// the shorthand for variables := cannot be used outside the variable

const age = 30
var years = 35

func main() {

	// decalring the constant this values cannot be reassigned. Vairable type is similar to 03_variables
	const name = "kartik"

	// cannot reassing to name
	// name = "pawar"

	const (
		port=5000
		host="localhost"
	)

	fmt.Println(port,host)
}