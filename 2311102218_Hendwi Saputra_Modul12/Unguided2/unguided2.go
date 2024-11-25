package main

import "fmt"

func insertion_Sort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func mencariMedian(arr []int) int {
	n := len(arr)
	if n%2 == 0 {
		return (arr[(n/2)-1] + arr[(n/2)]) / 2

	} else {

		return arr[(n / 2)]
	}
}

func main() {
	var input int
	var kumpulan_Angka = make([]int, 0)
	var angka_Sliced = make([]int, 0)
	for input != -5313 {
		fmt.Scan(&input)
		if input != -5313 {
			kumpulan_Angka = append(kumpulan_Angka, input)
		}
	}
	for _, angka := range kumpulan_Angka {
		if angka == 0 {
			insertion_Sort(angka_Sliced)
			fmt.Println(mencariMedian(angka_Sliced))

		} else {
			angka_Sliced = append(angka_Sliced, angka)
		}
	}
}
