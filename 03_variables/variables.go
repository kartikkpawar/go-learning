package main

import "fmt"

func main()  {
	
	// type is assigned by you
	// var name string = "Kartik Pawar"
	
	// in this case type is infered by golang 
	// var name = "Kartik Pawar"

	// type: boolean
	// var isHappy bool = true
	
	// type: int
	// int have various types such as int8,16,32 either use them or only int (In this case golang internally handles it)
	// var age int = 32

	// short-hand syntax 
	// name:="Kartik Pawar"

	// Here the variable is declared and accessed later in such case type specification is necessary
	var name string
	name = "kartik pawar"

	fmt.Println(name)



}