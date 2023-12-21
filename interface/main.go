package main

import (
	"fmt"
)

type Shape interface {
	getArea() float64
	getPerimeter() float64
}

type Rectangle struct {
	width  float64
	length float64
}

func (r Rectangle) getArea() float64 {
	return r.width * r.length
}

func (r Rectangle) getPerimeter() float64 {
	return 2 * r.width * r.length
}

type Square struct {
	side float64
}

func (s Square) getArea() float64 {
	return s.side * s.side
}

func (s Square) getPerimeter() float64 {
	return 4 * s.side
}

func printGetArea(r Shape) {
	fmt.Println("Luas:", r.getArea())
}

func printAreaOfSquare(s Shape) {
	fmt.Println("Luas:", s.getArea())
}

func getArea(s Shape) {
	fmt.Println("Luas: ", s.getArea())
}

func main() {
	// upcasting membuate instance dari suatu struct yang bertipe interface
	var shape1 Shape = Square{side: 5}
	var shape2 Shape = Rectangle{width: 4, length: 3}
	var shape3 Shape = Rectangle{width: 7, length: 9}

	fmt.Printf("shape1: %#v\n", shape1)
	fmt.Printf("shape2: %#v\n", shape2)
	fmt.Printf("shape3: %#v\n", shape3)

	fmt.Println()
	fmt.Println("===============================================")

	shapes := []Shape{shape1, shape2, shape3}

	for _, shape := range shapes {
		fmt.Printf("%#v memiliki luas %.1f dan keliling %.1f\n", shape, shape.getArea(), shape.getPerimeter())
	}

	fmt.Println()
	fmt.Println("===============================================")

	printGetArea(Rectangle{width: 3, length: 7})
	printAreaOfSquare(Square{side: 9})

	fmt.Println()
	fmt.Println("===================with shape==================")

	getArea(Rectangle{width: 6, length: 12})
	getArea(Square{side: 10})

	fmt.Println()
	fmt.Println("==================Down casting=================")
	// merubah tipe interface ke original struct

	var square1 Shape = Square{side: 6}
	fmt.Println("getArea:", square1.getArea())
	fmt.Println("getPerimeter:", square1.getPerimeter())

	var originalSquare Square = square1.(Square)
	fmt.Println("getArea:", originalSquare.getArea())
	fmt.Println("getPerimeter:", originalSquare.getPerimeter())
	fmt.Println("side:", originalSquare)

	fmt.Println()
	fmt.Println("=================interface kosong================")

	var anything any
	// var anything interface{}
	anything = "Kenji"
	fmt.Printf("anything %T: %v\n", anything, anything)

	anything = 20
	fmt.Printf("anything %T: %v\n", anything, anything)

	anything = false
	fmt.Printf("anything %T: %v\n", anything, anything)

	anything = []string{"Golang", "Java", "Python"}
	fmt.Printf("anything %T: %v\n", anything, anything)
}
