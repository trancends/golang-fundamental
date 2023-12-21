package main

import "fmt"

func main() {
	fmt.Println("Hasil penjumlahan:", add(5, 2))

	total := add(12, 123)
	fmt.Println("Total:", total)

	total2 := multiply(5, 2)
	fmt.Println("Hasil perkalian:", total2)

	calculation := add(7, multiply(5, 2))
	fmt.Println("Hasil kalkulasi:", calculation)
}

func add(num1 int, num2 int) int {
	result := num1 + num2
	return result
}

func multiply(num1, num2 int) int {
	return num2 * num1
}
