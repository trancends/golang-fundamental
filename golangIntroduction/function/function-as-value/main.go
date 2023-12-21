package main

import "fmt"

func main() {
	// param pada func harus sama dengan add
	var sum func(int, int) int = add
	fmt.Println("Hasil:", sum(3, 5))
}

func add(a, b int) int {
	return a + b
}
