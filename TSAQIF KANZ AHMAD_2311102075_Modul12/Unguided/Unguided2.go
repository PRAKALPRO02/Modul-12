package main

import (
	"fmt"
)

func selectionSort(array []int) {
	n := len(array)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if array[j] < array[minIdx] {
				minIdx = j
			}
		}
		array[i], array[minIdx] = array[minIdx], array[i]
	}
}

func hitungMedian(array []int) float64 {
	jumlah := len(array)
	if jumlah == 0 {
		return 0
	}

	selectionSort(array)

	if jumlah%2 == 1 {
		return float64(array[jumlah/2])
	}
	tengah1 := array[(jumlah/2)-1]
	tengah2 := array[jumlah/2]
	return float64(tengah1+tengah2) / 2.0
}

func main() {
	var jumlahData int
	fmt.Print("Masukkan jumlah elemen data: ")
	fmt.Scan(&jumlahData)

	if jumlahData <= 0 || jumlahData > 1000000 {
		fmt.Println("Jumlah elemen harus lebih besar dari 0 dan tidak boleh lebih dari 1.000.000.")
		return
	}

	data := make([]int, jumlahData)
	fmt.Printf("Masukkan %d elemen data:\n", jumlahData)
	for i := 0; i < jumlahData; i++ {
		fmt.Scan(&data[i])
	}
	median := hitungMedian(data)
	fmt.Printf("Median dari data adalah: %.2f\n", median)
}
