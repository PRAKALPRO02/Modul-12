package main

import (
	"fmt"
)

func hitungMedian(arr []int) float64 {
	n := len(arr)
	if n == 0 {
		return 0 // Menghindari pembagian dengan nol
	}
	if n%2 == 1 {
		return float64(arr[n/2])
	}
	return float64(arr[n/2-1]+arr[n/2]) / 2.0
}

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}

func main() {
	var data []int
	var input int

	fmt.Println("Masukkan angka (0 untuk menghitung median, -322 untuk keluar):")

	for {
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Silahkan masukkan angka.")
			// Clear the input buffer
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		if input == -322 {
			break
		}

		if input == 0 {
			if len(data) == 0 {
				fmt.Println("data kosong, tidak dapat menghitung median.")
				continue
			}
			insertionSort(data)
			median := hitungMedian(data)
			fmt.Printf("Median: %.2f\n", median)
		} else {
			data = append(data, input)
		}
	}
}