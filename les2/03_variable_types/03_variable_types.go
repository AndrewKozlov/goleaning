package main

import "fmt"

func main() {
	name := "Americano"
	price := 2.99
	ready := true
	orderedCount := 5
	stockCount := 5000

	fmt.Printf("%T %T %T %T %T\n", name, price, ready, orderedCount, stockCount)
}
