package main

import (
	"fmt"
	"time"
)

// NOTE: There are no classes in GO

// order struct

type customer struct {
	name  string
	phone string
}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // Time has nanoseconds precision
	customer            // This way 2 struct can be linked
}

// Adding methods to structs

// to attach method to struct before function name (name *struct_name). similar concept to pointer for *
func (o *order) changeStatus(status string) {
	o.status = status
}

// if not modifying any value then * (use of pointer) is not required
func (o order) getAmount() float32 {
	return o.amount
}

func main() {

	// Creating instance from struct
	// If no value is set to the files then the default zero value is used
	myOrder := order{
		id:     "kartik",
		amount: 500,
		status: "PENDING",
		customer: customer{
			name: "Kartik Pawar",
		},
	}

	// Setting filed after instance creation
	myOrder.createdAt = time.Now()

	fmt.Println(myOrder)

	myOrder.changeStatus("CONFIRMED")

	// getting single filed from instace
	fmt.Println(myOrder.status)
	fmt.Println(myOrder.getAmount())

	myOrder.customer.name = "Kartik Pawar"
	myOrder.customer.phone = "+91-98826255367"
	fmt.Println(myOrder)

	// one more way to define struct
	// This is used when the struct is supposed to be used only once
	language := struct {
		name   string
		isGood bool
	}{"goLang", true} // this order should be same to the struct def order

	fmt.Println(language)
}
