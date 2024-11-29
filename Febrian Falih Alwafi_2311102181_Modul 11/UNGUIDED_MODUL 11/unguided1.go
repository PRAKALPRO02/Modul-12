package main

import (
	"fmt"
	"strings"
)


func sortAscending(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}


func sortDescending(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}


func processHouses_2311102181(houses []int) (string, string) {
	var ganjil, genap []int

	for _, num := range houses {
		if num%2 == 0 {
			genap = append(genap, num)
		} else {
			ganjil = append(ganjil, num)
		}
	}

	sortDescending(ganjil)
	sortAscending(genap)

	ganjilStr := strings.Trim(fmt.Sprint(ganjil), "[]")
	genapStr := strings.Trim(fmt.Sprint(genap), "[]")

	return ganjilStr, genapStr
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah daerah (n): ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Jumlah daerah harus lebih besar dari 0.")
		return
	}

	results := make([]string, 0, n)

	for i := 1; i <= n; i++ {
		var m int
		fmt.Printf("\nMasukkan jumlah rumah kerabat untuk daerah ke-%d: ", i)
		fmt.Scan(&m)

		if m <= 0 {
			fmt.Println("Jumlah rumah kerabat harus lebih besar dari 0.")
			continue
		}

		fmt.Printf("Masukkan nomor rumah kerabat untuk daerah ke-%d: ", i)
		houses := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&houses[j])
		}

		ganjilStr, genapStr := processHouses_2311102181(houses)
		results = append(results, fmt.Sprintf("%s\n%s", ganjilStr, genapStr))
	}

	fmt.Println("\nHasil pengurutan rumah kerabat:")
	for i, result := range results {
		fmt.Printf("\nDaerah %d:\n%s\n", i+1, result)
	}
}
