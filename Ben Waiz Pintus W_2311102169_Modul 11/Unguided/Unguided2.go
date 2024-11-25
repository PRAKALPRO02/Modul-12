//Ben Waiz Pintus W
//2311102169

package main

import (
	"fmt"
)

// Fungsi untuk mengurutkan array menggunakan insertion sort
func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
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
	n := len(arr)
	if n%2 == 1 {
		// Jika jumlah elemen ganjil, median adalah elemen tengah
		return arr[n/2]
	}
	// Jika jumlah elemen genap, median adalah rata-rata dua elemen tengah
	return (arr[n/2-1] + arr[n/2]) / 2
}

func main() {
	var input int
	data := []int{} // Array untuk menyimpan data yang valid

	fmt.Println("Masukkan bilangan bulat (akhiri dengan -5313):")

	for {
		fmt.Scan(&input)

		if input == -5313 {
			// Marker untuk mengakhiri program
			break
		} else if input == 0 {
			// Jika menemukan angka 0, urutkan array dan cetak median
			insertionSort(data)
			if len(data) > 0 {
				median := getMedian(data)
				fmt.Println("Median saat ini:", median)
			}
		} else if input > 0 {
			// Hanya tambahkan bilangan bulat positif ke array
			data = append(data, input)
		}
	}
}
