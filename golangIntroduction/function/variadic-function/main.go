package main

import "fmt"

func main() {
	total := sum(2, 5, 6, 2)
	fmt.Println("Total:", total)
}

// variadic function, function yang dapat menerima
// lebih dari 1 argumen (dynamic argument)
// pada tiap pemanggilan function jumlah argumen
// bisa berbeda-beda

func sum(numbers ...int) int {
	// fmt.Println(numbers[0])
	// fmt.Println(numbers[1])
	// fmt.Println(numbers[2])
	// fmt.Println(numbers[3])
	// fmt.Printf("%T", numbers)

	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}
