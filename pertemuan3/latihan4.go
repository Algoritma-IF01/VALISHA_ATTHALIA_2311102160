package main

import (
	"fmt"
	"math"
)

func main() {
	var kantong1, kantong2 float64
	oleng := false

	for {
		// Input berat belanjaan untuk kedua kantong
		fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kantong1, &kantong2)

		// jika berat total melebihi 150 kg
		if kantong1+kantong2 > 150 {
			fmt.Println("Berat melebihi 150")
			fmt.Println("Proses selesai.")
			break
		}

		// kondisi berhenti jika berat negatif
		if kantong1 < 0 || kantong2 < 0 {
			fmt.Println("Proses selesai.")
			break
		}

		// selisih berat antara kedua kantong
		selisih := math.Abs(kantong1 - kantong2)
		oleng = selisih > 9

		// Output
		fmt.Printf("Sepeda motor Pak Andi akan oleng: %v\n", oleng)
	}
}
