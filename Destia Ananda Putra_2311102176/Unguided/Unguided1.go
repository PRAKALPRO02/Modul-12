package main

import (
	"fmt"
	"strings"
)

func selectionSortAsc(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func selectionSortDesc(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		maxIdx := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > arr[maxIdx] {
				maxIdx = j
			}
		}
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
	}
}

func splitOddEven(houses []int) ([]int, []int) {
	var ganjil, genap []int
	for _, num := range houses {
		if num%2 == 0 {
			genap = append(genap, num)
		} else {
			ganjil = append(ganjil, num)
		}
	}
	return ganjil, genap
}

func inputData(n int) []string {
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

		ganjil, genap := splitOddEven(houses)

		selectionSortDesc(ganjil)
		selectionSortAsc(genap)

		ganjilStr := strings.Trim(fmt.Sprint(ganjil), "[]")
		genapStr := strings.Trim(fmt.Sprint(genap), "[]")

		results = append(results, fmt.Sprintf("%s\n%s", ganjilStr, genapStr))
	}
	return results
}

func printResults(results []string) {
	fmt.Println("\nHasil pengurutan rumah kerabat:")
	for i, result := range results {
		fmt.Printf("\nDaerah %d:\n%s\n", i+1, result)
	}
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah daerah (n): ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Jumlah daerah harus lebih besar dari 0.")
		return
	}

	results := inputData(n)

	printResults(results)
}
