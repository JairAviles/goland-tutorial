package main

import "fmt"

func main() {
	coffeeTypes := [4]string{"Espresso", "Latte", "Cappuccino", "Mocha"}
	fmt.Println("Types of coffee:", coffeeTypes)
	fmt.Println("Length of array:", len(coffeeTypes))

	coffeeTypes[len(coffeeTypes)-1] = "Americano"
	fmt.Println("Types of coffee:", coffeeTypes)

	fmt.Println("String length is:", len("This is a coffee string."))
}
