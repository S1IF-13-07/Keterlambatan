package main

import "fmt"

func main() {
	var orang int

	fmt.Print("Masukkan jumlah orang: ")
	fmt.Scan(&orang)

	// Jika jumlah orang ganjil, tambahkan 1
	if orang%2 != 0 {
		orang = orang + 1
	}

	motor := orang / 2
	fmt.Println("Jumlah motor yang dibutuhkan:", motor)
}
