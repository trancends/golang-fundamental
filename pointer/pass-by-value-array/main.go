package main

import "fmt"

func main() {
	numbers := [4]int{4, 5, 3, 5}
	anotherNumbers := numbers

	fmt.Println("sebelum===========================")
	fmt.Println("numbers: ", numbers)
	fmt.Println("anotherNumbers", anotherNumbers)

	fmt.Println("sesudah===========================")
	anotherNumbers[0] = 10
	fmt.Println("numbers: ", numbers)
	fmt.Println("anotherNumbers", anotherNumbers)

	fmt.Println("=============================")
	scores := [4]int{7, 7, 8, 10}
	multiplyBy10(scores)
	fmt.Println("scores:", scores)
}

func multiplyBy10(n [4]int) {
	for i := range n {
		n[i] = n[i] * 10
	}
	fmt.Println("n in function:", n)
}
