package main

import "fmt"

func selectionSortgenap(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[maxIdx] { // Cari elemen terbesar
				maxIdx = j
			}
		}
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i] // Tukar elemen
	}
}
func main() {
	var arr []int
	var arr_Baru []int
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
			selectionSortgenap(arr_Baru)
			if len(arr_Baru)%2 == 0 {

				median = (arr_Baru[(len(arr_Baru)/2)-1] + arr_Baru[len(arr_Baru)/2]) / 2

			} else {

				fmt.Println("gan", len(arr_Baru))
				median = arr_Baru[len(arr_Baru)/2]
				fmt.Print()
			}
			fmt.Println(median)
		} else {
			arr_Baru = append(arr_Baru, arr[i])
		}
	}

}
