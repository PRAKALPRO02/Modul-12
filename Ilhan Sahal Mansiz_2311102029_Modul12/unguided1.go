package main

import (
	"fmt"
	"sort"
)

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

		// Masukkan nomor rumah
		houses := make([]int, m)
		fmt.Printf("Masukkan nomor rumah kerabat untuk daerah ke-%d: ", i+1)
		for j := 0; j < m; j++ {
			fmt.Scan(&houses[j])
		}

		// Pisahkan nomor ganjil dan genap
		var odd, even []int
		for _, house := range houses {
			if house%2 == 0 {
				even = append(even, house)
			} else {
				odd = append(odd, house)
			}
		}

		// Urutkan ganjil membesar
		sort.Ints(odd)
		// Urutkan genap mengecil
		sort.Sort(sort.Reverse(sort.IntSlice(even)))

		// Cetak hasil
		fmt.Printf("Hasil urutan rumah untuk daerah ke-%d: ", i+1)
		for _, house := range odd {
			fmt.Printf("%d ", house)
		}
		for _, house := range even {
			fmt.Printf("%d ", house)
		}
		fmt.Println()
	}
}