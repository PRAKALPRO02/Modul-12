package main

import "fmt"

func selectionSort(arr []int, ascending bool) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if (ascending && arr[j] < arr[idx]) || (!ascending && arr[j] > arr[idx]) {
				idx = j
			}
		}
		arr[i], arr[idx] = arr[idx], arr[i]
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)
		houses := make([]int, m)
		odd := []int{}
		even := []int{}

		for j := 0; j < m; j++ {
			fmt.Scan(&houses[j])
			if houses[j]%2 == 0 {
				even = append(even, houses[j])
			} else {
				odd = append(odd, houses[j])
			}
		}

		selectionSort(odd, true)
		selectionSort(even, false)

		for i := 0; i < len(odd); i++ {
			fmt.Print(odd[i], " ")
		}
		for i := 0; i < len(even); i++ {
			fmt.Print(even[i], " ")
		}
		fmt.Println()
	}
}