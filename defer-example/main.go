package main

import "fmt"

func main() {
	fmt.Println("defer selalu dieksekusi tepat sebelum fungsi selseai di eksekusi")
	// defer fmt.Println("defer Done")
	// fmt.Println("Start")
	// fmt.Println("Processing")

	// num := 4
	// nilai var dimasukan dengan nilai saat itu
	// namun pemanggilan defer tetap terakhir

	// num += 10
	// defer fmt.Println("Result Defer:", num)
	// num *= 2
	// fmt.Println("Result main: ", num)
	//
	// fmt.Println()
	// fmt.Println("=====defer if panic or error =====")

	// num1 := 4
	// num2 := 0

	// if error defer wont be called
	// but if it called before the panic
	// it will execute

	// fmt.Println("if error defer wont be called")
	// defer fmt.Println("Done")
	// fmt.Println("Start")
	// fmt.Println(num1 / num2)

	// fmt.Println()
	// fmt.Println("=====call func with defer=====")

	// defer still called last
	// defer add(7, 2)
	// defer add(7, multiply(2, 3)) // the ones that gor defered is the add()
	// add(3, 4)

	// fmt.Println()
	// fmt.Println("when there are multiple defer it will executes last in first out")
	// defer fmt.Println("Start")
	// defer fmt.Println("Done")
	// defer add(1, 2)
	// defer add(3, 4)
	// defer multiply(3, 3)
	//
	fmt.Println()
	fmt.Println("calling loop with defer still lifo")

	// fmt.Println("Start")
	// defer loop()
	// fmt.Println("Done")
	defer fmt.Println("Pertama")
	defer fmt.Println("Kedua")
	defer fmt.Println("Terakhir")
}

func add(num1 int, num2 int) {
	result := num1 + num2
	fmt.Println(result)
}

func multiply(num1, num2 int) int {
	result := num1 * num2
	fmt.Println(result)
	return result
}

func loop() {
	for i := 0; i <= 5; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done Loop")
}
