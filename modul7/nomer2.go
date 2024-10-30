package main

import (
	"fmt"
)

func isPerfectNumber(x int) bool {
	sum := 0
	for i := 1; i < x; i++ {
		if x%i == 0 {
			sum += i
		}
	}
	return sum == x
}

func perfectNumbersInRange(a int, b int) {
	fmt.Printf("Bilangan perfect number antara %d dan %d adalah:\n", a, b)
	found := false
	for i := a; i <= b; i++ {
		if isPerfectNumber(i) {
			fmt.Printf("%d ", i)
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ada perfect number dalam rentang ini.")
	} else {
		fmt.Println()
	}
}

func main() {
	var a, b int
	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)

	if a > b {
		fmt.Println("Error: a harus lebih kecil atau sama dengan b.")
		return
	}

	perfectNumbersInRange(a, b)
}
