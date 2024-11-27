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

func selectionSortDesc(arr []int) {
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

func main() {
	var n int
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

		rumah := make([]int, m)
		fmt.Printf("Masukkan nomor rumah kerabat untuk daerah ke-%d: ", i+1)
		for j := 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}

		var oddNumbers, evenNumbers []int
		for _, house := range rumah {
			if house%2 == 1 {
				oddNumbers = append(oddNumbers, house)
			} else { 
				evenNumbers = append(evenNumbers, house)
			}
		}

		selectionSortAsc(oddNumbers)
		selectionSortDesc(evenNumbers)

		fmt.Printf("Hasil urutan rumah untuk daerah ke-%d: ", i+1)

		for _, house := range oddNumbers {
			fmt.Printf("%d ", house)
		}

		for _, house := range evenNumbers {
			fmt.Printf("%d ", house)
		}
		fmt.Println()
	}
}