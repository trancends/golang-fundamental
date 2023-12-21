package main

import "fmt"

func main() {
	var x int = 4
	var y *int = &x // pointer variable saves the memory address of x (pass by reference)

	fmt.Println("x:", x)
	fmt.Println("alamat x:", &x)
	fmt.Println("y:", y)
	fmt.Println("alamat y", &y)

	fmt.Println("nilai reference dari pointer y", *y) // the * (asterisk) symbol is used to get the value of the reference that point by y

	*y = *y + 1
	fmt.Println("x: ", x)
	fmt.Println("*y:", *y)

	fmt.Println("=========func pass by reference=========")
	var a int = 3
	increase(&a) // n *int = &a
	fmt.Println("a:", a)
	// pass by reference berlaku untuk semua tipe data
}

func increase(n *int) {
	*n = *n + 1
	fmt.Println("n di func increase:", *n)
}
