package main

import "fmt"

func main() {
	if false {
		fmt.Println("Kode dijalankan")
	}

	fmt.Println("Done")

	fmt.Println("==============================")

	result := "gagal"

	if result == "lulus" {
		fmt.Println("Selamat anda,", result)
	} else {
		fmt.Println("Maaf anda,", result)
	}

	fmt.Println("==============================")

	if hour := 21; hour > 8 && hour < 17 {
		fmt.Println("Masih dalam rentang waktu yang diperbolehkan")
	} else {
		fmt.Println("diluar batas waktu")
	}

	fmt.Println("==============================")

	GPA := 3.44
	var graduationStatus string

	if GPA >= 3.5 && GPA <= 4.0 {
		graduationStatus = "CUMLAUDE"
	} else if GPA >= 3.0 && GPA < 3.5 {
		graduationStatus = "MEMUASKAN"
	} else if GPA > 2.75 && GPA < 3.0 {
		graduationStatus = "CUKUP MEMUASKAN"
	} else {
		graduationStatus = "KURANG MEMUASKAN"
	}

	fmt.Printf("selamat kamu lulus dengan dengan predikat %s dengan IPK %.2f\n", graduationStatus, GPA)

	fmt.Println("==============================")

	x := 3
	y := -4

	if x > 0 {
		if y > 0 {
			fmt.Println("x dan y lebih dari 0")
		} else {
			fmt.Println("x lebih besar dari 0 dan y kurang dari atau sama dengan 0 ")
		}
	}

	fmt.Println("==============================")

	poin := 4
	switch poin {
	case 8:
		fmt.Println("bagus")
	case 7, 6, 5:
		fmt.Println("cukup")
	default:
		fmt.Println("kurang")
	}

	switch {
	case poin > 8:
		{
			fmt.Println("Bagus")
		}
	case poin < 8 && poin > 5:
		{
			fmt.Println("Cukup")
		}
	default:
		{
			fmt.Println("Nilaimu kurang bagus")
			fmt.Println("Belajar lebig giat lagi")
		}
	}

	jam := 3

	switch {
	case jam >= 4 && jam <= 5:
		fmt.Println("Bangun Pagi")
	case jam >= 6 && jam <= 7:
		fmt.Println("Mandi dan Sarapan")
	case jam >= 8 && jam <= 11:
		fmt.Println("Berangkat Sekolah")
	case jam == 12:
		fmt.Println("Pulang Sekolah")
	case jam >= 13 && jam <= 15:
		fmt.Println("Tidur Siang")
	case jam >= 16 && jam <= 17:
		fmt.Println("Waktu Main")
	default:
		fmt.Println("Waktu Belajar dan Istirahat")
	}
}
