package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	path     = "/home/trancends/enigmaCamp/golang/data/"
	fileName = "example.txt"
	filePath = path + fileName
)

func main() {
	createFile(filePath)

	var input string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan nama : ")
	scanner.Scan()
	input = scanner.Text()

	writeFile(input)
	readFile(filePath)
}

func writeFile(input string) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(input + "\n")
	writer.Flush()
}

func createFile(filePath string) {
	_, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		file, err := os.Create(filePath)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer file.Close()

		fmt.Println("File", fileName, "berhasil dibuat")
	}
}

func readFile(filePath string) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
