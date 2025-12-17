package main

import "fmt"

func main() {
	var angka int

	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&angka)

	if angka < 0 {
		angka = angka * -1
	}

	fmt.Println("Nilai absolut:", angka)
}