package main

import "fmt"

func main() {
	numbers := []int{
		4, 1, 3, 7, 5,
	}

	kecil, besar := minMax(numbers)
	// if you only want to use one value
	// kecil, _ := minMax(numbers)
	fmt.Println("Kecil:", kecil)
	fmt.Println("Besar:", besar)
}

func minMax(numbers []int) (int, int) {
	min := numbers[0]
	max := numbers[0]

	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return min, max
}

// function with named return value, the returnc
// can be used as local variable

func minMaxNamed(numbers []int) (min int, max int) {
	min = numbers[0]
	max = numbers[0]

	for _, num := range numbers {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return min, max
}
