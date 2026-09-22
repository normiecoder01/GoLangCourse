package main

import "fmt"

type PaymentServices interface {
	handlePayment() string
}

type RazorPay struct {
	paymentID int
	senderID int
	recieverID int
	amount float64
	senderName string
	recieverName string
}

type PayPal struct {
	transactionID string
	senderID string
	recieverID string
	amount float64
	senderName string
	recieverName string
}

func (paypal PayPal) handlePayment () string {
	//business logic
	return fmt.Sprintln("Payment of Rs.",paypal.amount, "sent from", paypal.senderName, "to", paypal.recieverName,".")
}

func (razorPay RazorPay) handlePayment () string {
	//business logic
	return fmt.Sprintln("Payment of Rs.",razorPay.amount, "sent from", razorPay.senderName, "to", razorPay.recieverName,".")
}

func handlePayment2(ps PaymentServices) {
	fmt.Println(ps.handlePayment())	
}






func main(){
r1 := RazorPay{
	paymentID:    1001,
	senderID:     501,
	recieverID:   902,
	amount:       2499.75,
	senderName:   "Amit Sharma",
	recieverName: "Neha Patil",
}

r2 := RazorPay{
	paymentID:    1002,
	senderID:     777,
	recieverID:   888,
	amount:       499.00,
	senderName:   "Rahul Verma",
	recieverName: "Priya Deshmukh",
}

p1 := PayPal{
	transactionID: "TXN-9f3a7c21",
	senderID:      "user_abc123",
	recieverID:    "merchant_xyz789",
	amount:        19.99,
	senderName:    "John Doe",
	recieverName:  "Acme Corp",
}

p2 := PayPal{
	transactionID: "TXN-55b2d9ef",
	senderID:      "client_456def",
	recieverID:    "store_321ghi",
	amount:        150.50,
	senderName:    "Alice Brown",
	recieverName:  "Gadget Hub",
}

payments:=[]PaymentServices{p1,p2,r1,r2}

for _, payment := range payments{
	handlePayment2(payment)
}
fmt.Println("------------------------")
handlePayment2(p1)


}

