package main

import "fmt"

func getDrinkInfo(customerName string, drinkType string, price float64) {
	info := "%s's favorite drink is %s and it's price is $%.2f.\n"
	fmt.Printf(info, customerName, drinkType, price)
}

func main() {
	getDrinkInfo("Jair", "Mocha", 4.50)
	getDrinkInfo("Aline", "Cappuccino", 3.50)
}
