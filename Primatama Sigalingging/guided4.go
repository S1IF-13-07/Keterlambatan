package main

import "fmt"

func main() {
    var f float64

    fmt.Print("Masukkan suhu dalam Fahrenheit: ")
    fmt.Scan(&f)

    c := (f - 32) * 5 / 9
    fmt.Printf("Suhu dalam Celcius = %.0f\n", c)

}
