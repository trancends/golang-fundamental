package main

import (
	"fmt"
)

func main() {
	fruits := [4]string{"Apel", "Pisang", "Anggur", "Semangka"}

	fmt.Println(fruits)
	fmt.Println(fruits[1])
	fruits[2] = "Jeruk"
	fmt.Println(fruits)

	fmt.Println("===================================")

	var scores [3]int
	scores[0] = 87
	scores[1] = 78
	scores[2] = 92
	// scores[3] = 70 error out of bonds

	fmt.Println(scores)

	fmt.Println("===================================")

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	// days[7] = "Libur" error out of bounds

	fmt.Println(days)
	fmt.Println("Jumlah element,", len(days))

	fmt.Println("===========perulangan array dengan for==========")

	for i := 0; i < len(days); i++ {
		fmt.Printf("Elemen %d: %s\n", i, days[i])
	}

	fmt.Println()

	fmt.Println("===========perulangan array dengan range==========")

	for i, day := range days {
		fmt.Printf("Elemen %d: %s\n", i, day)
	}
	fmt.Println()

	fmt.Println("========perulangan array dengan range part2=======")

	for _, day := range days {
		fmt.Printf("Nama hari : %s\n", day)
	}

	fmt.Println("========perulangan array dengan range part3=======")

	for i := range days {
		fmt.Printf("index hari ke : %d\n", i)
	}
	fmt.Println()

	fmt.Println("========perulangan array dengan range part4=======")

	numbers := [5]int{3, 8, 2, 7, 4}

	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Println(number)
		}
	}

	fmt.Println()

	fmt.Println("========merubah nilai array dengan perulangan=======")

	fmt.Println("Sebelum: ", numbers)
	for i := 0; i < len(numbers); i++ {
		numbers[i] *= 2
	}
	fmt.Println("Sesudah: ", numbers)
	fmt.Println()

	fmt.Println("===============array 2d====================")

	matrix := [2][3]int{
		{3, 2, 3},
		{3, 4, 5},
	}

	fmt.Println(matrix)
	fmt.Println(matrix[1][1])

	// strings package can help
	// function from strings strings.Fields() and strings.Split()
	// create array with make function
	// exp with make:
	// var mark []float64 = make([]float64, int(s))
	// mark := make([]float64, int(s))
}
