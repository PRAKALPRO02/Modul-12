package main

import (
	"fmt"
)

// Fungsi untuk mengurutkan array secara ascending menggunakan selection sort
func selectionSortAsc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] { // Cari elemen terkecil
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i] // Tukar elemen
	}
}

// Fungsi untuk mengurutkan array secara descending menggunakan selection sort
func selectionSortDesc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[maxIdx] { // Cari elemen terbesar
				maxIdx = j
			}
		}
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i] // Tukar elemen
	}
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah daerah (n): ")
	fmt.Scan(&n)

	if n <= 0 || n >= 1000 {
		fmt.Println("n harus lebih besar dari 0 dan kurang dari 1000.")
		return
	}

	for i := 0; i < n; i++ {
		var m int
		fmt.Printf("\nMasukkan jumlah rumah untuk daerah ke-%d: ", i+1)
		fmt.Scan(&m)

		if m <= 0 {
			fmt.Println("Jumlah rumah harus lebih besar dari 0.")
			return
		}

		// Membaca array angka sebagai input
		fmt.Printf("Masukkan nomor rumah untuk daerah ke-%d:\n", i+1)
		houses := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&houses[j])
		}

		// Pisahkan bilangan ganjil dan genap
		var ganjil []int
		var genap []int
		for _, num := range houses {
			if num%2 == 0 {
				genap = append(genap, num)
			} else {
				ganjil = append(ganjil, num)
			}
		}

		// Urutkan ganjil (ascending) dan genap (descending)
		selectionSortAsc(ganjil)
		selectionSortDesc(genap)

		// Cetak hasil
		fmt.Printf("Hasil urutan rumah untuk daerah ke-%d: ", i+1)
		for _, num := range ganjil {
			fmt.Printf("%d ", num)
		}
		for _, num := range genap {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}