package main

import "fmt"

func selectionSortAsc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		minIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIdx] {
				minIdx = j
			}
		}
		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}
}

func main() {
	var n int
	fmt.Println("Masukan")
	fmt.Scan(&n)

	if n <= 0 || n >= 1000 {
		fmt.Println("Jumlah daerah harus di antara 1 dan 999.")
		return
	}

	masukan := make([][]int, n)
	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)

		if m <= 0 || m >= 1000000 {
			fmt.Println("Jumlah rumah harus di antara 1 dan 999999.")
			return
		}

		masukan[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&masukan[i][j])
		}
	}

	fmt.Println("\nKeluaran:")
	for _, daerah := range masukan {
		var ganjil []int
		var genap []int

		for _, num := range daerah {
			if num%2 == 0 {
				genap = append(genap, num)
			} else {
				ganjil = append(ganjil, num)
			}
		}

		selectionSortAsc(ganjil)
		selectionSortAsc(genap)
		hasil := append(ganjil, genap...)

		for _, val := range hasil {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}
}
