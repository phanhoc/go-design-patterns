package factory

import "fmt"

func ExampleCashPM_Pay() {
	payment, err := GetPaymentMethod(CASH)
	if err != nil {
		fmt.Printf("Unable to get payment method")
		return
	}
	msg := payment.Pay(20)
	fmt.Printf(msg)
	//output:
	//You have just pay amount 20 by cash
}

func ExampleDebitCard_Pay() {
	payment, err := GetPaymentMethod(DEBIT_CARD)
	if err != nil {
		fmt.Printf("Unable to get payment method")
		return
	}
	msg := payment.Pay(20)
	fmt.Printf(msg)
	//output:
	//You have just pay amount 20 by debit card
}

func ExampleInvalidateType() {
	_, err := GetPaymentMethod(30)
	if err != nil {
		fmt.Printf("Error: %v", err)
	}
	//output:
	//Error: Unsupport type payment 30
}
