package main

import "fmt"

func calculateLoyaltyPoints(purchaseAmount float64) int {
	return int(purchaseAmount * 2)
}

func main() {
	earnedPoints := calculateLoyaltyPoints(9.50)
	fmt.Println("Earned points today:", earnedPoints)
}
