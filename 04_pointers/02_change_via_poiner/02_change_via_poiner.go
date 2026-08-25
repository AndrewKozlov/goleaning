package main

import "fmt"

func main() {
	var coffeePrice = 4.50
	fmt.Println("Coffee price = ", coffeePrice)
	//compile time : var coffePrice = 4.50
	//runtime (что видит компьютер): [какойтто адрес в памяти1] = 4.50

	//compile time : fmt.Println("Coffee price = ", coffeePrice)
	//runtime (что видит компьютер): fmt.Println([какой то адрес в памяти2], [какой то адресс в памяти1])
	fmt.Println("Memory address of price 4.50", &coffeePrice)
	coffeePrice = 5.00
	fmt.Println("Memory address of price 5.00", &coffeePrice)

	//pointerToCoffeePrice:=&coffeePrice аналогично следующей строке
	var pointerToCoffeePrice *float64 = &coffeePrice
	*pointerToCoffeePrice = 7.50 //перейтив место в памяти на которое указывает poinerToCoffeePrice и сменить там значение на 5.50
	fmt.Printf("Updated coffeePrice via memory: %.2f", coffeePrice)

}
