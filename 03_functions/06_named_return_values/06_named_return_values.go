package main

import "fmt"

// func estimateBrewTime(cupsQty int, secondPerCup int) int {
// 	totalTimeSeconds := cupsQty * secondPerCup
// 	return totalTimeSeconds
// }

func estimateBrewTime(cupsQty int, secondPerCup int) (totalTimeSeconds int, info string) {
	info = "Estimated total brew time:"
	totalTimeSeconds = cupsQty * secondPerCup
	return //naked return
}

func main() {
	//12 cups, 20 sec per cup
	//12*20 =240 sec
	brewTime, info := estimateBrewTime(12, 20)
	fmt.Printf("%s %d\n", info, brewTime)
}
