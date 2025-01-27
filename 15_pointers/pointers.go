package main

import "fmt"

func main() {
	num := 1

	// In this the value is copied and passed i.e. pass by value
	changeNum(num)
	fmt.Println("In Main After Change before by ref", num)

	// &variable_name is used to see the memory address
	fmt.Println("Memory Address", &num)

	// This is pass by reference
	changeNumRef(&num)
	fmt.Println("In Main After Change after by ref", num)
}

func changeNum(num int) {
	num = 5
	fmt.Println("In ChangeNum", num)
}

// To recive pointer * is used before the variable type
func changeNumRef(num *int) {

	// De-refrencing the value. I.E. updating the value via pointer in the memory
	*num = 5
	fmt.Println("In ChangeNumRef", num)
}
