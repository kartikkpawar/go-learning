package main

import "fmt"

func main() {

	// m := make(map[string]string)

	// // setting key:value
	// m["fname"] = "Kartik"
	// m["lname"] = "Pawar"

	// fmt.Println(m["fname"])
	// // If the key does not exists in map in such case zero value is return in this case ""
	// fmt.Println(len(m)) // len in map's case give number of {key:value} present in the map in this case 2

	// delete(m,"fname") // delete is used to delete the element from the map, args: (map_name , key_to_delete)
	// // clear(m) // Empties the map fully
	// fmt.Println(m)

	// Other way to declare map
	m := map[string]int{"price": 4, "phone": 3}
	fmt.Println(m)

	value,ok:=m["price"] // value is the value of the key given. If not present return zero value. and ok has boolean for wether the key is present or not

	fmt.Println(value) 

	if ok{
		fmt.Println("all ok")
	}else{
		fmt.Println("not found")
	}

}