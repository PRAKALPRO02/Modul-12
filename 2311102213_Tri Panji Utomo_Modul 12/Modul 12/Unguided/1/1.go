package main

import "fmt"

func ganjilgenap(arr []int) ([]int, []int) {
	n := len(arr)
	var gnjil213 []int
	var gnp213 []int
	for i := 0; i < n; i++ {
		if arr[i]%2 == 0 {
			gnp213 = append(gnp213, arr[i])
		} else {
			gnjil213 = append(gnjil213, arr[i])
		}
	}
	return gnjil213, gnp213
}

func SSGanjil(arr []int) {
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
func SSGenap(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] > arr[maxIdx] { // Cari elemen terbesar
				maxIdx = j
			}
		}
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i] // Tukar elemen
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
		var m213 int
		fmt.Printf("Masukkan jumlah rumah kerabat untuk daerah ke-%d: ", i+1)
		fmt.Scan(&m213)

		if m213 <= 0 || m213 >= 1000000 {
			fmt.Println("m harus lebih besar dari 0 dan kurang dari 1000000.")
			return
		}

		houses := make([]int, m213)
		fmt.Printf("Masukkan nomor rumah kerabat untuk daerah ke-%d: ", i+1)
		for j := 0; j < m213; j++ {
			fmt.Scan(&houses[j])
		}
		gnjil213, gnp213 := ganjilgenap(houses)

		SSGanjil(gnjil213)
		SSGenap(gnp213)

		result := append(gnjil213, gnp213...)

		fmt.Printf("Hasil urutan rumah untuk daerah ke-%d: ", i+1)
		for _, house := range result {
			fmt.Printf("%d ", house)
		}
		fmt.Println()
	}
}