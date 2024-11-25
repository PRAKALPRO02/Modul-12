package main

import "fmt"

const nmax = 7919

type Buku struct {
	ID, Judul, Penulis, Penerbit string
	Eksemplar, Tahun, Ranting    int
}
type DaftarBuku [nmax]Buku

func insertion_Sort(Pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		key := Pustaka[i]
		j := i - 1

		for j >= 0 && Pustaka[j].Ranting < key.Ranting {
			Pustaka[j+1] = Pustaka[j]
			j--
		}
		Pustaka[j+1] = key
	}
}

func cari_Favorit(Pustaka DaftarBuku, n int) {
	max := Pustaka[0].Ranting
	favorit := 0
	for i := 0; i < n; i++ {
		if Pustaka[i].Ranting > max {
			max = Pustaka[i].Ranting
			favorit = i
		}
	}
	fmt.Printf("Buku Terfavorit Adalah : %s %s %s %s %v \n", Pustaka[favorit].ID, Pustaka[favorit].Judul, Pustaka[favorit].Penulis, Pustaka[favorit].Penerbit, Pustaka[favorit].Tahun)
}

func LimaRating(Pustaka DaftarBuku, n int) {
	insertion_Sort(&Pustaka, n)
	fmt.Print("Lima Rating Tertinggi : ")
	for i := 0; i < n; i++ {
		fmt.Print(Pustaka[i].Judul, " ")

	}
	fmt.Println()
}

func binary_Search(Pustaka DaftarBuku, n, target int) int {
	low, high := 0, n-1

	for low <= high {
		mid := (low + high) / 2

		if Pustaka[mid].Ranting == target {
			return mid
		} else if Pustaka[mid].Ranting < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return -1
}
func main() {
	var Buku DaftarBuku
	var n, Cari int
	fmt.Print("Masukan Banyak Buku : ")
	fmt.Scan(&n)
	fmt.Println("Masukan (ID, Judul, Penulis, Penerbit, Eksemplar, Tahun, Rating)")
	for i := 0; i < n; i++ {
		fmt.Print("Masukan : ")
		fmt.Scan(&Buku[i].ID, &Buku[i].Judul, &Buku[i].Penulis, &Buku[i].Penerbit, &Buku[i].Eksemplar, &Buku[i].Tahun, &Buku[i].Ranting)
	}

	fmt.Print("Masukan Rating Buku Yang Anda Cari : ")
	fmt.Scan(&Cari)

	cari_Favorit(Buku, n)
	LimaRating(Buku, n)
	Temukan := binary_Search(Buku, n, Cari)
	if Temukan != -1 {
		fmt.Printf("Buku dengan Rating %v : %s %s %s %s %v %v %v\n", Cari, Buku[Temukan].ID, Buku[Temukan].Judul, Buku[Temukan].Penulis, Buku[Temukan].Penerbit, Buku[Temukan].Eksemplar, Buku[Temukan].Tahun, Buku[Temukan].Ranting)
	} else {
		fmt.Printf("Buku dengan Reting %v tidak ditemukan", Cari)
	}

}
