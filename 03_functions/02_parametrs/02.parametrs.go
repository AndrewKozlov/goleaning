package main

import "fmt"

func greet(coffeeShopName string) {
	fmt.Println("Welcome to the Coffe Shop", coffeeShopName)
}

func main() {
	greet("Abobochka")
	greet("cooffeeeee and milk")
}
