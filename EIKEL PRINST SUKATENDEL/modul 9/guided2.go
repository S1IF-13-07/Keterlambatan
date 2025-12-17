package main

import "fmt"

func main() {
	var angka int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&angka)

	if angka > 0 {
		fmt.Println("Bilangan Positif")
	} else {
		fmt.Println("Bukan bilangan positif")
	}
}
