package main

import "fmt"

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[maxIdx] {
				maxIdx = j
			}
		}
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
	}
}
func main() {
	var arr []int
	var slice []int
	var n int

	var median int
	for {
		fmt.Scan(&n)
		arr = append(arr, n)
		if n == -5313 {
			break
		}
	}
	panjang := len(arr)

	for i := 0; i < panjang; i++ {
		if arr[i] == 0 {
			selectionSort(slice)
			if len(slice)%2 == 0 {

				median = (slice[(len(slice)/2)-1] + slice[len(slice)/2]) / 2

			} else {

				median = slice[len(slice)/2]
				fmt.Print()
			}
			fmt.Println(median)
		} else {
			slice = append(slice, arr[i])
		}
	}

}
