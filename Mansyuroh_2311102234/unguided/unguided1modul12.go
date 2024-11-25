package main

import (
	"fmt"
)

// Fungsi untuk mengurutkan array secara ascending menggunakan selection sort
func selectionSort(arr []int, ascending bool) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		selectedIdx := i
		for j := i + 1; j < n; j++ {
			if (ascending && arr[j] < arr[selectedIdx]) || (!ascending && arr[j] > arr[selectedIdx]) {
				selectedIdx = j
			}
		}
		// Tukar elemen
		arr[i], arr[selectedIdx] = arr[selectedIdx], arr[i]
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
		ganjil := make([]int, 0)
		genap := make([]int, 0)
		for _, num := range houses {
			if num%2 == 0 {
				genap = append(genap, num)
			} else {
				ganjil = append(ganjil, num)
			}
		}

		// Urutkan ganjil (ascending) dan genap (ascending)
		selectionSort(ganjil, true) // Urutkan ascending (Ganjil)
		selectionSort(genap, true)  // Urutkan ascending (Genap)

		campuran := append(ganjil, genap...)

		// Cetak hasil
		fmt.Printf("Hasil urutan rumah untuk daerah ke-%d:\n", i+1)

		// Baris pertama: Ganjil saja
		fmt.Print("Ganjil terurut: ")
		for _, num := range ganjil {
			fmt.Printf("%d ", num)
		}
		fmt.Println()

		fmt.Print("Campuran terurut: ")
		for _, num := range campuran {
			fmt.Printf("%d ", num)
		}
		fmt.Println()

		fmt.Print("Genap terurut: ")
		for _, num := range genap {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}
