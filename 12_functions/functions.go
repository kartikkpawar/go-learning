package main

import "fmt"

// main function is not needed to be calles as it is the entry point of the file
func main() {
	fmt.Println(add(4, 5))

	// direct call for multiple import
	fmt.Println(getLanguage())

	// another way of getting the values
	lang1, lang2, lang3 := getLanguage()

	fmt.Println(lang1, lang2, lang3, " <-- Destructured languages")

	fn := func(a, b int) int {
		return a + b
	}

	fmt.Println(processIt(fn))

}

// if both the parameters are same it can also be written as func add (a,b int){}
func add(a int, b int) int {
	return a + b
}

// since 3 string values were return
func getLanguage() (string, string, string) {
	return "javascript", "goLang", "java"
}

// function accepting function as the parameter

func processIt(fn func(a, b int) int) int {
	return fn(1, 2)
}
