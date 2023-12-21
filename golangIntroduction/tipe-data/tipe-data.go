package main

import "fmt"

func main() {
	fmt.Println("Tipe Data")
	fmt.Println("Numerik")
	var positiveNumber uint8 = 128
	var negativeNumber int8 = -128
	fmt.Printf("positiveNumber: %d\n", positiveNumber)
	fmt.Printf("negativeNumber: %d\n", negativeNumber)
	var pecahan float32 = 3.90
	fmt.Printf("Bilangan pecahan: %2.f\n", pecahan)

	fmt.Println("========================")

	fmt.Println("Boolean")
	var exist bool = true
	fmt.Printf("exist? %t", exist)

	fmt.Println("========================")
	fmt.Println("String")
	var message string = "hello"
	fmt.Printf("message: %s\n", message)

	otherMessage := `Nama saya "Beni"
Salam Kenal,
Mari belajar di "Eigma Camp"`
	fmt.Println(otherMessage)
}
