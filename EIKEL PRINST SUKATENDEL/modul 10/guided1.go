package main

import "fmt"

func main() {
	var usia int
	var punyaKK bool

	fmt.Print("Masukkan usia dan status KK (true/false): ")
	fmt.Scan(&usia, &punyaKK)

	if usia >= 17 && punyaKK {
		fmt.Println("Bisa membuat KTP")
	} else {
		fmt.Println("Belum bisa membuat KTP")
	}
}
