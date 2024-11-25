package main

import "fmt"

func selectionSort_Ganjil(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[maxIdx] {
				maxIdx = j
			}
		}
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
	}
}

func selectionSort_Genap(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[maxIdx] {
				maxIdx = j
			}
		}
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
	}
}

func main() {
	var n, input int
	fmt.Print("Masukkan jumlah daerah (n): ")
	fmt.Scan(&n)

	if n <= 0 || n >= 1000 {
		fmt.Println("n harus lebih besar dari 0 dan kurang dari 1000.")
		return
	}

	for i := 0; i < n; i++ {
		var m int
		fmt.Printf("Masukkan jumlah rumah kerabat untuk daerah ke-%d: ", i+1)
		fmt.Scan(&m)

		if m <= 0 || m >= 1000000 {
			fmt.Println("m harus lebih besar dari 0 dan kurang dari 1000000.")
			return
		}

		nomerGanjil := make([]int, 0)
		nomerGenap := make([]int, 0)
		fmt.Printf("Masukkan nomor rumah kerabat untuk daerah ke-%d: ", i+1)
		for j := 0; j < m; j++ {
			fmt.Scan(&input)
			if input%2 == 0 {
				nomerGenap = append(nomerGenap, input)
			} else {
				nomerGanjil = append(nomerGanjil, input)
			}
		}

		// Urutkan dengan selection sort
		selectionSort_Ganjil(nomerGanjil)
		selectionSort_Genap(nomerGenap)

		fmt.Printf("Hasil urutan rumah untuk daerah ke-%d: ", i+1)
		for _, rumah := range nomerGanjil {
			fmt.Printf("%d ", rumah)
		}
		for _, rumah := range nomerGenap {
			fmt.Printf("%d ", rumah)
		}
		fmt.Println()
	}
}
