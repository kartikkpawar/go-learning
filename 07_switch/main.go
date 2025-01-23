package main

import (
	"fmt"
	"time"
)

func main() {

	i := 20

	// normal switch
	// No need to write to break in go to break the statement
	switch i {
	case 1:
		fmt.Println(1)
	
	case 2:
		fmt.Println(2)

	case 3:
		fmt.Println(3)

		// Default case is optinoal in go-lang
	default:
		fmt.Println("Default Value")

	}

// Multivalue switch
	switch time.Now().Weekday(){

		// go also supports multiple conditions for same case
		case time.Saturday,time.Sunday:
			fmt.Println("Its Weekend")
		default:
			fmt.Println("Its workday")
	}

	// type switch

	// Fuctions can be stored in the variables 
	whoAmI:= func (i interface{})  {


		switch i.(type){  // i.(type) gives the type of the i 
		case int:
			fmt.Println("Its Integer")
		case string:
			fmt.Println("Its String")
		case bool:
			fmt.Println("Its Boolean")
		default:
			fmt.Println("Its Invalid (Other)")

		}
	}
	
	whoAmI("GoLang")  // string
	whoAmI(10)  // int
	whoAmI(10.3)  // other
	whoAmI(true)  // bool


}