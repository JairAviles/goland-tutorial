package main

import "fmt"

func calculateAfterDiscount(price float64, discountRate float64) float64 {
	discountAmount := price - (price * discountRate)
	return discountAmount
}

func main() {

	var coffeePrice float64 = 5.00
	var discountRate float64 = 0.1
	fmt.Printf("Basic coffer price: $%.2f\n", coffeePrice)

	coffeePrice = calculateAfterDiscount(coffeePrice, discountRate)
	fmt.Printf("Coffee price after discount: $%.2f\n", coffeePrice)
}
