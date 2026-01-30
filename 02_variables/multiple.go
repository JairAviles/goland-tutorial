package main

import "fmt"

func main() {
	var coffeeName = "Mocha"
	var size = "Medium"
	var price = 3.75

	fmt.Println("Medium Havana Cappuccino is $3.75")
	fmt.Println(size, coffeeName, "is $", price)
	fmt.Printf("%s %s is $%.2f \n", size, coffeeName, price)
}
