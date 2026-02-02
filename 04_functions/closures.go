package main

import "fmt"

func createTemperatureAdjuster() (func(adjustment float64) float64, float64) {
	baseTemperature := 90.0

	adjustTemperature := func(adjustment float64) float64 {
		baseTemperature = baseTemperature + adjustment
		return baseTemperature
	}

	return adjustTemperature, baseTemperature
}

func main() {
	adjustTemp, originalTemp := createTemperatureAdjuster()
	fmt.Printf("Original temperature: %.1f\n", originalTemp)
	fmt.Printf("Adjusted temperature +1.5: %.1fC\n", adjustTemp(1.5))
	fmt.Printf("Adjusted temperature -3.0: %.1fC\n", adjustTemp(-3.0))
	fmt.Printf("Adjusted temperature 5.0: %.1fC\n", adjustTemp(5.0))
}
