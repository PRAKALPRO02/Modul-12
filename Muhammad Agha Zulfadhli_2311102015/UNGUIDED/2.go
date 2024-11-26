package main

import (
	"fmt"
	"math"
	"strconv"
)

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[maxIdx] {
				maxIdx = j
			}
		}
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
	}
}

func main() {
	var x int
	var arr []int
	var text string
	for x >= 0 {
		fmt.Scan(&x)
		if x == 0 {
			selectionSort(arr)
			median := len(arr) / 2
			if len(arr)%median == 0 && median != 1 {
				text += "\n" + strconv.FormatFloat(math.Floor(float64(arr[median-1]+arr[median])/2), 'f', 0, 64)
			} else {
				text += "\n" + strconv.Itoa(arr[median])
			}
			continue
		} else if x < 0 {
			break
		}
		arr = append(arr, x)
	}
	fmt.Println(text)
}
