package main

import "fmt"

func main() {
	var qty int = 3
	var processed bool = false

	fmt.Println("Quantity: ", qty)
	fmt.Println("Processed: ", processed)

	processed = true
	fmt.Println("Processed after update: ", processed)
}
