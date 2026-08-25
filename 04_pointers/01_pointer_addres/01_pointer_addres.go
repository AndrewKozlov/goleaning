package main

import "fmt"

func main() {
	coffee := "Espresso"
	pointer := &coffee

	fmt.Println("Coffe name for coffee variable:", coffee)
	fmt.Println("Memory location:", pointer)
	fmt.Printf("pointer adress: %p\n", pointer)

	fmt.Println("------------------")

	coffeeCopy := coffee //значение скопировно в другое место в памяти

	fmt.Println("Coffe name for coffeCopy:", coffeeCopy)
	fmt.Println("Memory location:", &coffeeCopy)
	fmt.Printf("pointer adress: %p\n", &coffeeCopy)
}
