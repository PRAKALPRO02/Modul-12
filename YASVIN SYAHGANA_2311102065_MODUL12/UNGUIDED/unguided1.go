package main

import "fmt"

func GanjilSort(nomor []int) {
	for i := 0; i < len(nomor)-1; i++ {
		Max := i
		for j := i + 1; j < len(nomor); j++ {
			if nomor[j] < nomor[Max] {
				Max = j
			}
		}
		nomor[i], nomor[Max] = nomor[Max], nomor[i]
	}
}

func GenapSort(nomor []int) {
	for i := 0; i < len(nomor)-1; i++ {
		min := i
		for j := i + 1; j < len(nomor); j++ {
			if nomor[j] > nomor[min] {
				min = j
			}
		}
		nomor[i], nomor[min] = nomor[min], nomor[i]
	}
}

func main() {
	var n, m int

	fmt.Print("Inputkan banyaknya daerah: ")
	fmt.Scan(&n)
	if n <= 0 || n >= 1000 {
		fmt.Println("n harus lebih besar dari 0 dan kurang dari 1000.")
		return
	}

	for i := 0; i < n; i++ {
		var nmrGanjil, nmrGenap []int
		fmt.Print("Jumlah rumah kerabat: ")
		fmt.Scan(&m)

		if m <= 0 || m >= 1000000 {
			fmt.Println("m harus lebih besar dari 0 dan kurang dari 1000000.")
			return
		}

		nomor := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&nomor[j])
		}

		for _, x := range nomor {
			if x%2 == 0 {
				nmrGenap = append(nmrGenap, x)
			} else {
				nmrGanjil = append(nmrGanjil, x)
			}
		}

		GanjilSort(nmrGanjil)
		GenapSort(nmrGenap)

		fmt.Print("Urutan rumah: ")
		for _, rmhGanjil := range nmrGanjil {
			fmt.Print(rmhGanjil, " ")
		}
		for _, rmhGenap := range nmrGenap {
			fmt.Print(rmhGenap, " ")
		}
		fmt.Println()
	}
}
