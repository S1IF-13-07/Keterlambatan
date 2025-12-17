package main

import "fmt"

func main() {
	var token string

	fmt.Print("Masukkan token: ")
	fmt.Scan(&token)

	for token != "12345abcde" {
		fmt.Print("Token salah, masukkan ulang: ")
		fmt.Scan(&token)
	}

	fmt.Println("Congratulations you have successfully logged in")
}
