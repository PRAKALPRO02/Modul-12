package main

import (
	"fmt"
	"sort"
)

const nMax = 7919

type Buku struct {
	ID, Judul, Penulis, Penerbit string
	Eksemplar, Tahun, Rating     int
}

type DaftarBuku struct {
	Pustaka  []Buku
	nPustaka int
}

// Procedure untuk mengisi daftar buku
func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	var buku Buku
	for i := 0; i < n; i++ {
		fmt.Println("Masukkan data buku (ID, Judul, Penulis, Penerbit, Eksemplar, Tahun, Rating):")
		fmt.Scan(&buku.ID, &buku.Judul, &buku.Penulis, &buku.Penerbit, &buku.Eksemplar, &buku.Tahun, &buku.Rating)
		pustaka.Pustaka = append(pustaka.Pustaka, buku)
	}
	pustaka.nPustaka = n
}

// Procedure untuk mencetak buku dengan rating tertinggi
func CetakTerfavorit(pustaka DaftarBuku) {
	if pustaka.nPustaka == 0 {
		fmt.Println("Tidak ada data buku.")
		return
	}
	maxRating := pustaka.Pustaka[0].Rating
	var favorit Buku
	for _, buku := range pustaka.Pustaka {
		if buku.Rating > maxRating {
			maxRating = buku.Rating
			favorit = buku
		}
	}
	fmt.Printf("Buku terfavorit: %s, Penulis: %s, Penerbit: %s, Tahun: %d\n", favorit.Judul, favorit.Penulis, favorit.Penerbit, favorit.Tahun)
}

// Procedure untuk mengurutkan buku berdasarkan rating (Insertion Sort)
func UrutBuku(pustaka *DaftarBuku) {
	sort.Slice(pustaka.Pustaka, func(i, j int) bool {
		return pustaka.Pustaka[i].Rating > pustaka.Pustaka[j].Rating
	})
	fmt.Println("Buku telah diurutkan berdasarkan rating.")
}

// Procedure untuk mencetak 5 buku terbaru
func Cetak5Terbaru(pustaka DaftarBuku) {
	count := 5
	if pustaka.nPustaka < count {
		count = pustaka.nPustaka
	}
	fmt.Println("5 Buku Terbaru:")
	for i := 0; i < count; i++ {
		buku := pustaka.Pustaka[i]
		fmt.Printf("%s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n", buku.Judul, buku.Penulis, buku.Penerbit, buku.Tahun, buku.Rating)
	}
}

// Procedure untuk mencari buku dengan rating tertentu
func CariBuku(pustaka DaftarBuku, r int) {
	low, high := 0, pustaka.nPustaka-1
	for low <= high {
		mid := (low + high) / 2
		if pustaka.Pustaka[mid].Rating == r {
			buku := pustaka.Pustaka[mid]
			fmt.Printf("Buku ditemukan: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Eksemplar: %d, Rating: %d\n",
				buku.Judul, buku.Penulis, buku.Penerbit, buku.Tahun, buku.Eksemplar, buku.Rating)
			return
		} else if pustaka.Pustaka[mid].Rating < r {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	fmt.Println("Tidak ada buku dengan rating seperti itu.")
}

func main() {
	var pustaka DaftarBuku
	var n, rating int

	fmt.Println("Masukkan jumlah buku:")
	fmt.Scan(&n)

	DaftarkanBuku(&pustaka, n)
	UrutBuku(&pustaka)
	CetakTerfavorit(pustaka)
	Cetak5Terbaru(pustaka)

	fmt.Println("Masukkan rating yang ingin dicari:")
	fmt.Scan(&rating)
	CariBuku(pustaka, rating)
}
