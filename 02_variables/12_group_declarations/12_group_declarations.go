package main

import "fmt"

func main() {
	var coffeType string = "Latte"
	var quantity int = 3
	var unitPrice float64 = 4.25

	fmt.Printf("Ordered %d %s priced %.2f each\n", quantity, coffeType, unitPrice)

	var (
		customerName string = "bogdan"
		tableNumber  int    = 8
		isReadyToPay bool   = false
	)
	fmt.Printf("Ordered by %s on table %d is ready to pay: %t\n", customerName, tableNumber, isReadyToPay)
	const (
		sizeSmall  = "S"
		sizeMdeium = "M"
		sizeLarge  = "L"
	)

}
