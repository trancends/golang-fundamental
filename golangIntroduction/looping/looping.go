package main

import "fmt"

func main() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Angka,", i)
	}

	fmt.Println("=====================")

	index := 1
	for index <= 10 {
		fmt.Println("Index,", index)
		index++
	}

	fmt.Println("==========================")

	index2 := 0
	for {
		fmt.Println("Index2,", index2)
		index2++
		if index2 == 5 {
			break
		}
	}

	fmt.Println("==========================")

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}

	fmt.Println("===========break===============")

	fmt.Println("Start")
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break
		}
		fmt.Print(i, " ")
	}
	fmt.Println("")
	fmt.Println("Done")

	fmt.Println("===========continue===============")

	fmt.Println("Start")
	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Print(i, " ")
	}
	fmt.Println("")
	fmt.Println("Done")

	fmt.Println("===========break & nested===============")

	fmt.Println("Start")
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 5; j++ {
			if j == 4 {
				break
			}
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
	fmt.Println("")
	fmt.Println("Done")

	fmt.Println("===========continue & nested===============")

	fmt.Println("Start")
	for i := 1; i <= 5; i++ {
		for j := 1; j <= 5; j++ {
			if j == 4 {
				continue
			}
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
	fmt.Println("")
	fmt.Println("Done")
}
