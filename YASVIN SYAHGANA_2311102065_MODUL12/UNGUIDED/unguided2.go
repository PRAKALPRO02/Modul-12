package main

import "fmt"

func InsertionSorting(median []int) {
	for i := 0; i < len(median); i++ {
		cur := median[i]
		j := i - 1
		for j >= 0 && median[j] > cur {
			median[j+1] = median[j]
			j--
		}
		median[j+1] = cur
	}
}

func HitungMedian(median []int) float64 {
	temp := make([]int, len(median))
	copy(temp, median)

	InsertionSorting(temp)
	n := len(temp)
	if n%2 == 1 {
		return float64(temp[n/2])
	} else {
		return float64(temp[n/2-1]+temp[n/2]) / 2
	}
}

func main() {
	var n int
	var median []int
	for {
		fmt.Scan(&n)

		if n == -5313 {
			break
		}

		if n == 0 {
			fmt.Println(HitungMedian(median))
		} else {
			median = append(median, n)
		}

	}
}
