package main

import "fmt"

// package scope variable has to be declared using var
var name = "Benedictus" // package scope or main scope?

func main() {
	helloWorld()
	fmt.Println("main: ", name)
	greet("Beni", 23)
	greet("Thanos", 1500)
}

func helloWorld() {
	name := "Jullian"
	fmt.Println("Hello World!")
	fmt.Println("Hello", name)
}

func greet(name string, age int) {
	fmt.Println("Hello my name is", name, "and I'm", age, "years old")
}
