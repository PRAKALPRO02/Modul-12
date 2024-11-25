package main

import (
	"fmt"
	"math"
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

func isDataConsistentlySpaced(arr []int) (bool, int) {
	if len(arr) < 2 {
		return true, 0
	}

	diff := int(math.Abs(float64(arr[1] - arr[0])))

	for i := 1; i < len(arr)-1; i++ {
		currentDiff := int(math.Abs(float64(arr[i+1] - arr[i])))
		if currentDiff != diff {
			return false, 0
		}
	}

	return true, diff
}

func main() {
	var data []int
	var input int

	fmt.Println("Masukkan data (input negatif untuk selesai):")
	for {
		fmt.Scan(&input)
		if input < 0 {
			break
		}
		data = append(data, input)
	}

	insertionSort(data)

	isConsistent, diff := isDataConsistentlySpaced(data)

	fmt.Println("Hasil pengurutan:", data)
	if isConsistent {
		fmt.Printf("Data berjarak %d\n", diff)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}
