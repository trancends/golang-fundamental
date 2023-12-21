package main

import (
	"fmt"
	"strconv"
)

var names []string

func main() {
	var input string
	fmt.Println("names:", names)
	fmt.Print("Masukan jumlah nama :")
	fmt.Scan(&input)
	len, _ := strconv.Atoi(input)
	for i := 0; i < len; i++ {
		fmt.Print("Masukkan Nama :")
		fmt.Scanln(&input)
		names = append(names, input)

	}
	fmt.Println("names:", names)
}
