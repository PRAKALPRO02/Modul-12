package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah daerah (n): ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Jumlah daerah harus lebih besar dari 0.")
		return
	}

	for i := 0; i < n; i++ {
		var m int
		fmt.Printf("Masukkan jumlah rumah kerabat untuk daerah ke-%d: ", i+1)
		fmt.Scan(&m)

		if m <= 0 {
			fmt.Println("Jumlah rumah harus lebih besar dari 0.")
			return
		}

		// Input nomor rumah
		houses := make([]int, m)
		fmt.Printf("Masukkan nomor rumah kerabat untuk daerah ke-%d: ", i+1)
		for j := 0; j < m; j++ {
			fmt.Scan(&houses[j])
		}

		// Pisahkan bilangan ganjil dan genap
		var oddNumbers, evenNumbers []int
		for _, house := range houses {
			if house%2 == 0 {
				evenNumbers = append(evenNumbers, house)
			} else {
				oddNumbers = append(oddNumbers, house)
			}
		}

		// Urutkan ganjil secara menaik dan genap secara menurun
		sort.Ints(oddNumbers)
		sort.Sort(sort.Reverse(sort.IntSlice(evenNumbers)))

		// Cetak hasil
		fmt.Printf("Hasil untuk daerah ke-%d: ", i+1)
		for _, odd := range oddNumbers {
			fmt.Printf("%d ", odd)
		}
		for _, even := range evenNumbers {
			fmt.Printf("%d ", even)
		}
		fmt.Println()
	}
}
