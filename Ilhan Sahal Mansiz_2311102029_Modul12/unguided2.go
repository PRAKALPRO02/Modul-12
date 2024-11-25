package main

import (
	"fmt"
)

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

func findMedian(arr []int) int {
	n := len(arr)

	// Jika jumlah elemen ganjil, ambil nilai tengah
	if n%2 == 1 {
		return arr[n/2]
	}

	// Jika jumlah elemen genap, ambil rerata kedua nilai tengah (dibulatkan ke bawah)
	return (arr[(n/2)-1] + arr[n/2]) / 2
}

func main() {
	var data []int

	for {
		var num int
		fmt.Scan(&num)

		if num == -5313 { // Marker untuk menghentikan input
			break
		}

		if num == 0 { // Jika 0 ditemukan, hitung median
			// Urutkan data menggunakan insertion sort
			insertionSort(data)

			// Cari median dan cetak hasil
			median := findMedian(data)
			fmt.Println(median)
		} else {
			// Tambahkan data ke array
			data = append(data, num)
		}
	}
}