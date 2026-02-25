package main

import "fmt"

func main() {
	menu := []string{"Egg", "bean"}

	fmt.Println("Initial menu:", menu)
	fmt.Println("Length:", len(menu), "Capacity:", cap(menu))
	fmt.Printf("Memory location of menu: %p\n", &menu)
	fmt.Printf("Memory location of \"Egg\": %p\n", &menu[0])

	menu = append(menu, "tomato")
	fmt.Println("After appending:", menu)
	fmt.Println("Length:", len(menu), "Capacity:", cap(menu))
	fmt.Printf("Memory location of menu: %p\n", &menu)
	fmt.Printf("Memory location of \"Egg\": %p\n", &menu[0])

	menu = append(menu, "carrot")
	fmt.Println("After appending again:", menu)
	fmt.Println("Length:", len(menu), "Capacity:", cap(menu))
	fmt.Printf("Memory location of menu: %p\n", &menu)
	fmt.Printf("Memory location of \"Egg\": %p\n", &menu[0])

	fmt.Println()

	cupSizes := make([]string, 0, 5)
	fmt.Println("Len of cupSizes:", len(cupSizes), "Cap of cupSize:", cap(cupSizes))
	//cupSize[0] = "Small"
	cupSizes = append(cupSizes, "Small", "Medium")
	fmt.Println("Len of cupSize:", len(cupSizes), "Cap of cupSize:", cap(cupSizes))
	fmt.Println(cupSizes)
	cupSizes[0] = "Extra Small"
	fmt.Println(cupSizes)
	fmt.Println("Len of cupSize:", len(cupSizes), "Cap of cupSize:", cap(cupSizes))
}
