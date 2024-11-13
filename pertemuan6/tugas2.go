package main

import (
	"fmt"
	"math"
)

func tampilArr(arr []int) {
	fmt.Println("Isi array:", arr)
}

func ganjil(arr []int) {
	fmt.Print("Elemen dengan nilai ganjil: ")
	for _, value := range arr {
		if value%2 != 0 {
			fmt.Print(value, " ")
		}
	}
	fmt.Println()
}

func genap(arr []int) {
	fmt.Print("Elemen dengan nilai genap: ")
	for _, value := range arr {
		if value%2 == 0 && value != 0 {
			fmt.Print(value, " ")
		}
	}
	fmt.Println()
}

func kelipatan(arr []int, x int) {
	fmt.Printf("Elemen pada indeks kelipatan %d: ", x)
	for i := x - 1; i < len(arr); i += x {
		if arr[i] != 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()
}

func hapus(arr []int, idx int) []int {
	return append(arr[:idx], arr[idx+1:]...)
}

func mean(arr []int) float64 {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return float64(sum) / float64(len(arr))
}

func deviasi(arr []int, rata float64) float64 {
	sumSquares := 0.0
	for _, value := range arr {
		sumSquares += math.Pow(float64(value)-rata, 2)
	}
	return math.Sqrt(sumSquares / float64(len(arr)))
}

func frekuensi(arr []int, target int) int {
	count := 0
	for _, value := range arr {
		if value == target {
			count++
		}
	}
	return count
}

func main() {
	var n, x, deleteIndex, target int

	fmt.Print("Masukkan jumlah elemen array (N): ")
	fmt.Scan(&n)

	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan nilai untuk elemen ke-%d: ", i)
		fmt.Scan(&array[i])
	}

	for {
		fmt.Println("\nPilih operasi:")
		fmt.Println("1. Tampilkan seluruh isi array")
		fmt.Println("2. Tampilkan elemen dengan nilai ganjil")
		fmt.Println("3. Tampilkan elemen dengan nilai genap")
		fmt.Println("4. Tampilkan elemen pada indeks kelipatan x")
		fmt.Println("5. Hapus elemen pada indeks tertentu")
		fmt.Println("6. Tampilkan rata-rata dari elemen array")
		fmt.Println("7. Tampilkan standar deviasi elemen array")
		fmt.Println("8. Tampilkan frekuensi dari suatu bilangan")
		fmt.Println("9. Keluar")

		var pilihan int
		fmt.Print("Masukkan pilihan: ")
		fmt.Scan(&pilihan)

		switch pilihan {
		case 1:
			tampilArr(array)

		case 2:
			ganjil(array)

		case 3:
			genap(array)

		case 4:
			fmt.Print("Masukkan nilai x untuk kelipatan indeks: ")
			fmt.Scan(&x)
			kelipatan(array, x)

		case 5:
			fmt.Print("Masukkan indeks untuk dihapus: ")
			fmt.Scan(&deleteIndex)
			if deleteIndex >= 0 && deleteIndex < len(array) {
				array = hapus(array, deleteIndex)
				tampilArr(array)
			} else {
				fmt.Println("Indeks tidak valid!")
			}

		case 6:
			rata := mean(array)
			fmt.Printf("Rata-rata elemen array: %.2f\n", rata)

		case 7:
			rata := mean(array)
			stdDev := deviasi(array, rata)
			fmt.Printf("Standar deviasi elemen array: %.2f\n", stdDev)

		case 8:
			fmt.Print("Masukkan nilai untuk mencari frekuensi: ")
			fmt.Scan(&target)
			frequency := frekuensi(array, target)
			fmt.Printf("Frekuensi %d dalam array: %d kali\n", target, frequency)

		case 9:
			fmt.Println("Keluar dari program.")
			return

		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}
