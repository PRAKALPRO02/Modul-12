package main

import (
	"fmt"
)

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func calculateMedian(arr []int) int {
	n := len(arr)
	if n%2 == 1 {

		return arr[n/2]
	}

	return (arr[(n/2)-1] + arr[n/2]) / 2
}

func main() {
	var data []int
	var input int

	fmt.Println("Masukkan data (akhiri dengan -5313):")
	for {
		fmt.Scan(&input)

		if input == -5313 {
			break
		}

		if input == 0 {

			insertionSort(data)
			median := calculateMedian(data)
			fmt.Println(median)
		} else {

			data = append(data, input)
		}
	}
}
