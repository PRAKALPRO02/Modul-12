package main

import "fmt"

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
	if n%2 == 0 {
		return (arr[(n/2)-1] + arr[(n/2)]) / 2
	} else {
		return arr[(n / 2)]
	}
}

func main() {
	var inputData int
	var numberCollection = make([]int, 0)
	var currentNumbers = make([]int, 0)

	for inputData != -5313 {
		fmt.Scan(&inputData)
		if inputData != -5313 {
			numberCollection = append(numberCollection, inputData)
		}
	}

	for _, number := range numberCollection {
		if number == 0 {
			insertionSort(currentNumbers)
			fmt.Println(calculateMedian(currentNumbers))
		} else {
			currentNumbers = append(currentNumbers, number)
		}
	}
}
