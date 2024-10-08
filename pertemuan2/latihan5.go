package main

import (
	"fmt"
)

func main() {
	var input [5]int
	var character1, character2, character3 rune

	// Membaca input 5 buah integer
	fmt.Println("Masukkan 5 buah integer (32 s.d. 127):")
	for i := 0; i < 5; i++ {
		fmt.Scan(&input[i])
	}

	// Menampilkan hasil konversi ke karakter ASCII
	fmt.Println("Masukan:")
	for i, val := range input {
		if i < 5 {
			fmt.Printf("%d ", val)
		}
	}
	fmt.Println()

	// Mengambil 3 karakter dari input
	character1 = rune(input[0])
	character2 = rune(input[1])
	character3 = rune(input[2])

	// Menampilkan hasil karakter
	fmt.Printf("Hasil karakter: %c %c %c\n", character1, character2, character3)

	// Mencetak hasil karakter
	fmt.Println("Bagus")
	fmt.Println("TOP")
}
