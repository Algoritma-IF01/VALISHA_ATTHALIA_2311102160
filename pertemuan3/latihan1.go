package main

import "fmt"

func main() {
	var beratParsel int
	fmt.Print("Berat parsel (gram): ")
	fmt.Scanln(&beratParsel)

	kg := beratParsel / 1000
	gram := beratParsel % 1000

	biayaKg := kg * 10000
	biayaGram := 0
	if gram >= 500 {
		biayaGram = gram * 5
	} else if gram > 0 && gram < 500 {
		biayaGram = gram * 15
	}

	totalBiaya := biayaKg + biayaGram
	if kg > 10 && gram > 0 && gram < 500 {
		totalBiaya = totalBiaya - (gram * 15)
	}

	fmt.Println("Detail berat:", kg, "kg +", gram, "gr")
	fmt.Println("Detail biaya: Rp.", biayaKg, "+ Rp.", biayaGram)
	fmt.Println("Total biaya: Rp.", totalBiaya)
}
