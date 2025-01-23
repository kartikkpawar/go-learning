package main

import "fmt"

func main() {

	var nums [4]int // Syntax to declare array. var array_name [array_length] array_element_type
	fmt.Println(len(nums)) //len is the method to get length of array

	nums[0] = 1 // adding element to array no push pop method present.

	fmt.Println(nums[0]) // getting the element of the array at index
	fmt.Println(nums) // getting all the values of the array

	// if the values are not filled in the array then go-lang fills it with the zeroy values
	// int -> 0, string -> "", bool -> false

	nums2:=[3]int{1,2,3} // Another way of declaring the array and add value at the time of declaring
	fmt.Println(nums2)

	// 2D arrays 
	nums3:=[3][3]int{{4,5,6},{7,8,9},{10,11,12}}
	fmt.Println(nums3)

}