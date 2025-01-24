package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5, 6}

	// using normal for loop
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	// same loop using range
	for index, num := range nums {
		fmt.Println(num, "Index is ",index)
	}

	// using range with map
	m := map[string]int{"price": 4, "phone": 3}

	for key, value := range m {
		fmt.Println(key, value)
	}

	// using range with string

	// byte size is the actual byte size of the char 
	// char gives the unicode and not the ascii code

	for byteSize,char :=range "��=� kartik pawar"{
		fmt.Println(string(char),byteSize)
	}	

}