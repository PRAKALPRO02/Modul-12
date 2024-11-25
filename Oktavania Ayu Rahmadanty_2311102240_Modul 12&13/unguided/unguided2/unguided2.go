package main

import (
	"fmt"
)

func main() {
	var data []int // Menyimpan data
	for {
		var bilangan int
		fmt.Scan(&bilangan) // Membaca input bilangan

		// Jika bilangan adalah tanda akhir
		if bilangan == -5313 {
			break
		}

		// Jika bilangan adalah 0, hitung median
		if bilangan == 0 {
			// Urutkan data dengan insertion sort
			insertionSort(data)

			// Hitung dan cetak median
			median := hitungMedian(data)
			fmt.Println(median)
		} else {
			// Tambahkan bilangan ke array data
			data = append(data, bilangan)
		}
	}
}

func insertionSort(data []int) {
	// Algoritma insertion sort untuk pengurutan array
	for i := 1; i < len(data); i++ {
		key := data[i]
		j := i - 1

		// Geser elemen yang lebih besar dari key
		for j >= 0 && data[j] > key {
			data[j+1] = data[j]
			j--
		}

		// Tempatkan key pada posisi yang benar
		data[j+1] = key
	}
}

func hitungMedian(data []int) int {
	n := len(data)
	if n == 0 {
		return 0 // Tidak ada data
	}

	if n%2 == 1 {
		// Jika jumlah elemen ganjil
		return data[n/2]
	} else {
		// Jika jumlah elemen genap
		return (data[(n/2)-1] + data[n/2]) / 2 // Dibulatkan ke bawah
	}
}
