package main

import "fmt"

func getDrinkInfo(customerNmae string, drink string, price float64) {
	info := "%s`s fav drink is %s and it`s price is $%.2f\n"
	fmt.Printf(info, customerNmae, drink)
}

func main() {
	getDrinkInfo("Vasia", "Latte", 4.50)
	getDrinkInfo("Alice", "macciato", 2.99)
}
