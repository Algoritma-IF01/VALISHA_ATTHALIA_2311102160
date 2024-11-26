package main

import (
	"fmt"
)

func selectionSortAsc(arr []int) []int {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
	return arr
}

func main() {
	var n int
	fmt.Scan(&n)

	results := make([][]int, n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)
		houses := make([]int, m)

		for j := 0; j < m; j++ {
			fmt.Scan(&houses[j])
		}

		results[i] = selectionSortAsc(houses)
	}

	for i := 0; i < len(results); i++ {
		sortedHouses := results[i]
		for j := 0; j < len(sortedHouses); j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(sortedHouses[j])
		}
		fmt.Println()
	}
}
