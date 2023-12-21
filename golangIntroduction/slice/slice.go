package main

import "fmt"

func main() {
	fmt.Println("========================================")
	numbersSlice := []int{2, 7, 9, 4}
	fmt.Println(numbersSlice)

	fmt.Println("Panjang slice: ", len(numbersSlice))
	fmt.Println("Kapasitas slice: ", cap(numbersSlice))

	fmt.Println("=========deklarasi dengan make ==========")

	scores := make([]int, 4, 7) // argument ke 3 opsional
	scores[0] = 89
	scores[1] = 89
	scores[2] = 89
	scores[3] = 89
	fmt.Println(scores)
	fmt.Println("Panjang slice: ", len(scores))
	fmt.Println("Kapasitas slice: ", cap(scores))

	fmt.Println("========================================")
	scores2 := make([]int, 4)
	scores[0] = 89
	scores[1] = 89
	scores[2] = 89
	scores[3] = 89
	fmt.Println(scores)
	fmt.Println("Panjang slice: ", len(scores2))
	fmt.Println("Kapasitas slice: ", cap(scores2))

	fmt.Println("======menambahkan element pada slice======")

	villain := make([]string, 3, 5)
	villain[0] = "Thanos"
	villain[1] = "Homelander"
	villain[2] = "Joker"
	fmt.Println("Villain: ", villain)
	fmt.Println("Panjang villain: ", len(villain))
	fmt.Println("Kapasitas villain: ", cap(villain))
	villain = append(villain, "Ultron")
	villain = append(villain, "Sauron", "Voldemort")
	fmt.Println("Villain: ", villain)
	fmt.Println("Panjang villain: ", len(villain))
	fmt.Println("Kapasitas villain: ", cap(villain))

	fmt.Println("======array passed by value example======")

	// array is passed by value
	numbers := [4]int{2, 1, 6, 8}
	anotherNumbers := numbers
	fmt.Println("Numbers: ", numbers)
	fmt.Println("AnotherNumbers: ", anotherNumbers)
	anotherNumbers[1] = 11
	fmt.Println("Numbers: ", numbers)
	fmt.Println("AnotherNumbers: ", anotherNumbers)

	fmt.Println("====== slice is passed by ref example======")

	prices := []int{2000, 1300, 1000}
	anotherPrices := prices
	fmt.Println("Prices: ", prices)
	fmt.Println("Anotherprices: ", anotherPrices)
	anotherPrices[1] = 1100
	fmt.Println("Prices: ", prices)
	fmt.Println("Anotherprices: ", anotherPrices)

	fmt.Println("========================================")

	ages := [4]int{20, 22, 25, 19}
	sliceAges := ages[0:3]
	fmt.Println("Ages : ", ages)
	fmt.Println("Slice Ages: ", sliceAges)
	sliceAges[2] = 33
	fmt.Println("Ages : ", ages)
	fmt.Println("Slice Ages: ", sliceAges)
	sliceAges = append(sliceAges, 37, 23)
	fmt.Println("Ages : ", ages)
	fmt.Println("Slice Ages: ", sliceAges)
	sliceAges[2] = 25
	sliceAges[3] = 19
	fmt.Println("Ages : ", ages)
	fmt.Println("Slice Ages: ", sliceAges)
}
