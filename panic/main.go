package main

import "fmt"

// panic merupakan kondisi ketika terdapat error yang tidak
// dihandle
// panic terjadi ketika ada kesalahan logic
// panic dapat dilakukan dengan sengaja
// code setelah panic tidak akan dijalankan
func main() {
	// pemanggilan recover harus dilakukan bersamaan dengan defer
	// dan dilakukan sebelum panic
	// this code below wont work
	// r := recover()
	// if r != nil {
	// 	fmt.Println("Terjadi panic: ", r)
	// }

	// this work
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("Terjadi panic: ", r)
	// 	}
	// }()

	// var input string
	// fmt.Print("Masukan Nama :")
	// fmt.Scanln(&input)
	//
	// fmt.Println("Start")
	//
	// if !isEmtpy(input) {
	// 	fmt.Println(input)
	// }
	//
	// fmt.Println("done")

	defer func() {
		fmt.Println("Deferred Function")
	}()

	fmt.Println("Before panic")
	panic("Panic occured")
	defer func() {
		fmt.Println("Deferred Function")
	}()
}

func isEmtpy(input string) (empty bool) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Terjadi panic: ", r)
			empty = false
		}
	}()
	if input == "" {
		panic("Input tidak boleh kosong")
	}
	empty = false
	return
}
