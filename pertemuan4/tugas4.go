package main

import (
	"fmt"
)

// Fungsi cetakDeret untuk mencetak deret sesuai aturan yang diberikan
func cetakDeret(n int) {
	for n != 1 {
		fmt.Print(n, " ")

		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	fmt.Println(1)
}

func main() {
	var n int

	// input dari pengguna
	fmt.Print("Masukkan angka: ")
	fmt.Scanln(&n)

	// memastikan bahwa input valid (kurang dari 1 juta)
	if n <= 0 || n >= 1000000 {
		fmt.Println("Input harus bilangan positif dan kurang dari 1 juta!")
		return
	}
	cetakDeret(n)
}
