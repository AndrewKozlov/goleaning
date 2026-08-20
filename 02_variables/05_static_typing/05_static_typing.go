package main

import "fmt"

func main() {
	//указываем тип явно
	var cupsQty int = 3

	//cupsQty ="four"//ошибка копиляции

	fmt.Println("number of cups:", cupsQty)

	//указываем тип явно
	var wasProcessed = true

	//cupsQty ="yes"//ошибка копиляции

	fmt.Println("Order was prodessed:", wasProcessed)

}
