package main

import "fmt"

func main() {
	menu := [3]string{"egg", "bacon", "sausage"}
	slice := menu[:2]

	fmt.Println("Before slice change:")
	fmt.Println("menu:", menu)
	fmt.Println("slice:", slice)

	fmt.Println("Length of menu:", len(menu))
	fmt.Println("Length of slice:", len(slice))

	slice[0] = "hot dog"

	fmt.Println("After slice change:")
	fmt.Println("menu:", menu)
	fmt.Println("slice:", slice)
}
