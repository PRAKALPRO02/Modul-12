package main

import (
	"fmt"
	"sort"
)

// Definisi struct Buku
type Buku struct {
	ID        string
	Judul     string
	Penulis   string
	Penerbit  string
	Eksemplar int
	Tahun     int
	Rating    int
}

// DaftarBuku adalah slice dari Buku
type DaftarBuku []Buku

// Prosedur DaftarkanBuku
func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		var buku Buku
		fmt.Println("Masukkan data buku (ID, Judul, Penulis, Penerbit, Eksemplar, Tahun, Rating):")
		fmt.Scan(&buku.ID, &buku.Judul, &buku.Penulis, &buku.Penerbit, &buku.Eksemplar, &buku.Tahun, &buku.Rating)
		*pustaka = append(*pustaka, buku)
	}
}

// Prosedur CetakTerfavorit
func CetakTerfavorit(pustaka DaftarBuku) {
	if len(pustaka) == 0 {
		fmt.Println("Tidak ada buku.")
		return
	}
	maxRating := pustaka[0]
	for _, buku := range pustaka {
		if buku.Rating > maxRating.Rating {
			maxRating = buku
		}
	}
	fmt.Printf("Buku terfavorit: %s, %s, %s, %d\n", maxRating.Judul, maxRating.Penulis, maxRating.Penerbit, maxRating.Tahun)
}

// Prosedur UrutBuku
func UrutBuku(pustaka *DaftarBuku) {
	sort.Slice(*pustaka, func(i, j int) bool {
		return (*pustaka)[i].Rating > (*pustaka)[j].Rating
	})
}

// Prosedur Cetak5Terbaru
func Cetak5Terbaru(pustaka DaftarBuku) {
	fmt.Println("5 Buku dengan rating tertinggi:")
	for i := 0; i < len(pustaka) && i < 5; i++ {
		fmt.Println(pustaka[i].Judul)
	}
}

// Prosedur CariBuku
func CariBuku(pustaka DaftarBuku, r int) {
	low, high := 0, len(pustaka)-1
	for low <= high {
		mid := (low + high) / 2
		if pustaka[mid].Rating == r {
			buku := pustaka[mid]
			fmt.Printf("Buku ditemukan: %s, %s, %s, %d, %d, %d\n", buku.Judul, buku.Penulis, buku.Penerbit, buku.Tahun, buku.Eksemplar, buku.Rating)
			return
		} else if pustaka[mid].Rating > r {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	fmt.Println("Tidak ada buku dengan rating seperti itu.")
}

func main() {
	var pustaka DaftarBuku
	var n int
	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(&n)

	// Daftarkan buku
	DaftarkanBuku(&pustaka, n)

	// Urutkan buku berdasarkan rating
	UrutBuku(&pustaka)

	// Cetak buku terfavorit
	CetakTerfavorit(pustaka)

	// Cetak 5 buku dengan rating tertinggi
	Cetak5Terbaru(pustaka)

	// Cari buku berdasarkan rating
	var rating int
	fmt.Print("Masukkan rating yang akan dicari: ")
	fmt.Scan(&rating)
	CariBuku(pustaka, rating)
}
