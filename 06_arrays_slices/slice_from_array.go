package main

import "fmt"

func main() {
	dessertMenu := [4]string{"ice cream", "cake", "pie", "cookies"}
	fmt.Println("Dessert menu:", dessertMenu)

	slice := dessertMenu[1:3]
	fmt.Println("Slice of dessert menu [1:3] :", slice)

	slice = dessertMenu[:]
	fmt.Println("Slice of entire dessert menu [:] :", slice)

	slice = dessertMenu[2:]
	fmt.Println("Slice of dessert menu [2:] :", slice)

	slice = dessertMenu[:3]
	fmt.Println("Slice of dessert menu [:3] :", slice)
}
