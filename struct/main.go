package main

import "fmt"

type mesin struct {
	tipe string
	cc   int
}

type vehicle struct { // embedded struct, struct that used another struct as it's fields
	merek string
	tahun int
	model string
	harga int
	mesin
}

type kendaraan struct {
	merek string
	tahun int
	model string
	harga int
}

func main() {
	var a kendaraan
	a.merek = "Toyota"
	a.tahun = 2023
	a.model = "Inova"
	a.harga = 700000000

	fmt.Println("a:", a)
	fmt.Println("merek a:", a.merek)

	fmt.Println("==============another way to initialize struct===============")

	b := kendaraan{
		merek: "Wuling",
		tahun: 2023,
		model: "Conferro",
		harga: 450000000,
	}

	fmt.Println("b:", b)

	c := kendaraan{}
	c.merek = "Honda"
	c.tahun = 2023
	c.model = "Brio"
	c.harga = 250000000

	fmt.Println("c:", c)

	d := kendaraan{"Xenia", 2023, "Daihatsu", 300000000} // nilai field harus berurutan
	fmt.Println("d: ", d)

	fmt.Println()
	fmt.Println("==============pass by value with struct===============")

	x := kendaraan{
		merek: "Honda",
		tahun: 2023,
		model: "HRV",
		harga: 350000000,
	}

	y := x
	fmt.Println("Sebelum===========")
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("alamat x:", &x)
	fmt.Println("alamat y:", &y)

	fmt.Println("Sesudah===========")
	y.model = "CRV"
	fmt.Println("x:", x)
	fmt.Println("y:", y)

	fmt.Println()
	fmt.Println("==============Pass by value struct to func===============")

	z := kendaraan{model: "Civic", merek: "Honda", tahun: 2023, harga: 700000000}
	updateKendaraan(z)
	fmt.Println("kendaraan di main:", z)

	fmt.Println()
	fmt.Println("==============Pass by reference struct ===============")

	var z2 *kendaraan = &z
	fmt.Printf("alamat z: %p\n", &z)
	fmt.Printf("alamat z2: %p\n", z2)
	z2.model = "CRV"
	fmt.Println("z:", z)
	fmt.Println("z2:", *z2)

	fmt.Println()
	fmt.Println("==============Pass by reference struct to func===============")
	newZ := kendaraan{
		model: "CBR",
		merek: "Honda",
		harga: 25000000,
		tahun: 2023,
	}
	fmt.Println("newZ:", newZ)
	fmt.Println("after")
	updateKendaraanRef(&newZ)
	fmt.Println("newZ:", newZ)

	fmt.Println()
	fmt.Println("==============create struct pointer with new keyword===============")

	k := new(kendaraan)
	fmt.Printf("alamat k %p\n", k)

	fmt.Println()
	fmt.Println("==============embedded struct===============")

	newA := vehicle{
		merek: "Toyota",
		tahun: 2023,
		model: "Camry",
		harga: 450000000,
		mesin: mesin{
			tipe: "premium",
			cc:   2000,
		},
	}

	fmt.Println("newA :", newA)
	fmt.Println("tipe mesin newA:", newA.mesin.tipe)
}

func updateKendaraan(newKendaraan kendaraan) {
	newKendaraan.model = "Camry"
	newKendaraan.merek = "Toyota"
	newKendaraan.harga = 650000000
	newKendaraan.tahun = 2023
	fmt.Println("kendaraan di function:", newKendaraan)
}

func updateKendaraanRef(newKendaraan *kendaraan) {
	newKendaraan.model = "Camry"
	newKendaraan.merek = "Toyota"
	newKendaraan.harga = 650000000
	newKendaraan.tahun = 2023
	fmt.Println("kendaraan di function updateKendaraanRef:", newKendaraan)
}
