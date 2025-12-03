package main

import "fmt"

func main() {
    var nama, nim, kelas string

    fmt.Println("Masukkan nama:")
    fmt.Scanln(&nama)

    fmt.Println("Masukkan NIM:")
    fmt.Scanln(&nim)

    fmt.Println("Masukkan kelas:")
    fmt.Scanln(&kelas)

    fmt.Printf("Perkenalkan nama saya adalah %s, saya adalah salah satu mahasiswa Telkom University prodi S1-Informatika dari kelas %s dengan NIM %s.\n", nama, kelas, nim)
}
