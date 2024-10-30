package main

import (
	"fmt"
)

func hitungPertemuan(hari int, x int, y int, count int) int {
	if hari > 365 {
		return count
	}

	if hari%x == 0 && hari%y != 0 {
		count++
	}
	return hitungPertemuan(hari+1, x, y, count)
}

func main() {
	var x, y int

	fmt.Print("Masukkan nilai x (kelipatan pertemuan): ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y (kelipatan yang dilarang): ")
	fmt.Scan(&y)

	jumlahPertemuan := hitungPertemuan(1, x, y, 0)

	fmt.Printf("Jumlah pertemuan rahasia dalam setahun: %d\n", jumlahPertemuan)
}
