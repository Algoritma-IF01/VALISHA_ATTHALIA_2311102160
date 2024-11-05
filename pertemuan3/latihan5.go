package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung f(k)
func f(k float64) float64 {
	numerator := math.Pow((4*k + 2), 2)
	denominator := (4*k + 1) * (4*k + 3)
	return numerator / denominator
}

// Fungsi untuk menghitung akar dari 2 dengan jumlah iterasi K
func sqrt2(k int) float64 {
	result := 1.0
	for i := 0; i <= k; i++ {
		result *= f(float64(i))
	}
	return result
}

func main() {
	var K int

	for i := 1; i <= 3; i++ {
		// Input nilai K
		fmt.Print("Nilai K = ")
		fmt.Scan(&K)

		// // Menghitung f(K)
		// fk := f(float64(K))
		// fmt.Printf("Nilai f(K) = %.10f\n", fk)

		// Menghitung aproksimasi sqrt(2) dengan K iterasi
		approxSqrt2 := sqrt2(K)
		fmt.Printf("Nilai akar 2 = %.10f\n\n", approxSqrt2)
	}

	fmt.Println("Proses selesai.")
}
