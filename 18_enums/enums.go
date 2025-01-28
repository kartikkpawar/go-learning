package main

import "fmt"

// This is a custom type
type OrderStatus int

const (
	Received OrderStatus = iota
	Confirmed
	Prepared
	Delivered
)

func chnageOrderStatus(status OrderStatus) {
	fmt.Println("Changing order status:", status)
}

func main() {
	chnageOrderStatus(Delivered)
}
