package main

import "fmt"

type PaymentMethod interface {
	ProcessPayment(amount float64) string
}

type CreditCard struct {
	CardNumber string
}

func (cc CreditCard) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Process payment of $%.2f using Credit Card: %s", amount, cc.CardNumber)
}

type PayPal struct {
	Email string
}

func (pp PayPal) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Process payment of $%.2f using PayPal: %s", amount, pp.Email)
}

type BankTransfer struct {
	AccountNumber string
}

func (bt BankTransfer) ProcessPayment(amount float64) string {
	return fmt.Sprintf("Process payment of $%.2f using Bank Transfer: %s", amount, bt.AccountNumber)
}

func ProcessPayments(payments []PaymentMethod, amount float64) {
	for _, payment := range payments {
		result := payment.ProcessPayment(amount)
		fmt.Println(result)
	}
}

func main() {
	cc := CreditCard{CardNumber: "1234-5678-9102-2031"}
	pp := PayPal{Email: "example@example.com"}
	bt := BankTransfer{AccountNumber: "987654321"}

	payments := []PaymentMethod{cc, pp, bt}

	ProcessPayments(payments, 150.75)
}

/*
Explanation:
Interface Definition: The PaymentMethod interface defines a ProcessPayment method that all
payment types must implement.

Structs for Payment Types: We define three structs (CreditCard, PayPal, and BankTransfer),
each implementing the ProcessPayment method in their own way.

Function to Process Payments: The ProcessPayments function accepts a slice of PaymentMethod.
It iterates through each payment method and calls ProcessPayment.

Main Function: In main, we create instances of our payment types and store them in a slice.
We then call ProcessPayments to demonstrate how each payment type handles the transaction.

*/
