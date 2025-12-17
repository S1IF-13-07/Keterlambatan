package main

import "fmt"

func main() {
	var huruf string

	fmt.Print("Masukkan huruf: ")
	fmt.Scan(&huruf)

	// Ambil karakter pertama
	r := []rune(huruf)[0]

	// Ubah huruf besar ke huruf kecil
	if r >= 'A' && r <= 'Z' {
		r = r + ('a' - 'A')
	}

	switch r {
	case 'a', 'i', 'u', 'e', 'o':
		fmt.Println("Vokal")
	default:
		if r >= 'a' && r <= 'z' {
			fmt.Println("Konsonan")
		} else {
			fmt.Println("Bukan huruf")
		}
	}
}
