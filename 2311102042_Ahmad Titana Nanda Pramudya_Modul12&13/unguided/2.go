package main

import (
	"fmt"
	"sort"
)

func hitungMedian(data []int) int {
	n := len(data)
	if n == 0 {
		return 0
	}

	if n%2 != 0 {
		return data[n/2]
	}

	return (data[n/2-1] + data[n/2]) / 2
}

func main() {
	var masukan []int
	var data []int

	fmt.Printf("Masukkan angka :")

	for {
		
		var num int
		fmt.Scan(&num)

		
		if num < 0 {
			break
		}

		
		if num == 0 {
			
			tempData := make([]int, len(data))
			copy(tempData, data)
			sort.Ints(tempData)

			
			median := hitungMedian(tempData)
			masukan = append(masukan, median)
		} else {
			
			data = append(data, num)
		}
	}

	
	fmt.Println("Hasil keluaran:")
	for _, median := range masukan {
		fmt.Println(median)
	}
}
