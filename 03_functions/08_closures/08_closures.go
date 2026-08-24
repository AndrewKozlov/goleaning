package main

import "fmt"

func createTemperatureAdjuster() (func(change float64) float64, float64, func() float64) {
	baseTemperature := 90.0

	adjustTemperature := func(change float64) float64 {
		baseTemperature += change
		return baseTemperature
	}

	getBaseTemperature := func() float64 {
		return baseTemperature
	}

	return adjustTemperature, baseTemperature, getBaseTemperature
}

func main() {
	adjustTemp, originalTemp, getBase := createTemperatureAdjuster()
	fmt.Printf("Original temperature %.1f\n", originalTemp)
	fmt.Printf("Adjusted temperature +1.5 %.1f grad C\n", adjustTemp(1.5)) // baseTemperature is changed
	fmt.Printf("Adjusted temperature +1.5 %.1f grad C\n", adjustTemp(-3.0))
	fmt.Println(getBase())
	fmt.Printf("Adjusted temperature +1.5 %.1f grad C\n", adjustTemp(5))
	fmt.Println(getBase())
	fmt.Printf("Original temperature %.1f\n", originalTemp)

	// 90.0
	// +1.5 -> 91.5
	// -3.0 -> 88,5
	// +5.0 -> 93.5

}
