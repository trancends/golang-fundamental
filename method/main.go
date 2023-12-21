package main

import "fmt"

type triangle struct {
	base   float64
	height float64
}

// value receiver
func (t triangle) area() float64 { // receiver is a struct where the method attached to so we can manipulate the field of the struct
	return 0.5 * t.base * t.height
}

func (t *triangle) increaseSize() {
	t.base += 1
	t.height += 1
}

type Counter struct {
	Value int
}

func (c Counter) Increment() {
	c.Value++
}

func main() {
	instanceTriangle := triangle{10, 12}
	area := instanceTriangle.area()
	fmt.Println("area:", area)

	fmt.Println()
	fmt.Println("=========reference receiver========")

	fmt.Println("instanceTriangle:", instanceTriangle)
	instanceTriangle.increaseSize()
	fmt.Println("after increase:")
	fmt.Println("instanceTriangle:", instanceTriangle)

	counter := &Counter{Value: 3}
	counter.Increment()
	fmt.Println("Value counter after Increment", counter.Value)
}
