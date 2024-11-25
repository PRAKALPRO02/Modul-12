package main

import (
	"fmt"
	"sort"
)

const nMax = 7919

type Buku struct {
	ID       string
	Judul    string
	Penulis  string
	Penerbit string
	Eksemplar int
	Tahun    int
	Rating   int
}

type DaftarBuku []Buku

func tampilkanBuku(buku Buku) {
	fmt.Printf("ID: %s, Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n",
		buku.ID, buku.Judul, buku.Penulis, buku.Penerbit, buku.Tahun, buku.Rating)
}

func cetakTerfavorit(pustaka DaftarBuku) {
	if len(pustaka) == 0 {
		fmt.Println("Tidak ada buku di pustaka.")
		return
	}
	terfavorit := pustaka[0]
	for _, buku := range pustaka {
		if buku.Rating > terfavorit.Rating {
			terfavorit = buku
		}
	}
	fmt.Println("Buku Terfavorit:")
	tampilkanBuku(terfavorit)
}

func urutkanBuku(pustaka DaftarBuku) {
	sort.Slice(pustaka, func(i, j int) bool {
		return pustaka[i].Rating > pustaka[j].Rating
	})
}

func cetak5Terbaru(pustaka DaftarBuku) {
	if len(pustaka) == 0 {
		fmt.Println("Tidak ada buku di pustaka.")
		return
	}
	fmt.Println("5 Buku dengan Rating Tertinggi:")
	for i := 0; i < 5 && i < len(pustaka); i++ {
		tampilkanBuku(pustaka[i])
	}
}

func cariBuku(pustaka DaftarBuku, rating int) {
	urutkanBuku(pustaka)
	left, right := 0, len(pustaka)-1
	for left <= right {
		mid := (left + right) / 2
		if pustaka[mid].Rating == rating {
			fmt.Println("Buku ditemukan dengan rating tersebut:")
			tampilkanBuku(pustaka[mid])
			return
		} else if pustaka[mid].Rating < rating {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	fmt.Println("Tidak ada buku dengan rating seperti itu.")
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(&n)

	pustaka := make(DaftarBuku, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data untuk buku ke-%d (ID, Judul, Penulis, Penerbit, Eksemplar, Tahun, Rating):\n", i+1)
		fmt.Scan(&pustaka[i].ID, &pustaka[i].Judul, &pustaka[i].Penulis, &pustaka[i].Penerbit,
			&pustaka[i].Eksemplar, &pustaka[i].Tahun, &pustaka[i].Rating)
	}

	var ratingCari int
	fmt.Print("Masukkan rating buku yang ingin dicari: ")
	fmt.Scan(&ratingCari)

	cetakTerfavorit(pustaka)

	urutkanBuku(pustaka)
	cetak5Terbaru(pustaka)

	cariBuku(pustaka, ratingCari)
}
