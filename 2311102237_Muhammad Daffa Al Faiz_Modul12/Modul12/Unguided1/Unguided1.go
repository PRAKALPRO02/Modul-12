package main

import "fmt"

func ganjilgenap(arr []int) ([]int, []int) {
	var ganjil, genap []int
	for _, num := range arr {
		if num%2 == 0 {
			genap = append(genap, num)
		} else {
			ganjil = append(ganjil, num)
		}
	}
	return ganjil, genap
}

func selectionSortAsc(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}

func selectionSortDesc(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		max := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] > arr[max] {
				max = j
			}
		}
		arr[i], arr[max] = arr[max], arr[i]
	}
}

func main() {
	var n int
	fmt.Print("Masukan jumlah daerah (n): ")
	fmt.Scan(&n)
	if n <= 0 || n >= 1000 {
		fmt.Println("n harus lebih besar dari 0 dan kurang dari 1000")
		return
	}

	for i := 0; i < n; i++ {
		var m int
		fmt.Printf("Masukan jumlah rumah kerabat untuk daerah ke- %d: ", i+1)
		fmt.Scan(&m)
		if m <= 0 || m >= 1000000 {
			fmt.Println("m harus lebih besar dari 0 dan kurang dari 1000000")
			return
		}

		houses := make([]int, m)
		fmt.Printf("Masukan nomor rumah kerabat untuk daerah ke- %d: ", i+1)
		for j := range houses {
			fmt.Scan(&houses[j])
		}

		ganjil, genap := ganjilgenap(houses)
		selectionSortAsc(ganjil)
		selectionSortDesc(genap)
		result := append(ganjil, genap...)

		fmt.Printf("Hasil urutan rumah untuk daerah ke- %d: ", i+1)
		for _, house := range result {
			fmt.Printf("%d ", house)
		}
		fmt.Println()
	}
}
