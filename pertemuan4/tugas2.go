package main

import (
	"fmt"
	"math"
)

// untuk menghitung jarak antara dua titik
func jarak(x1, y1, x2, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}

// untuk menentukan apakah titik berada di dalam lingkaran
func diDalamLingkaran(x, y, cx, cy, radius float64) bool {
	return jarak(x, y, cx, cy) <= radius
}

func main() {
	var x1, y1, r1 float64
	var x2, y2, r2 float64
	var x, y float64

	fmt.Scan(&x1, &y1, &r1)
	fmt.Scan(&x2, &y2, &r2)
	fmt.Scan(&x, &y)

	// Mengecek posisi titik terhadap lingkaran 1 dan 2
	diDalam1 := diDalamLingkaran(x, y, x1, y1, r1)
	diDalam2 := diDalamLingkaran(x, y, x2, y2, r2)

	// Menentukan output berdasarkan posisi titik
	if diDalam1 && diDalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diDalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diDalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
