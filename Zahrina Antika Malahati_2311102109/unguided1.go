package main

import "fmt"

func selectionSort(arr []int, ascending bool) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		targetIdx := i
		for j := i + 1; j < n; j++ {
			if (ascending && arr[j] < arr[targetIdx]) || (!ascending && arr[j] > arr[targetIdx]) {
				targetIdx = j
			}
		}
		arr[i], arr[targetIdx] = arr[targetIdx], arr[i]
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
		fmt.Printf("Masukkan jumlah rumah kerabat untuk daerah ke-%d: ", i+1)
		fmt.Scan(&m)

		if m <= 0 || m >= 1000000 {
			fmt.Println("m harus lebih besar dari 0 dan kurang dari 1000000.")
			return
		}

		houses := make([]int, m)
		fmt.Printf("Masukkan nomor rumah kerabat untuk daerah ke-%d: ", i+1)
		for j := 0; j < m; j++ {
			fmt.Scan(&houses[j])
		}

		var oddHouses []int
		var evenHouses []int

		for _, house := range houses {
			if house%2 == 1 {
				oddHouses = append(oddHouses, house) // Nomor ganjil
			} else {
				evenHouses = append(evenHouses, house) // Nomor genap
			}
		}

		// Urutkan nomor rumah ganjil membesar
		selectionSort(oddHouses, true)

		// Urutkan nomor rumah genap mengecil
		selectionSort(evenHouses, false)

		// Cetak hasil
		fmt.Printf("Hasil urutan rumah untuk daerah ke-%d: ", i+1)

		// Cetak nomor rumah ganjil
		for _, house := range oddHouses {
			fmt.Printf("%d ", house)
		}

		// Cetak nomor rumah genap
		for _, house := range evenHouses {
			fmt.Printf("%d ", house)
		}

		fmt.Println()
	}
}

// Zahrina Antika Malahati_2311102109
