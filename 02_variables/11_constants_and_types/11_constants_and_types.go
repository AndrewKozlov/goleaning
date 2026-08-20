package main

import "fmt"

func main() {
	const rewardPoints = 10 //не типизированная константа со значением integer
	fmt.Printf("Dedfaulr type of rewaedPoints is %T\n", rewardPoints)

	var totalRewardPoints float64 = 150.3

	totalRewardPoints += rewardPoints //добавление нетипизированной константы к float64 возможно, константа адаптируется. Работает тоьлко с const

	fmt.Printf("updated loyalty points %.2f\n", totalRewardPoints)
}
