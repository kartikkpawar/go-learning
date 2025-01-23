package main

import "fmt"

// slices are dynamic array: used when no. of elements inside arrays are not fixed

func main() {

	// unintialized silce is null
	// var nums []int //  no need to mention size

	var nums = make([]int,2,5) // This means make a slice with intial length of 2. And 5 is the initial capacity of the slice. Here the values are zeroy similar behaviour to array

	fmt.Println("Start Capacity",cap(nums)) // cap means Capacity which means the maximum number of elements the array can hold

	nums = append(nums, 1) // adding 1 to the array
	nums = append(nums, 2) // adding 2 to the array
	nums = append(nums, 3) // adding 3 to the array
	nums = append(nums, 4) // adding 4 to the array

	fmt.Println("After Adding Capacity",cap(nums)) // when the capacity is reached the capacity doubles

	fmt.Println(nums)

	// copy slice
	var nums2 = make([]int,len(nums),cap(nums))
	fmt.Println("Before Copy",nums,nums2)
	copy(nums2,nums) // copy is a inbuilt function it takes params (destination, source)
	fmt.Println("After Copy",nums,nums2)

	// Slice operator
	fmt.Println(nums[2:3]) // this is syntax for slicing the slice. [from_index : no_of_indexes_to_be_covered(last is excluded, i.e. in case 3 it will go to next 2 indexes)]. 
	// If from_index is not given it will be considered from 0
	// If no_of_indexes_to_be_covered is not given it will consider the length

	// 2D slices
	nums3:=[][]int{{4,5,6},{7,8,9},{10,11,12}}
	fmt.Println(nums3)
}