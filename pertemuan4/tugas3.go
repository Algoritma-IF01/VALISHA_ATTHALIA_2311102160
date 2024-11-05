package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// untuk menghitung total soal yang diselesaikan dan total skor
func hitungSkor(input string) (int, int) {
	parts := strings.Fields(input)
	waktuPengerjaan := parts[1:]
	totalSoal := 0
	totalWaktu := 0

	// Iterasi waktu pengerjaan untuk setiap soal
	for _, waktu := range waktuPengerjaan {
		var waktuInt int
		fmt.Sscanf(waktu, "%d", &waktuInt)

		if waktuInt <= 300 {
			totalSoal++
			totalWaktu += waktuInt
		}
	}
	return totalSoal, totalWaktu
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var pemenang string
	var maxSoal, minWaktu int
	minWaktu = 1000000

	for {
		// Input dari user
		fmt.Println("\nMasukkan data nama peserta dan waktu : ")
		scanner.Scan()
		input := scanner.Text()

		// Cek jika user sudah selesai memasukkan data
		if strings.ToLower(input) == "selesai" {
			break
		}

		// Hitung skor berdasarkan input peserta
		soal, waktu := hitungSkor(input)

		// Tampilkan hasil untuk peserta yang sedang diinput
		fmt.Println("Nama:", strings.Fields(input)[0], "Soal diselesaikan:", soal, "Total waktu:", waktu)

		// jika soal yang diselesaikan sama, yang memiliki waktu lebih sedikit yang menang
		if soal > maxSoal || (soal == maxSoal && waktu < minWaktu) {
			pemenang = strings.Fields(input)[0]
			maxSoal = soal
			minWaktu = waktu
		}
	}

	// Cetak pemenang setelah semua data dimasukkan
	if pemenang != "" {
		fmt.Printf("Pemenang: %s %d soal, total waktu: %d menit\n", pemenang, maxSoal, minWaktu)
	} else {
		fmt.Println("Tidak ada peserta yang valid.")
	}
}
