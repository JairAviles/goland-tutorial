package main

import "fmt"

func estimateBrewTime(cupsQty int, secondsPerCup int) (totalTimeSeconds int, info string) {
	totalTimeSeconds = cupsQty * secondsPerCup
	info = "Coffee brewing estimated:"
	return
}

func main() {
	brewTime, information := estimateBrewTime(12, 20)
	fmt.Printf("%s %d seconds\n", information, brewTime)
}
