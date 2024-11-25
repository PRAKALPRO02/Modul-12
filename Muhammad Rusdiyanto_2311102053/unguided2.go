package main

import "fmt"

func sortArr(arr *[]int, n int) {
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if (*arr)[j] < (*arr)[minIdx] { // Cari elemen terkecil
				minIdx = j
			}
		}
		(*arr)[i], (*arr)[minIdx] = (*arr)[minIdx], (*arr)[i] // Tukar elemen
	}
}

func main() {
	const ARR_MAX int = 1000000
	var arrNum []int
	var elemCount int

	for i := 0; i < ARR_MAX; i++ {
		var temp int
		fmt.Scan(&temp)

		if temp < 0 {
			break
		} else {
			arrNum = append(arrNum, temp)
			elemCount++
		}
	}

	var arrFiltered []int
	for i := 0; i < elemCount; i++ {
		if arrNum[i] == 0 {
			n := len(arrFiltered)

			if n == 0 {
				continue // do nothing
			} else {
				sortArr(&arrFiltered, n)
				var index int = (n - 1) / 2
				if n%2 == 0 {
					fmt.Print((arrFiltered[index] + arrFiltered[index+1]) / 2)
				} else {
					fmt.Print(arrFiltered[index])
				}
			}
			fmt.Println()
		} else {
			arrFiltered = append(arrFiltered, arrNum[i])
		}
	}
}
