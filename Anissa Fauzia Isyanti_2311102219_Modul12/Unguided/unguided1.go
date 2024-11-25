package main

import "fmt"

func selectionSortAsc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		idxMin := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[idxMin] {
				idxMin = j
			}
		}
		arr[i], arr[idxMin] = arr[idxMin], arr[i]
	}
}

func selectionSortDesc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		idxMax := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[idxMax] {
				idxMax = j
			}
		}
		arr[i], arr[idxMax] = arr[idxMax], arr[i]
	}
}

func main() {
	var jmlDaerah int
	fmt.Println("Input")
	fmt.Scan(&jmlDaerah)

	if jmlDaerah <= 0 || jmlDaerah >= 1000 {
		fmt.Println("Jumlah daerah harus di antara 1 dan 999.")
		return
	}
	data := make([][]int, jmlDaerah)

	for i := 0; i < jmlDaerah; i++ {
		var jmlRumah int
		fmt.Scan(&jmlRumah)

		if jmlRumah <= 0 || jmlRumah >= 1000000 {
			fmt.Println("Jumlah rumah harus di antara 1 dan 999999.")
			return
		}

		data[i] = make([]int, jmlRumah)
		for j := 0; j < jmlRumah; j++ {
			fmt.Scan(&data[i][j])
		}
	}

	fmt.Println("\nOutput:")
	for _, daerah := range data {
		var ganjil []int
		var genap []int

		for _, angka := range daerah {
			if angka%2 == 0 {
				genap = append(genap, angka)
			} else {
				ganjil = append(ganjil, angka)
			}
		}

		selectionSortAsc(ganjil)
		selectionSortDesc(genap)
		hasil := append(ganjil, genap...)

		for _, nilai := range hasil {
			fmt.Print(nilai, " ")
		}
		fmt.Println()
	}
}
