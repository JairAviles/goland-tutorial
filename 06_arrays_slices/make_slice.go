package main

import "fmt"

func main() {
	ratings := []int{5, 4, 3, 5, 3}
	fmt.Println("Original ratings", ratings)

	ratings[2] = 4
	fmt.Println("Changed element with index", ratings)
	fmt.Println("Length of the slice is:", len(ratings))

	coffeeTypes := make([]string, 3)
	fmt.Println("Coffee types:", coffeeTypes)
}
