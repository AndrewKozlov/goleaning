package main

import "fmt"

func main() {
	var coffeeTypes = [3]string{"Espresso", "Latte", "Cappuccino"}
	fmt.Println(coffeeTypes)
	coffeeSize := [3]string{"s", "m", "l"}
	fmt.Println(coffeeSize)

	fmt.Println("lenght of the arary:", len(coffeeSize))
	coffeeTypes[len(coffeeTypes)-1] = "milk"
	fmt.Println(coffeeTypes)

}
