package main

import "fmt"

func main() {
	var pengeluaran int
	var lama_join int

	fmt.Print("Input : ")
	fmt.Scan(&pengeluaran, &lama_join)

	if pengeluaran >= 5000 && lama_join >= 24 {
		fmt.Println("VVIP Backstage")
	} else if pengeluaran >= 2000 && lama_join >= 12 {
		fmt.Println("VIP Souncheck")
	} else if pengeluaran >= 500 {
		fmt.Println("Festival Ground")
	} else {
		fmt.Println("Menonton dari YouTube")
	}
}