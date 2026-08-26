package main

import "fmt"

func main() {
	dessertMenu := [4]string{"Maffin", "Brownie", "Croissant", "Cookie"}
	fmt.Println("Dessert Menu:", dessertMenu)

	slice := dessertMenu[1:2]
	fmt.Println("slice [1:2]: ", slice) //элементы с индексами 1 и

	slice = dessertMenu[:]
	fmt.Println("slice: ", slice) // все элемементы

	slice = dessertMenu[2:]
	fmt.Println("slice: ", slice) // все элемементы со 2

	slice = dessertMenu[:3]
	fmt.Println("slice: ", slice) // все элемементы до 3го не включительно
}
