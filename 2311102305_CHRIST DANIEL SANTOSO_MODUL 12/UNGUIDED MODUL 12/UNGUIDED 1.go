package main

import (
	"fmt"
	"strings"
)

func sortAscending(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
}

func sortDescending(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		maxIndex := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > arr[maxIndex] {
				maxIndex = j
			}
		}
		arr[i], arr[maxIndex] = arr[maxIndex], arr[i]
	}
}

func main() {
	var numRegions int
	fmt.Print("Masukkan jumlah daerah: ")
	fmt.Scan(&numRegions)

	if numRegions <= 0 {
		fmt.Println("Jumlah daerah harus lebih besar dari 0.")
		return
	}

	results := make([]string, 0, numRegions)

	for region := 1; region <= numRegions; region++ {
		var numHouses int
		fmt.Printf("\nMasukkan jumlah rumah kerabat untuk daerah ke-%d: ", region)
		fmt.Scan(&numHouses)

		if numHouses <= 0 {
			fmt.Println("Jumlah rumah kerabat harus lebih besar dari 0.")
			continue
		}

		fmt.Printf("Masukkan nomor rumah kerabat untuk daerah ke-%d: ", region)
		houses := make([]int, numHouses)
		for i := 0; i < numHouses; i++ {
			fmt.Scan(&houses[i])
		}

		var oddNumbers, evenNumbers []int
		for _, house := range houses {
			if house%2 == 0 {
				evenNumbers = append(evenNumbers, house)
			} else {
				oddNumbers = append(oddNumbers, house)
			}
		}

		sortDescending(oddNumbers)
		sortAscending(evenNumbers)

		oddString := strings.Trim(fmt.Sprint(oddNumbers), "[]")
		evenString := strings.Trim(fmt.Sprint(evenNumbers), "[]")

		results = append(results, fmt.Sprintf("%s\n%s", oddString, evenString))
	}

	fmt.Println("\nHasil pengurutan rumah kerabat:")
	for i, result := range results {
		fmt.Printf("\nDaerah %d:\n%s\n", i+1, result)
	}
}
