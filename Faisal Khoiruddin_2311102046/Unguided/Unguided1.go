package main

import "fmt"

func selectionSort(arr []int, menaik bool) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		targetIdx := i
		for j := i + 1; j < n; j++ {
			if menaik {
				if arr[j] < arr[targetIdx] {
					targetIdx = j
				}
			} else {
				if arr[j] > arr[targetIdx] {
					targetIdx = j
				}
			}
		}
		arr[i], arr[targetIdx] = arr[targetIdx], arr[i]
	}
}

func main() {
	var jumlahDaerah int
	fmt.Print("Masukkan jumlah daerah (n): ")
	fmt.Scan(&jumlahDaerah)

	if jumlahDaerah <= 0 || jumlahDaerah >= 1000 {
		fmt.Println("Jumlah daerah harus lebih besar dari 0 dan kurang dari 1000.")
		return
	}

	for i := 0; i < jumlahDaerah; i++ {
		var jumlahRumah int
		fmt.Printf("Masukkan jumlah rumah kerabat untuk daerah ke-%d: ", i+1)
		fmt.Scan(&jumlahRumah)

		if jumlahRumah <= 0 || jumlahRumah >= 1000000 {
			fmt.Println("jumlah rumah harus lebih besar dari 0 dan kurang dari 1000000.")
			return
		}

		rumah := make([]int, jumlahRumah)
		fmt.Printf("Masukkan nomor rumah kerabat untuk daerah ke-%d: ", i+1)
		for j := 0; j < jumlahRumah; j++ {
			fmt.Scan(&rumah[j])
		}

		var rumahGanjil, rumahGenap []int
		for _, nomorRumah := range rumah {
			if nomorRumah%2 == 1 {
				rumahGanjil = append(rumahGanjil, nomorRumah)
			} else {
				rumahGenap = append(rumahGenap, nomorRumah)
			}
		}

		selectionSort(rumahGanjil, true)
		selectionSort(rumahGenap, false)

		fmt.Printf("Hasil urutan rumah untuk daerah ke-%d: ", i+1)
		for _, nomorRumah := range rumahGanjil {
			fmt.Printf("%d ", nomorRumah)
		}
		for _, nomorRumah := range rumahGenap {
			fmt.Printf("%d ", nomorRumah)
		}
		fmt.Println()
	}
}
