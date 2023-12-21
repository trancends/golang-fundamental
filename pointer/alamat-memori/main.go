package main

import "fmt"

func main() {
	name := "Beni"
	fmt.Println("name: ", name)
	fmt.Println("alamat dari name: ", &name)

	age := 25
	fmt.Println("age: ", age)
	fmt.Println("alamat dari age: ", &age)
}
