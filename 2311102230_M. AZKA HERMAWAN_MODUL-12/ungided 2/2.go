package main

import "fmt"

func hitungMedian(data []int, panjang int) float64 {
	if panjang%2 == 1 {
		return float64(data[panjang/2])
	}
	return float64(data[panjang/2-1]+data[panjang/2]) / 2.0
}

func SelectionSort(data []int, panjang int) {
	for i := 0; i < panjang-1; i++ {
		minimal := i
		for j := i; j < panjang; j++ {
			if data[j] < data[minimal] {
				minimal = j
			}
		}
		data[i], data[minimal] = data[minimal], data[i]
	}
}

func main() {
	var input int
	var data [1000000]int
	var semuaMedian [1000000]float64
	panjang := 0
	totalMedian := 0

	fmt.Println("Masukan bilangan : ")

	for {
		fmt.Scan(&input)

		if input == -5313 {
			break
		}

		if input == 0 {
			if panjang > 0 {
				SelectionSort(data[:], panjang)

				median := hitungMedian(data[:], panjang)
				semuaMedian[totalMedian] = median
				totalMedian++
			}
		} else {
			data[panjang] = input
			panjang++
		}
	}

	fmt.Println("\nHasil median: ")
	for i := 0; i < totalMedian; i++ {
		fmt.Printf("Median %d: %.0f\n", i+1, semuaMedian[i])
	}
	fmt.Println()
}
