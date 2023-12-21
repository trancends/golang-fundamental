package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var firstName, lastName string
	var age int

	fmt.Print("Enter your name: ")
	fmt.Scanln(&firstName, &lastName)
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)

	fmt.Println("Your name is ", firstName, lastName, "You are born in ", 2023-age)

	fmt.Println("==========================================")

	var fullName, address string

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Data diri peserta Enigma Camp")
	fmt.Print("Masukan nama anda : ")
	scanner.Scan()
	fullName = scanner.Text()
	fmt.Print("Masukan Umur Anda : ")
	scanner.Scan()
	age, _ = strconv.Atoi(scanner.Text())
	fmt.Print("Masukan Alamat Anda : ")
	scanner.Scan()
	address = scanner.Text()
	fmt.Printf("%s | %d | %s", fullName, age, address)
}
