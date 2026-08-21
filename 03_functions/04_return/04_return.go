package main

import "fmt"

func calculateLoyaltyPoints(amountSpent float64) int {
	loyaltyPoints := int(amountSpent * 2)
	return loyaltyPoints
}

func addTotalPoints(currentPoints int, newPoints int) int {
	return currentPoints + newPoints
}

func main() {
	totalPoints := 120
	var newarlyEarnedPoints int = calculateLoyaltyPoints(5.90)
	fmt.Println("Earned points today:", newarlyEarnedPoints)
	totalPoints = addTotalPoints(totalPoints, newarlyEarnedPoints)
	fmt.Println("Updated loyalty points:", totalPoints)
}
