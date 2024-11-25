package main

import "fmt"

// Fungsi untuk mengurutkan array menggunakan selection sort
func ganjilgenap(arr []int) ([]int, []int) {
	n := len(arr)
	var ganjil []int
	var genap []int
	for i := 0; i < n; i++ {
		if arr[i]%2 == 0 {
			genap = append(genap, arr[i])
		} else {
			ganjil = append(ganjil, arr[i])
		}
	}
	return ganjil, genap
}

func selectionSortganjil(arr []int) {
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
func selectionSortgenap(arr []int) {
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
		var m int
		fmt.Printf("Masukkan jumlah rumah kerabat untuk daerah ke-%d: ", i+1)
		fmt.Scan(&m)

		if m <= 0 || m >= 1000000 {
			fmt.Println("m harus lebih besar dari 0 dan kurang dari 1000000.")
			return
		}

		// Masukkan nomor rumah
		houses := make([]int, m)
		fmt.Printf("Masukkan nomor rumah kerabat untuk daerah ke-%d: ", i+1)
		for j := 0; j < m; j++ {
			fmt.Scan(&houses[j])
		}
		ganjil, genap := ganjilgenap(houses)

		// Urutkan dengan selection sort
		selectionSortganjil(ganjil)
		selectionSortgenap(genap)

		result := append(ganjil, genap...)

		// Cetak hasil
		fmt.Printf("Hasil urutan rumah untuk daerah ke-%d: ", i+1)
		for _, house := range result {
			fmt.Printf("%d ", house)
		}
		fmt.Println()
	}
}
