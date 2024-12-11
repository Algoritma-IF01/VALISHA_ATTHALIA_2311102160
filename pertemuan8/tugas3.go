package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Membaca input
	scanner.Scan()
	inputLine := scanner.Text()
	parts := strings.Fields(inputLine)

	// Parsing input ke dalam array bilangan bulat
	var numbers []int
	for _, part := range parts {
		num, _ := strconv.Atoi(part)
		if num < 0 {
			break // Menghentikan pembacaan saat menemukan bilangan negatif
		}
		numbers = append(numbers, num)
	}

	// Urutkan array menggunakan insertion sort
	insertionSort(numbers)

	// Cetak array terurut
	for i, num := range numbers {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(num)
	}
	fmt.Println()

	// Periksa apakah data memiliki jarak tetap
	if isConsistentDifference(numbers) {
		fmt.Printf("Data berjarak %d\n", numbers[1]-numbers[0])
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}

// Fungsi untuk mengurutkan array menggunakan insertion sort
func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1
		// Pindahkan elemen yang lebih besar dari key ke depan
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// Fungsi untuk memeriksa apakah array memiliki jarak tetap
func isConsistentDifference(arr []int) bool {
	if len(arr) < 2 {
		return true // Jika hanya ada satu elemen, dianggap konsisten
	}

	difference := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != difference {
			return false
		}
	}
	return true
}
