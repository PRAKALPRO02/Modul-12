package main

import "fmt"

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func printMedian(arr []int) {
	selectionSort(arr)
	n := len(arr)
	if n%2 == 1 {
		fmt.Printf("Median: %d\n", arr[n/2])
	} else {
		mid := n / 2
		median := (arr[mid-1] + arr[mid]) / 2
		fmt.Printf("Median: %d\n", median)
	}
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
				printMedian(data)
			}
		} else {
			data = append(data, input)
		}
	}
}
