package main

import "fmt"

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		// Geser elemen yang lebih besar dari key ke kanan
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func median(arr []int) int {
	n := len(arr)
	if n%2 == 0 {
		return (arr[(n/2)-1] + arr[(n/2)]) / 2

	} else {

		return arr[(n / 2)]
	}
}

func main() {
	var input int
	var kumpulanAngka = make([]int, 0)
	var angkaSliced = make([]int, 0)
	for input != -5313 {
		fmt.Scan(&input)
		if input != -5313 {
			kumpulanAngka = append(kumpulanAngka, input)
		}
	}
	for _, angka := range kumpulanAngka {
		if angka == 0 {
			insertionSort(angkaSliced)
			fmt.Println(median(angkaSliced))

		} else {
			angkaSliced = append(angkaSliced, angka)
		}
	}
}
