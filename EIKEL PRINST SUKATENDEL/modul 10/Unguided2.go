package main

import "fmt"

func main() {
	var nilai float64
	var nilaiHuruf string

	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scan(&nilai)

	if nilai > 80 {
		nilaiHuruf = "A"
	} else if nilai > 72.5 {
		nilaiHuruf = "AB"
	} else if nilai > 65 {
		nilaiHuruf = "B"
	} else if nilai > 57.5 {
		nilaiHuruf = "BC"
	} else if nilai > 50 {
		nilaiHuruf = "C"
	} else if nilai > 40 {
		nilaiHuruf = "D"
	} else {
		nilaiHuruf = "E"
	}

	fmt.Println("Nilai mata kuliah:", nilaiHuruf)
}
