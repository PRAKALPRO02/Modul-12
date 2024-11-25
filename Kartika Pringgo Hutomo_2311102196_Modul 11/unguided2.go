package main

import (
	"fmt"
)

// Fungsi untuk mengurutkan array menggunakan insertion sort
func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		// Pindahkan elemen yang lebih besar dari key ke satu posisi di depan
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

// Fungsi untuk menghitung median dari array yang sudah terurut
func getMedian(arr []int) int {
	if len(arr) == 0 {
		return 0 // Jika array kosong, median adalah 0
	}
	if len(arr)%2 == 1 {
		// Jika jumlah elemen ganjil, median adalah elemen tengah
		return arr[len(arr)/2]
	}
	// Jika jumlah elemen genap, median adalah rata-rata dua elemen tengah
	return (arr[len(arr)/2-1] + arr[len(arr)/2]) / 2
}

func main() {
	var input int
	data := make([]int, 0) // Inisialisasi array kosong

	fmt.Println("Masukkan bilangan bulat (akhiri dengan -5313):")

	for {
		fmt.Scan(&input)

		if input == -5313 {
			// Marker untuk mengakhiri program
			break
		}

		switch {
		case input == 0:
			// Jika menemukan angka 0, urutkan array dan cetak median
			insertionSort(data)
			fmt.Println("Median saat ini:", getMedian(data))
		case input > 0:
			// Hanya tambahkan bilangan bulat positif ke array
			data = append(data, input)
		}
	}
}
