package main

import "fmt"

func dynamicGreet(coffeeShopNme string) {
	greetMessage := "Welcome to"
	fmt.Println(greetMessage, coffeeShopNme, "coffee shop!")
}

func main() {
	dynamicGreet("Grand Central")
	dynamicGreet("Mi familia")
}
