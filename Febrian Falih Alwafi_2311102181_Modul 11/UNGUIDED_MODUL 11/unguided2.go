package main

import (
	"fmt"
	"sort"
)


func calculateMedian_2311102181(arr []int) int {
	n := len(arr)
	if n%2 == 1 {
		return arr[n/2]
	}
	return (arr[n/2-1] + arr[n/2]) / 2
}

func main() {
	var data []int
	var input int

	fmt.Println("Masukkan angka (0 untuk menghitung median, -5313 untuk keluar):")
	for {
		fmt.Scan(&input)

		if input == -5313 {
			break 
		}

		if input == 0 {
			
			if len(data) == 0 {
				fmt.Println("Data kosong, tidak dapat menghitung median.")
				continue
			}

			
			sort.Ints(data)
			median := calculateMedian_2311102181(data)
			fmt.Println("Median:", median)
		} else {
			
			data = append(data, input)
		}
	}
}
