package main

import (
	"fmt"
	"sort"
)

// Fungsi untuk menghitung median dari array
func findMedian(data []int) int {
	n := len(data)
	if n == 0 {
		return 0
	}
	if n%2 == 1 { // Jika jumlah data ganjil
		return data[n/2]
	}
	// Jika jumlah data genap
	return (data[n/2-1] + data[n/2]) / 2
}

func main() {
	var numbers []int
	var input int

	fmt.Println("Masukkan data (akhiri dengan -5313):")

	for {
		fmt.Scan(&input)
		if input == -5313 {
			break
		}
		if input == 0 {
			// Jika menemukan 0, urutkan data dan hitung median
			sort.Ints(numbers)
			median := findMedian(numbers)
			fmt.Println(median)
		} else {
			// Tambahkan data ke array
			numbers = append(numbers, input)
		}
	}
}
