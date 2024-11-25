package main

import "fmt"

func selectionSort(arr213 []int) {
	n := len(arr213)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if arr213[j] < arr213[maxIdx] { // Cari elemen terbesar
				maxIdx = j
			}
		}
		arr213[i], arr213[maxIdx] = arr213[maxIdx], arr213[i] // Tukar elemen
	}
}
func main() {
	var arr213 []int
	var slice213 []int
	var n int

	var median int
	for {
		fmt.Scan(&n)
		arr213 = append(arr213, n)
		if n == -5313 {
			break
		}
	}
	panjang := len(arr213)

	for i := 0; i < panjang; i++ {
		if arr213[i] == 0 {
			selectionSort(slice213)
			if len(slice213)%2 == 0 {

				median = (slice213[(len(slice213)/2)-1] + slice213[len(slice213)/2]) / 2

			} else {

				
				median = slice213[len(slice213)/2]
				fmt.Print()
			}
			fmt.Println(median)
		} else {
			slice213 = append(slice213, arr213[i])
		}
	}

}