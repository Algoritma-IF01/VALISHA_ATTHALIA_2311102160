package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen array (N): ")
	fmt.Scanln(&n)

	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i)
		fmt.Scanln(&array[i])
	}

	var choice, x, index int
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Menampilkan keseluruhan isi dari array.")
		fmt.Println("2. Menampilkan elemen-elemen array dengan indeks ganjil saja.")
		fmt.Println("3. Menampilkan elemen-elemen array dengan indeks genap saja.")
		fmt.Println("4. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x.")
		fmt.Println("5. Menghapus elemen array pada indeks tertentu.")
		fmt.Println("6. Menampilkan rata-rata dari bilangan yang ada di dalam array.")
		fmt.Println("7. Menampilkan standar deviasi dari bilangan yang ada di dalam array.")
		fmt.Println("8. Menampilkan frekuensi dari setiap bilangan tertentu di dalam array.")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih opsi: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Println("Isi array:", array)

		case 2:
			fmt.Print("Elemen dengan indeks ganjil: ")
			for i := 1; i < len(array); i += 2 {
				fmt.Print(array[i], " ")
			}
			fmt.Println()

		case 3:
			fmt.Print("Elemen dengan indeks genap: ")
			for i := 0; i < len(array); i += 2 {
				fmt.Print(array[i], " ")
			}
			fmt.Println()

		case 4:
			fmt.Print("Masukkan nilai x: ")
			fmt.Scanln(&x)
			fmt.Printf("Elemen dengan indeks kelipatan %d: ", x)
			for i := x; i < len(array); i += x {
				fmt.Print(array[i], " ")
			}
			fmt.Println()

		case 5:
			fmt.Print("Masukkan indeks yang ingin dihapus: ")
			fmt.Scanln(&index)
			if index >= 0 && index < len(array) {
				array = append(array[:index], array[index+1:]...)
				fmt.Println("Isi array setelah penghapusan:", array)
			} else {
				fmt.Println("Indeks tidak valid!")
			}

		case 6:
			total := 0
			for _, value := range array {
				total += value
			}
			avg := float64(total) / float64(len(array))
			fmt.Printf("Rata-rata dari array: %.2f\n", avg)

		case 7:
			total := 0
			for _, value := range array {
				total += value
			}
			mean := float64(total) / float64(len(array))
			variance := 0.0
			for _, value := range array {
				variance += math.Pow(float64(value)-mean, 2)
			}
			stdDev := math.Sqrt(variance / float64(len(array)))
			fmt.Printf("Standar deviasi dari array: %.2f\n", stdDev)

		case 8:
			frequency := make(map[int]int)
			for _, value := range array {
				frequency[value]++
			}
			fmt.Println("Frekuensi setiap bilangan dalam array:")
			for num, freq := range frequency {
				fmt.Printf("%d: %d kali\n", num, freq)
			}

		case 0:
			fmt.Println("Terima kasih!")
			return

		default:
			fmt.Println("Pilihan tidak valid, silakan coba lagi.")
		}
	}
}
