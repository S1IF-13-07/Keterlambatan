package main

import (
    "fmt"
    "math"
)

func main () {
    var r float64

    fmt.Print("Masukkan jari jari lingkaran: ")
    fmt.Scan(&r)

    luas := math.Pi * r * r

    fmt.Printf("Luas lingkaran  %.1f\n", luas)
}
