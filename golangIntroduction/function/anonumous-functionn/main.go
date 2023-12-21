package main

import "fmt"

func main() {
	func() {
		fmt.Println("hello world")
	}() // call the function directly

	// assign anon func to a variable
	hallo := func() {
		fmt.Println("Hello Universe!!")
	}

	hallo()

	// pass arguments to anon func

	func(word string) {
		fmt.Println(word)
	}("EnigmaCamp") // you can pass the arg when calling it directly

	// pass argument to anon func then assign it to a var

	hello2 := func(name string) {
		fmt.Println("Hello", name)
	}

	hello2("John")

	// passing anon func as argument
	greetEnglish := func(name string) string {
		return "Hello " + name
	}

	greetIndonesian := func(name string) string {
		return "Halo " + name
	}

	greet("Rick", greetEnglish)
	greet("Rick", greetIndonesian)

	add := func(num1, num2 int) int {
		return num1 + num2
	}

	multiply := func(num1, num2 int) int {
		return num1 * num2
	}

	calculate(3, 2, add)
	calculate(5, 2, multiply)
}

func greet(name string, f func(name string) string) {
	fmt.Println(f(name))
}

func calculate(a int, b int, operator func(x int, y int) int) {
	fmt.Println(operator(a, b))
}
