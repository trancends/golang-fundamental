package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Book struct {
	Id          int
	Title       string
	Author      string
	ReleaseYear string
	Pages       int
}

var books []Book

var fileName = "data.csv"

// kunci jawaban
// git.enigmacamp.com/enigma-camp-2.0/golang/golang-fundamental/simple-crud

func main() {
	// createFile(fileName)
	// addNewBook()
	readDataFromCSV(fileName)
	viewAllBooks()
}

func createFile(filename string) {
	_, err := os.Stat(filename)
	// if err == nil {
	// 	fmt.Println("File", filename, "sudah ada")
	// 	return
	// }

	if os.IsNotExist(err) {
		file, err := os.Create(filename)
		if err != nil {
			fmt.Println(err)
			return
		}

		defer file.Close()

		fmt.Println("File", fileName, "berhasil dibuat")
	}
}

func addNewBook() error {
	var newBook Book

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter Book Details :")
	fmt.Print("Book Id: ")
	scanner.Scan()
	newBook.Id, _ = strconv.Atoi(scanner.Text())

	fmt.Print("Book Title: ")
	scanner.Scan()
	newBook.Title = scanner.Text()

	fmt.Print("Book Author: ")
	scanner.Scan()
	newBook.Author = scanner.Text()

	fmt.Print("Book ReleaseYear: ")
	scanner.Scan()
	newBook.ReleaseYear = scanner.Text()

	fmt.Print("Book Pages: ")
	scanner.Scan()
	newBook.Pages, _ = strconv.Atoi(scanner.Text())

	_, err := findBookById(newBook.Id)
	if err != nil {
		books = append(books, newBook)
	} else {
		return fmt.Errorf("Book with id: %d already exist", newBook.Id)
	}

	err = writeDataToCSV(fileName)
	if err != nil {
		return err
	}

	fmt.Println("Book added successfully")

	return nil
}

func readDataFromCSV(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error opening csv file: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		record := strings.Split(scanner.Text(), ",")
		id, _ := strconv.Atoi(record[0])
		title := record[1]
		author := record[2]
		releaseYear := record[3]
		pages, _ := strconv.Atoi(record[4])
		book := Book{id, title, author, releaseYear, pages}
		books = append(books, book)
		// fmt.Println("=======================================")
		// fmt.Println("Book Id:", record[0])
		// fmt.Println("Book Title:", record[1])
		// fmt.Println("Book Author:", record[2])
		// fmt.Println("Book ReleaseYear:", record[3])
		// fmt.Println("Book Pages:", record[4])
	}

	return nil
}

func viewAllBooks() {
	if len(books) == 0 {
		fmt.Println("No books available")
	} else {
		for i, book := range books {
			fmt.Println(strings.Repeat("=", 30))
			fmt.Println("Book - ", i)
			fmt.Println("Book Id:", book.Id)
			fmt.Println("Book Title:", book.Title)
			fmt.Println("Book Author:", book.Author)
			fmt.Println("Book ReleaseYear:", book.ReleaseYear)
			fmt.Println("Book Pages:", book.Pages)
		}
	}
}

func writeDataToCSV(filename string) error {
	_, err := os.Stat(filename)
	if err == nil {
		fmt.Println("Menambahkan data....")
		file, _ := os.OpenFile(filename, os.O_APPEND|os.O_RDWR, 0666)
		for _, book := range books {
			row := strconv.Itoa(book.Id) + "," + book.Title + "," + book.Author + "," + book.ReleaseYear + "," + strconv.Itoa(book.Pages) + "\n"
			file.WriteString(row)
		}
		defer file.Close()

	} else {
		file, err := os.Create(filename)
		if err != nil {
			return fmt.Errorf("error opening csv file: %w", err)
		}

		defer file.Close()

		for _, book := range books {
			row := strconv.Itoa(book.Id) + "," + book.Title + "," + book.Author + "," + book.ReleaseYear + "," + strconv.Itoa(book.Pages) + "\n"
			fmt.Println(row)
			file.WriteString(row)
		}
	}

	return nil
}

func findBookById(id int) (Book, error) {
	for _, book := range books {
		if book.Id == id {
			return book, nil
		}
	}

	return Book{}, fmt.Errorf("id: %d not found", id)
}
