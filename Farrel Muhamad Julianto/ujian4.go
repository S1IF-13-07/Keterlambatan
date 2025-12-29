package main

import "fmt"
func main() {
 var N int
 fmt.Scan(&N)

 asli := N

 jumlah := 0

 for N > 0 {
	digit := N % 10

	jumlah += digit

	N = N / 10
 }

 fmt.Printf("Masukan Angka : %d\n", asli)
 fmt.Printf("Hasil Penjumlahan : %d\n", jumlah)
}