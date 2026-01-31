package main

import "fmt"

func main() {
	// Untyped constant with integer value
	const weekRewardPoints = 10

	fmt.Printf("Default type of weekRewardPoints: %T\n", weekRewardPoints)

	var totalRewardPoints float64 = 150.3

	// Adding untyped constant to a float64 variable
	totalRewardPoints = totalRewardPoints + weekRewardPoints

	fmt.Printf("Updated reward points: $%.2f\n", totalRewardPoints)
}
