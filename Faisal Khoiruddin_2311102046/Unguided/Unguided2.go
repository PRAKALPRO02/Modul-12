package main

import (
	"fmt"
	"sort"
)

func hitungMedian(arr []int) int {
	sort.Ints(arr)
	n := len(arr)

	if n%2 == 0 {
		return (arr[n/2-1] + arr[n/2]) / 2
	}
	return arr[n/2]
}

func main() {
	var num int
	var data []int

	for {
		fmt.Scan(&num)
		if num == -5313 {
			break
		}
		if num == 0 {
			if len(data) > 0 {
				median := hitungMedian(data)
				fmt.Println(median)
			}
		} else {
			data = append(data, num)
		}
	}
}
