package main

import "fmt"

func processPayment(orderTotal float64, tip float64, amountPaid float64) (float64, float64) {
	totalAmountDue := orderTotal + tip
	change := amountPaid - totalAmountDue
	return totalAmountDue, change

}

func main() {
	totalCost, change := processPayment(6.50, 2.00, 10.00)
	fmt.Printf("Total cost (with tip): $%.2f\nChange return to customer: $%.2f\n", totalCost, change)
	fmt.Println("______________________")
	totalCost, change = processPayment(28.50, 1.50, 50.00)
	fmt.Printf("Total cost (with tip): $%.2f\nChange return to customer: $%.2f\n", totalCost, change)
}
