package main

import "fmt"

func printArray(arr []int) {
	n := len(arr)
	if n == 0 {
		return
	}
	for i := 0; i < n; i++ {
		fmt.Printf("%v ", arr[i])
	}
}

func sortArray(arr *[]int, n int) {
	if n == 0 {
		return
	}

	if (*arr)[0]%2 == 0 {
		for i := 0; i < n-1; i++ {
			maxIdx := i
			for j := i + 1; j < n; j++ {
				if (*arr)[j] > (*arr)[maxIdx] { // Cari elemen terbesar
					maxIdx = j
				}
			}
			(*arr)[i], (*arr)[maxIdx] = (*arr)[maxIdx], (*arr)[i] // Tukar elemen
		}
	} else {
		for i := 0; i < n-1; i++ {
			minIdx := i
			for j := i + 1; j < n; j++ {
				if (*arr)[j] < (*arr)[minIdx] { // Cari elemen terkecil
					minIdx = j
				}
			}
			(*arr)[i], (*arr)[minIdx] = (*arr)[minIdx], (*arr)[i] // Tukar elemen
		}
	}
}

func splitEvenOdd(arr []int, n int) ([]int, []int) {
	var even, odd []int
	for i := 0; i < n; i++ {
		if arr[i]%2 == 0 {
			even = append(even, arr[i])
		} else {
			odd = append(odd, arr[i])
		}
	}
	return even, odd
}

func main() {
	var nIteration int

	fmt.Print("Masukkan banyak rangkaian rumah kerabat: ")
	fmt.Scan(&nIteration)

	for i := 0; i < nIteration; i++ {
		var nHouses int
		fmt.Printf("Masukkan banyak rumah untuk rangkian ke - %v: ", i+1)
		fmt.Scan(&nHouses)

		arrHouses := make([]int, nHouses)
		fmt.Printf("Masukkan nomor - nomor rumah untuk rangkian ke - %v:\n", i+1)
		for j := 0; j < nHouses; j++ {
			fmt.Scan(&arrHouses[j])
		}

		evenHouses, oddHouses := splitEvenOdd(arrHouses, nHouses)

		sortArray(&evenHouses, len(evenHouses))
		sortArray(&oddHouses, len(oddHouses))

		printArray(oddHouses)
		printArray(evenHouses)
		fmt.Println()
	}
}
