package main

import "fmt"

func main() {
	var coffeePrice = 4.50
	fmt.Println("Coffee price:", coffeePrice)
	fmt.Println("Memory address of coffeePrice $4.50:", &coffeePrice)

	coffeePrice = 5.00
	fmt.Println("Memory address of coffeePrice $5.00:", &coffeePrice)

	// pointerToCoffePrice := &coffeePrice
	var pointerToCoffeePrice *float64 = &coffeePrice
	*pointerToCoffeePrice = 7.50

	fmt.Println("Updated coffeePrice value in memory", coffeePrice)

}
