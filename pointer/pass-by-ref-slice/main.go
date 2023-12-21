package main

import "fmt"

func main() {
	// slice secara default adalah pass by reference
	numbers := []int{4, 3, 5, 11}
	anotherNumbers := numbers

	fmt.Println("numbers:", numbers)
	fmt.Println("anotherNumbers:", anotherNumbers)

	numbers[1] = 9
	fmt.Println("sesudah=========================")
	fmt.Println("numbers:", numbers)
	fmt.Println("anotherNumbers:", anotherNumbers)

	scores := []int{7, 9, 10}
	multipleBy10(scores)
	fmt.Println("scores:", scores)
}

// pass by reference pada slice juga berlaku ketika digunakan sebagai func param
func multipleBy10(numbers []int) {
	for i := range numbers {
		numbers[i] = numbers[i] * 10
	}

	fmt.Println("numbers di function:", numbers)
}
