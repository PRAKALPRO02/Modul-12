package main

import (
	"fmt"
	"sort"
)

type Buku struct {
	ID        int
	Judul     string
	Penulis   string
	Penerbit  string
	Eksemplar int
	Tahun     int
	Rating    int
}

type DaftarBuku struct {
	Pustaka  []Buku
	nPustaka int
}

// Fungsi untuk memasukkan daftar buku
func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	pustaka.nPustaka = n
	pustaka.Pustaka = make([]Buku, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data buku ke-%d (ID, Judul, Penulis, Penerbit, Eksemplar, Tahun, Rating):\n", i+1)
		fmt.Scanf("%d\n", &pustaka.Pustaka[i].ID)
		fmt.Scanln(&pustaka.Pustaka[i].Judul)
		fmt.Scanln(&pustaka.Pustaka[i].Penulis)
		fmt.Scanln(&pustaka.Pustaka[i].Penerbit)
		fmt.Scanf("%d %d %d\n", &pustaka.Pustaka[i].Eksemplar, &pustaka.Pustaka[i].Tahun, &pustaka.Pustaka[i].Rating)
	}
}

// Fungsi untuk mencetak buku dengan rating tertinggi
func CetakTerfavorit(pustaka DaftarBuku) {
	var favorit Buku
	for _, buku := range pustaka.Pustaka {
		if buku.Rating > favorit.Rating {
			favorit = buku
		}
	}
	fmt.Println("Buku Terfavorit:")
	fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d\n", favorit.Judul, favorit.Penulis, favorit.Penerbit, favorit.Tahun)
}

// Fungsi untuk mengurutkan buku berdasarkan rating (descending)
func UrutBuku(pustaka *DaftarBuku) {
	sort.SliceStable(pustaka.Pustaka, func(i, j int) bool {
		return pustaka.Pustaka[i].Rating > pustaka.Pustaka[j].Rating
	})
}

// Fungsi untuk mencetak 5 buku dengan rating tertinggi
func CetakTerbaru(pustaka DaftarBuku) {
	fmt.Println("5 Buku dengan Rating Tertinggi:")
	for i := 0; i < 5 && i < pustaka.nPustaka; i++ {
		buku := pustaka.Pustaka[i]
		fmt.Printf("%d. Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d\n",
			i+1, buku.Judul, buku.Penulis, buku.Penerbit, buku.Tahun)
	}
}

// Fungsi untuk mencari buku berdasarkan rating tertentu
func CariBuku(pustaka DaftarBuku, rating int) {
	fmt.Printf("Buku dengan rating %d:\n", rating)
	found := false
	for _, buku := range pustaka.Pustaka {
		if buku.Rating == rating {
			fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d\n",
				buku.Judul, buku.Penulis, buku.Penerbit, buku.Tahun)
			found = true
		}
	}
	if !found {
		fmt.Println("Tidak ada buku dengan rating tersebut.")
	}
}

func main() {
	var daftarBuku DaftarBuku
	var n, cariRating int

	// Input jumlah buku
	fmt.Println("Masukkan jumlah buku:")
	fmt.Scanf("%d\n", &n)

	// Daftarkan buku
	DaftarkanBuku(&daftarBuku, n)

	// Cetak buku terfavorit
	CetakTerfavorit(daftarBuku)

	// Urutkan buku berdasarkan rating
	UrutBuku(&daftarBuku)

	// Cetak 5 buku dengan rating tertinggi
	CetakTerbaru(daftarBuku)

	// Cari buku berdasarkan rating
	fmt.Println("Masukkan rating untuk mencari buku:")
	fmt.Scanf("%d\n", &cariRating)
	CariBuku(daftarBuku, cariRating)
}

// Zahrina Antika Malahati_2311102109
