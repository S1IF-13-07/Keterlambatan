package main

import "fmt"

func main() {
	var x, y int
	hasil := 0

	fmt.Print("Masukkan x dan y: ")
	fmt.Scan(&x, &y)

	for x >= y {
		x = x - y
		hasil++
	}

	fmt.Println("Hasil pembagian:", hasil)
}
