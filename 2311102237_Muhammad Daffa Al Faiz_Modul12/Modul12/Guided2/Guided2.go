package main

import (
	"fmt"
	"math"
)

// Fungsi insertion sort untuk mengurutkan array
func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		// Geser elemen yang lebih besar dari key ke kanan
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}
// Fungsi untuk memeriksa apakah data berjarak tetap
func isDataConsistentlySpaced(arr []int) (bool, int) {
	if len(arr) < 2 {
		return true, 0 // Array dengan kurang dari 2 elemen dianggap berjarak tetap
	}

	// Hitung selisih awal
	diff := int(math.Abs(float64(arr[1] - arr[0])))

	for i := 1; i < len(arr)-1; i++ {
		currentDiff := int(math.Abs(float64(arr[i+1] - arr[i])))
		if currentDiff != diff {
			return false, 0 // Jika ada selisih yang berbeda, tidak berjarak tetap
		}
	}

	return true, diff
}
func main() {
	var data []int
	var input int

	fmt.Println("Masukkan data (akhiri dengan bilangan negatif):")
	for {
		fmt.Scan(&input)
		if input < 0 {
			break
		}
		data = append(data, input)
	}

	// Urutkan data menggunakan insertion sort
	insertionSort(data)

	// Periksa apakah data berjarak tetap
	isConsistent, diff := isDataConsistentlySpaced(data)

	// Cetak hasil
	fmt.Println("Hasil pengurutan:", data)
	if isConsistent {
		fmt.Printf("Data berjarak %d\n", diff)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}