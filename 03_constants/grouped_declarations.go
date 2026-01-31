package main

import "fmt"

func main() {
	var coffeeType string = "coffee"
	var quantity int = 3
	var unitPrice float64 = 4.99

	fmt.Printf("Ordered %d %s priced $%.2f\n", quantity, coffeeType, unitPrice)

	var (
		customerName string = "Alice"
		tableNumber  int    = 5
		isReady      bool   = true
	)

	fmt.Printf("CustomerName: %s at table %d is ready to pay: %t\n", customerName, tableNumber, isReady)

	const (
		SizeSmall  = "S"
		SizeMedium = "M"
		SizeLarge  = "L"
	)

}
