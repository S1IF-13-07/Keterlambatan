package main

import "fmt"

func main() {
	var n int
	var s1, s2, temp int

	fmt.Print("Masukkan jumlah deret: ")
	fmt.Scan(&n)

	s1 = 0
	s2 = 1

	for i := 0; i < n; i++ {
		fmt.Print(s1, " ")
		temp = s1 + s2
		s1 = s2
		s2 = temp
	}
}
