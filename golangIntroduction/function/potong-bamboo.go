package main

import (
	"fmt"
)

func main() {
	inputJumlah := 0
	potong := 0

	fmt.Print("Masukkan banyak bamboo: ")
	fmt.Scanln(&inputJumlah)
	bamboos := jumlahBamboo(inputJumlah)
	bamboos = sekatBamboos(bamboos)

	fmt.Print("Berapa kali potong? ")
	fmt.Scanln(&potong)
	potongBamboo(bamboos, potong)
}

func jumlahBamboo(num int) []int {
	return make([]int, num)
}

func sekatBamboos(nums []int) []int {
	for index := range nums {
		fmt.Printf("sekat bamboo ke %d : ", index+1)
		fmt.Scanln(&nums[index])
	}

	return nums
}

func potongBamboo(bamboos []int, potong int) {
	for i := 0; i < potong; i++ {
		fmt.Println("Potongan ke-", i+1)
		for index := range bamboos {
			if checkSekat(bamboos[index]) {
				bamboos[index] = 0
			} else {
				bamboos[index] -= 1
			}
			fmt.Println(bamboos[index])

		}

	}
}

func checkSekat(sekat int) bool {
	return sekat == 0
}
