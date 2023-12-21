package main

import "fmt"

func main() {
	x := 4
	y := x

	y = y + 1
	fmt.Println("x:", x)
	fmt.Println("y:", y)

	fmt.Println("Alamat x:", &x)
	fmt.Println("Alamat y:", &y)

	fmt.Println("===========pass by value when passing argument===========")

	a := 3
	fmt.Println("a:", a)
	increase(a)
}

func increase(n int) {
	n = n + 1
	fmt.Println("n:", n)
}
