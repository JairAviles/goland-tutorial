package main

import "fmt"

func calculateLoyaltyPoints(purchaseAmount float64) int {
	return int(purchaseAmount * 2)
}

func updateTotalPoints(currentPoints int, pointsToAdd int) int {
	return currentPoints + pointsToAdd
}

func main() {
	totalPoints := 120
	earnedPoints := calculateLoyaltyPoints(9.30)
	fmt.Println("Earned points today:", earnedPoints)

	fmt.Println("Updated loyalty points:", updateTotalPoints(totalPoints, earnedPoints))
}
