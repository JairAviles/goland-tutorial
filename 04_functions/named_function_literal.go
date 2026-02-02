package main

import "fmt"

func main() {
	taxRate := 0.10
	subtotal := 25.00

	calculateTax := func(amount float64, taxRate float64) float64 {
		return amount * taxRate
	}

	tax := calculateTax(subtotal, taxRate)
	total := subtotal + tax
	fmt.Printf("Total amount including tax: $%.2f\n", total)
}
