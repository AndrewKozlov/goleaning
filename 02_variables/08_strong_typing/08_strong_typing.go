package main

import "fmt"

func main() {
	price := 4.50 //float64

	quntity := 15 //int

	total := price * float64(quntity) //float64=float64*float64

	fmt.Printf("%.2f", total)
}
