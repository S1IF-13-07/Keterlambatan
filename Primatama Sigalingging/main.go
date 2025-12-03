package main

import "fmt"

func main() {
	var nama string
	fmt.Print("Masukkan nama kalian : ")
	fmt.Scanln(&nama)
	fmt.Println(nama)
}