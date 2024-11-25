package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah daerah (n): ")
	_, err := fmt.Scan(&n)
	if err != nil || n <= 0 || n >= 1000 {
		fmt.Println("n harus lebih besar dari 0 dan kurang dari 1000.")
		return
	}

	for i := 0; i < n; i++ {
		var m int
		fmt.Printf("Masukkan jumlah rumah kerabat untuk daerah ke-%d: ", i+1)
		_, err := fmt.Scan(&m)
		if err != nil || m <= 0 || m >= 1000000 {
			fmt.Println("m harus lebih besar dari 0 dan kurang dari 1000000.")
			return
		}

		houses := make([]int, m)
		fmt.Printf("Masukkan nomor rumah kerabat untuk daerah ke-%d: ", i+1)
		for j := 0; j < m; j++ {
			_, err := fmt.Scan(&houses[j])
			if err != nil {
				fmt.Println("Input tidak valid. Silakan masukkan angka.")
				return
			}
		}

		var odd, even []int
		for _, house := range houses {
			if house%2 == 0 {
				even = append(even, house)
			} else {
				odd = append(odd, house)
			}
		}

		sort.Ints(odd)
		sort.Ints(even)

		result := append(odd, even...)
		fmt.Printf("Hasil urutan rumah untuk daerah ke-%d: %s\n", i+1, strings.Trim(strings.Join(strings.Fields(fmt.Sprint(result)), " "), "[]"))
	}
}