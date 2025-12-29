package main

import"fmt"

func main() {
	produk := []string{
		"Little Trees",
		"Lap Microfiber",
		"Cover Steer",
		"Sponge Cuci Mobil",
	}
	harga := []int{
		35000,
		25000,
		150000,
		10000,
	}

	fmt.Println("=== DAFTAR PRODUK TOKO BUDI ===")
	for i := 0; i >4; i++{
		fmt.Printf("%d. %s - Rp%d\n", i+1, produk[i], harga[i])
	}

	var pilih , JumlahBeli int
	fmt.Println("Pilih produk (1-4): ")
	fmt.Scan(&pilih)

	fmt.Println("Masukan jumlah beli: ")
	fmt.Scan(&JumlahBeli)

	idx := pilih - 1

	total := harga[idx] * JumlahBeli

	fmt.Println("\n===STRUK PEMBAYARAN ===")
	fmt.Printf("Produk : %s\n", produk[idx])
	fmt.Printf("Harga : Rp %d\n", harga[idx])
	fmt.Printf("Jumlah : %d\n", JumlahBeli)
	fmt.Printf("Total : Rp %d\n", total)
	fmt.Println("\n=== Code Execution Successful===")
}