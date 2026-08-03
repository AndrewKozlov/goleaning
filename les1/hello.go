package main

import "fmt"

func main() {
	mass := []int{7, 12, 4, 6, 78, 2, 52, 1, 2, 5}
	/*
		fmt.Println(summas(mass))
		fmt.Print(len(mass), "\n")
		mas_sort(mass)
		fmt.Println(mass)
	*/
	fmt.Println("before:", mass)
	fmt.Println("after:", qsort(mass))
}

func sum(n int) int {
	if n == 1 {
		return n
	}
	return n + sum(n-1)
}

func summas(n []int) int {
	itog := 0
	for _, i := range n {
		itog += i
	}
	return itog
}

func mas_sort(n []int) []int {
	var tmp int
	for j := 0; j < len(n); j++ {
		for i := 0; i < len(n)-1; i++ {
			if n[i] > n[i+1] {
				tmp = n[i]
				n[i] = n[i+1]
				n[i+1] = tmp
				fmt.Println(n[i], "знаечние", i)
			}

		}
	}
	return n
}

func qsort(n []int) []int {
	if len(n) <= 1 {
		return n
	}
	pivot := n[len(n)-1]
	var sorted_mass1, sorted_mass2 []int
	for i := 0; i < len(n)-1; i++ {
		if n[i] >= pivot {
			sorted_mass2 = append(sorted_mass2, n[i])
		} else {
			sorted_mass1 = append(sorted_mass1, n[i])
		}
	}
	return append(append(qsort(sorted_mass1), pivot), qsort(sorted_mass2)...)
}
