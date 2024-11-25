package main

import (
	"fmt"
)

const nMax = 7919

type Buku struct {
	id        int
	judul     string
	penulis   string
	penerbit  string
	eksemplar int
	tahun     int
	rating    int
}

type DaftarBuku [nMax]Buku

// Prosedur 1: Daftarkan Buku
func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data buku ke-%d:\n", i+1)

		// Input masing-masing atribut buku
		fmt.Print("Masukkan id buku: ")
		fmt.Scan(&pustaka[i].id)

		fmt.Print("Masukkan judul buku: ")
		fmt.Scan(&pustaka[i].judul)

		fmt.Print("Masukkan penulis buku: ")
		fmt.Scan(&pustaka[i].penulis)

		fmt.Print("Masukkan penerbit buku: ")
		fmt.Scan(&pustaka[i].penerbit)

		fmt.Print("Masukkan jumlah eksemplar buku: ")
		fmt.Scan(&pustaka[i].eksemplar)

		fmt.Print("Masukkan tahun buku: ")
		fmt.Scan(&pustaka[i].tahun)

		fmt.Print("Masukkan rating buku: ")
		fmt.Scan(&pustaka[i].rating)
	}
}

// Prosedur 2: Cetak Terfavorit
func CetakTerfavorit(pustaka DaftarBuku, n int) {
	maxRating := -1
	var favorit Buku
	for i := 0; i < n; i++ {
		if pustaka[i].rating > maxRating {
			maxRating = pustaka[i].rating
			favorit = pustaka[i]
		}
	}
	fmt.Println("Buku terfavorit:")
	fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d\n", favorit.judul, favorit.penulis, favorit.penerbit, favorit.tahun)
}

// Prosedur 3: Urutkan Buku (Insertion Sort)
func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		key := pustaka[i]
		j := i - 1
		for j >= 0 && pustaka[j].rating < key.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = key
	}
}

// Prosedur 4: Cetak 5 Terbaru
func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	fmt.Println("5 Buku dengan rating tertinggi:")
	for i := 0; i < 5 && i < n; i++ {
		fmt.Printf("%d. Judul: %s, Rating: %d\n", i+1, pustaka[i].judul, pustaka[i].rating)
	}
}

// Prosedur 5: Cari Buku (Pencarian Biner)
func CariBuku(pustaka DaftarBuku, n int, r int) {
	low, high := 0, n-1
	for low <= high {
		mid := (low + high) / 2
		if pustaka[mid].rating == r {
			fmt.Printf("Buku dengan rating %d ditemukan:\n", r)
			fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Eksemplar: %d, Rating: %d\n",
				pustaka[mid].judul, pustaka[mid].penulis, pustaka[mid].penerbit, pustaka[mid].tahun, pustaka[mid].eksemplar, pustaka[mid].rating)
			return
		} else if pustaka[mid].rating < r {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	fmt.Println("Tidak ada buku dengan rating seperti itu.")
}

func main() {
	var pustaka DaftarBuku
	var n, ratingCari int

	// Meminta input jumlah buku
	fmt.Print("Masukkan jumlah data buku: ")
	fmt.Scan(&n)

	// Pastikan n > 0 dan tidak melebihi batas maksimum
	if n <= 0 || n > nMax {
		fmt.Println("Jumlah buku harus lebih besar dari 0 dan kurang dari 7919.")
		return
	}

	// Prosedur Daftarkan Buku
	DaftarkanBuku(&pustaka, n)

	// Prosedur Cetak Terfavorit
	CetakTerfavorit(pustaka, n)

	// Prosedur Urutkan Buku
	UrutBuku(&pustaka, n)

	// Prosedur Cetak 5 Terbaru
	Cetak5Terbaru(pustaka, n)

	// Prosedur Cari Buku
	fmt.Print("Masukkan rating buku yang ingin dicari: ")
	fmt.Scan(&ratingCari)
	CariBuku(pustaka, n, ratingCari)
}