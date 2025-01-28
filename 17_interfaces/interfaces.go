package main

import "fmt"

// naming convention for interface add r to end
type paymenter interface {
	pay(amount float32)
}

type payment struct {
	// Any strcut with pay method on it is allowed
	gateway paymenter
}

func (p payment) makePayment(amount float32) {

	// razorpayPaymentGw := razorpay{}
	// razorpayPaymentGw.pay(amount)

	// stripePaymentGw := stripe{}
	// stripePaymentGw.pay(amount)

	p.gateway.pay(amount)

}

type razorpay struct {
}

func (r razorpay) pay(amount float32) {
	fmt.Println("Making payment using razorpay. PAYMENT:", amount)
}

type stripe struct {
}

func (r stripe) pay(amount float32) {
	fmt.Println("Making payment using stripe. PAYMENT:", amount)
}

func main() {
	stripePaymentGw := stripe{}
	// razorpayPaymentGw := razorpay{}

	newPayment := payment{
		gateway: stripePaymentGw,
	}
	newPayment.makePayment(520)
}
