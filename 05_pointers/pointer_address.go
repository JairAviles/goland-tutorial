package main

import "fmt"

func main() {
	coffee := "Mocha"
	pointer := &coffee

	fmt.Println("Coffee name:", coffee)
	fmt.Println("Memory address of coffee variable:", pointer)
	fmt.Printf("Pointer address: %p\n", pointer)

	fmt.Println("_________________________________________________")

	coffeeCopy := coffee

	fmt.Println("Coffee copy name:", coffeeCopy)
	fmt.Println("Memory address of coffee variable:", &coffeeCopy)
	fmt.Printf("Pointer address: %p\n", &coffeeCopy)
}
