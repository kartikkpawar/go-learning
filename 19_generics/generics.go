package main

import "fmt"

// either the types can be specifed or any, interface{} can be used for all types to be allowed

func printSlice[T int | string](items []T) {
	for _, item := range items {
		fmt.Print(item, ",")
	}
	fmt.Println()
}

type stack[T any] struct {
	elements []T
}

func main() {

	// sliceInt := []int{1, 2, 3, 4, 5}
	// sliceStr := []string{"golang", "typescript", "javascript"}
	// // sliceBool := []bool{true, false, true, false}

	// printSlice(sliceInt)
	// printSlice(sliceStr)
	// printSlice(sliceBool) // will give error

	myStackInt := stack[int]{
		elements: []int{1, 2, 3, 4, 5, 6},
	}
	myStackStr := stack[string]{
		elements: []string{"golang", "typescript", "javascript"},
	}

	fmt.Println(myStackInt)
	fmt.Println(myStackStr)
}
