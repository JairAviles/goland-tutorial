package main

import "fmt"

func applyDiscount(price *float64, discountRate float64) {
	*price = *price - (*price * discountRate)
}

func main() {

	var coffeePrice float64 = 5.00
	fmt.Println("Memory address of coffeePrice in main:", &coffeePrice)
	var discountRate float64 = 0.1
	fmt.Printf("Basic coffer price: $%.2f\n", coffeePrice)

	applyDiscount(&coffeePrice, discountRate)
	fmt.Printf("Coffee price after discount: $%.2f\n", coffeePrice)
}
