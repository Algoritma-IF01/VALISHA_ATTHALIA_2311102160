package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64

	// Input
	fmt.Print("Jejari = ")
	fmt.Scanln(&radius)

	// Perhitungan volume dan luas permukaan
	volume := (4.0 / 3.0) * math.Pi * math.Pow(radius, 3)
	surfaceArea := 4 * math.Pi * math.Pow(radius, 2)

	// Output
	fmt.Printf("Bola dengan jejari %.0f memiliki volume %.4f dan luas kulit %.4f\n", radius, volume, surfaceArea)
}
