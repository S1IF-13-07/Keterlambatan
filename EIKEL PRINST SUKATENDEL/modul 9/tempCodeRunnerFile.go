package main

import "fmt"

func main() {
	var x int

	fmt.Print("Masukkan bilangan x: ")
	fmt.Scan(&x)

	if x%2 == 0 && x < 0 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}