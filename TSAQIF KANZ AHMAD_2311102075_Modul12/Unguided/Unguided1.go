package main

import (
	"fmt"
)

func main() {
	var nDaerah int
	fmt.Print("Masukan Jumlah daerah: ")
	fmt.Scan(&nDaerah)

	if nDaerah <= 0 {
		fmt.Println("Error: Jumlah daerah harus lebih dari 0")
		return
	}

	for i := 1; i <= nDaerah; i++ {
		var nRumah int
		fmt.Printf("Masukan Jumlah rumah di daerah ke-%d: ", i)
		fmt.Scan(&nRumah)

		if nRumah <= 0 {
			fmt.Println("Jumlah rumah harus positif!")
			return
		}

		dataRumah := make([]int, nRumah)
		fmt.Printf("Masukkan nomor rumah di daerah ke-%d: ", i)
		for j := 0; j < nRumah; j++ {
			fmt.Scan(&dataRumah[j])
		}

		ganjil := []int{}
		genap := []int{}
		for _, x := range dataRumah {
			if x%2 == 0 {
				genap = append(genap, x)
			} else {
				ganjil = append(ganjil, x)
			}
		}
		sortSelection(ganjil, false)
		sortSelection(genap, true)

		fmt.Printf("Hasil daerah ke-%d: ", i)
		for _, g := range ganjil {
			fmt.Printf("%d ", g)
		}
		for _, g := range genap {
			fmt.Printf("%d ", g)
		}
		fmt.Println()
	}
}

func sortSelection(arr []int, ascending bool) {
	for i := 0; i < len(arr)-1; i++ {
		idx := i
		for j := i + 1; j < len(arr); j++ {
			if (ascending && arr[j] < arr[idx]) || (!ascending && arr[j] > arr[idx]) {
				idx = j
			}
		}
		arr[i], arr[idx] = arr[idx], arr[i]
	}
}
