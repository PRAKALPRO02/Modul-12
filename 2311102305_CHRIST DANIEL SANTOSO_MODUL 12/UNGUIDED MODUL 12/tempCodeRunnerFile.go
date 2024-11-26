package main

import "fmt"

func sortArray(arr []int) {
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

func calculateMedian(arr []int) int {
	sortArray(arr)
	n := len(arr)
	if n%2 == 1 {
		return arr[n/2]
	}
	mid := n / 2
	return (arr[mid-1] + arr[mid]) / 2
}

func main() {
	var data []int
	var input int

	fmt.Println("Masukkan data (ketik 0 untuk mencetak median, -5313 untuk keluar):")
	for {
		fmt.Scan(&input)

		if input == -5313 {
			fmt.Println("Program selesai.")
			break
		}

		if input == 0 {
			if len(data) == 0 {
				fmt.Println("Data kosong, tidak dapat menghitung median.")
			} else {
				median := calculateMedian(data)
				fmt.Printf("Median: %d\n", median)
			}
		} else {
			data = append(data, input)
		}
	}
}
