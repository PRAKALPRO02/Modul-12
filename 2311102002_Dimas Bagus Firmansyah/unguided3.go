package main

import (
	"fmt"
	"sort"
)

const nMax = 7919

type Buku struct {
	id       int
	judul    string
	penulis  string
	penerbit string
	eksemplar int
	tahun    int
	rating   int
}

type DaftarBuku []Buku

// Subprogram untuk mencetak buku favorit
func CetakTerFavorit(pustaka DaftarBuku) {
	if len(pustaka) == 0 {
		fmt.Println("Tidak ada data buku.")
		return
	}

	// Cari buku dengan rating tertinggi
	terfavorit := pustaka[0]
	for _, buku := range pustaka {
		if buku.rating > terfavorit.rating {
			terfavorit = buku
		}
	}

	// Cetak buku favorit
	fmt.Printf("Buku Terfavorit: %s, %s, %s, %d\n", terfavorit.judul, terfavorit.penulis, terfavorit.penerbit, terfavorit.tahun)
}

// Subprogram untuk mengurutkan buku berdasarkan rating secara menurun
func UrutBuku(pustaka *DaftarBuku) {
	sort.Slice(*pustaka, func(i, j int) bool {
		return (*pustaka)[i].rating > (*pustaka)[j].rating
	})
}

// Subprogram untuk mencari buku berdasarkan rating
func CariBuku(pustaka DaftarBuku, r int) {
	for _, buku := range pustaka {
		if buku.rating == r {
			fmt.Printf("Buku ditemukan: %s, %s, %s, %d\n", buku.judul, buku.penulis, buku.penerbit, buku.tahun)
			return
		}
	}
	fmt.Println("Buku dengan rating yang diminta tidak ditemukan.")
}

func main() {
	var n, ratingCari int

	// Input jumlah buku
	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(&n)

	if n <= 0 || n > nMax {
		fmt.Printf("Jumlah buku harus antara 1 hingga %d.\n", nMax)
		return
	}

	// Input data buku
	pustaka := make(DaftarBuku, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data buku ke-%d (id, judul, penulis, penerbit, eksemplar, tahun, rating):\n", i+1)
		fmt.Scan(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis, &pustaka[i].penerbit, &pustaka[i].eksemplar, &pustaka[i].tahun, &pustaka[i].rating)
	}

	// Input rating yang akan dicari
	fmt.Print("Masukkan rating buku yang akan dicari: ")
	fmt.Scan(&ratingCari)

	// Cetak buku terfavorit
	CetakTerFavorit(pustaka)

	// Urutkan buku berdasarkan rating
	UrutBuku(&pustaka)

	// Cetak daftar buku setelah diurutkan
	fmt.Println("Daftar buku setelah diurutkan berdasarkan rating:")
	for _, buku := range pustaka {
		fmt.Printf("%s, %s, %s, %d, Rating: %d\n", buku.judul, buku.penulis, buku.penerbit, buku.tahun, buku.rating)
	}

	// Cari buku berdasarkan rating
	CariBuku(pustaka, ratingCari)
}
