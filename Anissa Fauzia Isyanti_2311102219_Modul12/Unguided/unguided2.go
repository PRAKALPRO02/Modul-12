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

func calculateMedian(arr []int) int {
	n := len(arr)
	if n%2 == 1 {
		return arr[n/2]
	}
	return (arr[n/2-1] + arr[n/2]) / 2
}

func main() {
	var input int
	var data []int

	for {
		fmt.Scan(&input)
		if input == -5313 {
			break
		}
		if input == 0 {
			if len(data) == 0 {
				continue
			}
			selectionSort(data)
			median := calculateMedian(data)
			fmt.Println(median)
		} else {
			data = append(data, input)
		}
	}
}
