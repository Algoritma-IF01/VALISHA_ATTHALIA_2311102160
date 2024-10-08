package main

import (
	"fmt"
)

func main() {
	var celsius float64

	// Input suhu dalam derajat Celsius
	fmt.Print("Masukkan suhu dalam Celsius: ")
	fmt.Scanln(&celsius)

	// Konversi suhu
	fahrenheit := (celsius * 9 / 5) + 32
	reamur := celsius * 4 / 5
	kelvin := celsius + 273.15

	// Menampilkan hasil konversi
	fmt.Printf("Temperatur Celsius: %.2f\n", celsius)
	fmt.Printf("Derajat Fahrenheit: %.2f\n", fahrenheit)
	fmt.Printf("Derajat Reamur: %.2f\n", reamur)
	fmt.Printf("Derajat Kelvin: %.2f\n", kelvin)
}
