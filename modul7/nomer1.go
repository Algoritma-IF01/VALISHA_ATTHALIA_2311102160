package main

import (
	"fmt"
)

func hitungBiayaSewa(jam int, menit int, isMember bool) int {
	tarifPerJamMember := 3500
	tarifPerJamNonMember := 5000

	var durasiJam int
	if jam == 0 && menit > 0 {
		durasiJam = 1
	} else if menit >= 10 {
		durasiJam = jam + 1
	} else {
		durasiJam = jam
	}

	var tarifPerJam int
	if isMember {
		tarifPerJam = tarifPerJamMember
	} else {
		tarifPerJam = tarifPerJamNonMember
	}

	biayaAwal := durasiJam * tarifPerJam

	if durasiJam > 3 {
		biayaAkhir := int(float64(biayaAwal) * 0.9) // diskon 10%
		return biayaAkhir
	}
	return biayaAwal
}

func main() {
	var jam, menit int
	var isMemberInput string
	var isMember bool

	fmt.Print("Masukkan durasi (jam): ")
	fmt.Scan(&jam)

	fmt.Print("Masukkan durasi (menit): ")
	fmt.Scan(&menit)

	fmt.Print("Apakah Anda member? (y/n): ")
	fmt.Scan(&isMemberInput)

	isMember = isMemberInput == "y"

	biaya := hitungBiayaSewa(jam, menit, isMember)

	fmt.Printf("Biaya sewa: Rp %d\n", biaya)
}
