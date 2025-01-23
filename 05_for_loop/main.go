package main

import "fmt"

// for it the only construct in go for looping
func main()  {
	//  While loop
	i:=1
	for i<=3{
		fmt.Println(i)
		i= i+1
	}
	fmt.Println("While Loop ended")

	// Regular for loop
	for j:=0;j<3;j++{
		fmt.Println(j)
	}
	fmt.Println("For Loop ended")

// break and continue keyword works in same fashion as other lang

// Another way of writing for loop using range keyword. Range is introduced in go after v1.22.
// ranges starts from 0 till given value -1 so range 11 will go from 0 to 10
for k:=range 11{
	fmt.Println(k)
}
fmt.Println("For Loop with range ended")


}