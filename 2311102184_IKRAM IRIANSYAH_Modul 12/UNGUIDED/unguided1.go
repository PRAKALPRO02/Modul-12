//IKRAM IRIANSYAH
//2311102184

package main

import "fmt"

func slice_data(arr []int) ([]int, []int) {
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

func asc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func desc(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func main() {
	var n int
	fmt.Print("Masukkan Jumlah Daerah (n) : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("\nMasukkan Jumlah Rumah Daerah %d : ", i+1)
		var m int
		fmt.Scan(&m)

		arr := make([]int, m)
		fmt.Printf("Masukkan %d Nomor Rumah Daerah %d : ", m, i+1)
		for j := 0; j < m; j++ {
			fmt.Scan(&arr[j])
		}

		ganjil, genap := slice_data(arr)

		desc(genap)
		asc(ganjil)

		fmt.Printf("Nomor Rumah Terurut [ Ganjil(+), Genap(-) ] Daerah %d : ", i+1)
		for _, num := range ganjil {
			fmt.Printf("%d ", num)
		}
		for _, num := range genap {
			fmt.Printf("%d ", num)
		}
		fmt.Println()
	}
}