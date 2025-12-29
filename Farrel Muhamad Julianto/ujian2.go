package main

import"fmt"

func main() {

 var N int
 fmt.Printf("Jumlah Piring : ")
 fmt.Scan(&N)

 totalYes := 0

 for i := 0; i < N; i++ {
	var skor int
	fmt.Printf("Skor : ")
	fmt.Scan(&skor)
 

	if skor > 75 {
		fmt.Println("Juri bilang YES")
		totalYes++
	} else {
		fmt.Println("Juri bilang NO")
	}
 }

 fmt.Printf("Total YES: %d\n", totalYes)
}