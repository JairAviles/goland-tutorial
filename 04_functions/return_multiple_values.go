package main

import "fmt"

func processPayment(orderTotal float64, tip float64, purchaseAmount float64) (float64, float64) {
	totalAmountDue := orderTotal + tip
	change := purchaseAmount - totalAmountDue
	return totalAmountDue, change
}

func main() {
	totalCost, change := processPayment(6.50, 2.00, 10.00)
	fmt.Printf("Total cost (with tip): %.2f\n", totalCost)
	fmt.Printf("Change returned: %.2f\n", change)
}
