package main

import "fmt"

func main() {
	var coffeeName string = "Espresso"

	var size = "Small"

	price := 2.50

	fmt.Println("Small Espresso price is $2.60")

	fmt.Println(size, coffeeName, "price is $", price)
	fmt.Printf("%s %s price is $%.2f\n", size, coffeeName, price)
}
