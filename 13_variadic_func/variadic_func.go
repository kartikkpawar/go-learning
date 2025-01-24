package main

import "fmt"

func main() {
	// This is the example of the variadic function. Its basically the fuction which can accept n number of arguments
	fmt.Println(1, 2, 3, 4, 5, "Hello World")
	fmt.Println(sum(1, 2, 3, 4, 5))
}

// Currently the function is only accepting the int type. To accept any type use interface{} this is basically type: any of go
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total = total + num
	}
	return total
}
