package main

import "fmt"

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[idx] {
				idx = j
			}
		}
		arr[i], arr[idx] = arr[idx], arr[i]
	}
}

func printMedian(arr []int) {
	selectionSort(arr)
	n := len(arr)
	if n%2 == 1 {
		fmt.Println(arr[n/2])
	} else {
		mid := n / 2
		median := (arr[mid-1] + arr[mid]) / 2
		fmt.Println(median)
	}
}

func main() {
	var data []int
	var input int
	for {
		fmt.Scan(&input)
		if input == -5313 {
			break
		}
		if input == 0 {
			printMedian(data)
		} else {
			data = append(data, input)
		}
	}
}