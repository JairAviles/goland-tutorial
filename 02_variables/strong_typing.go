package main

import "fmt"

func main() {
	price := 4.50
	quantity := 15
	total := price * float64(quantity)

	fmt.Printf("Total income: $%.2f\n", total)
}
