package main

import "fmt"

const maxBuku = 7919

type Buku struct {
	id        int
	judul     string
	penulis   string
	penerbit  string
	eksemplar int
	tahun     int
	rating    int
}

type KoleksiBuku struct {
	daftar []Buku
}

func TambahkanBuku(koleksi *KoleksiBuku, jumlah int) {
	for i := 0; i < jumlah; i++ {
		var buku Buku
		fmt.Printf("Masukkan data untuk buku ke-%d (id, judul, penulis, penerbit, eksemplar, tahun, rating):\n", i+1)
		fmt.Scan(&buku.id, &buku.judul, &buku.penulis, &buku.penerbit, &buku.eksemplar, &buku.tahun, &buku.rating)
		if len(koleksi.daftar) < maxBuku {
			koleksi.daftar = append(koleksi.daftar, buku)
		}
	}
}

func TampilkanFavorit(koleksi KoleksiBuku) {
	if len(koleksi.daftar) == 0 {
		fmt.Println("Koleksi kosong.")
		return
	}
	favorit := koleksi.daftar[0]
	for _, buku := range koleksi.daftar {
		if buku.rating > favorit.rating {
			favorit = buku
		}
	}
	fmt.Println("\nBuku dengan rating tertinggi:")
	fmt.Printf("ID: %d, Judul: %s, Penulis: %s, Penerbit: %s, Eksemplar: %d, Tahun: %d, Rating: %d\n",
		favorit.id, favorit.judul, favorit.penulis, favorit.penerbit, favorit.eksemplar, favorit.tahun, favorit.rating)
}

func SortirBuku(koleksi *KoleksiBuku) {
	n := len(koleksi.daftar)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if koleksi.daftar[j].rating < koleksi.daftar[j+1].rating {
				koleksi.daftar[j], koleksi.daftar[j+1] = koleksi.daftar[j+1], koleksi.daftar[j]
			}
		}
	}
}

func TampilkanTop5(koleksi KoleksiBuku) {
	fmt.Println("\nLima buku dengan rating tertinggi:")
	for i := 0; i < 5 && i < len(koleksi.daftar); i++ {
		buku := koleksi.daftar[i]
		fmt.Printf("ID: %d, Judul: %s, Penulis: %s, Penerbit: %s, Eksemplar: %d, Tahun: %d, Rating: %d\n",
			buku.id, buku.judul, buku.penulis, buku.penerbit, buku.eksemplar, buku.tahun, buku.rating)
	}
}

func CariBerdasarkanRating(koleksi KoleksiBuku, rating int) {
	ditemukan := false
	for _, buku := range koleksi.daftar {
		if buku.rating == rating {
			if !ditemukan {
				fmt.Printf("\nDitemukan buku dengan rating %d:\n", rating)
				ditemukan = true
			}
			fmt.Printf("ID: %d, Judul: %s, Penulis: %s, Penerbit: %s, Eksemplar: %d, Tahun: %d, Rating: %d\n",
				buku.id, buku.judul, buku.penulis, buku.penerbit, buku.eksemplar, buku.tahun, buku.rating)
		}
	}
	if !ditemukan {
		fmt.Println("\nTidak ada buku dengan rating tersebut.")
	}
}

func main() {
	var koleksi KoleksiBuku
	var jumlahBuku int

	fmt.Print("Masukkan jumlah buku di perpustakaan: ")
	fmt.Scan(&jumlahBuku)
	if jumlahBuku <= 0 || jumlahBuku > maxBuku {
		fmt.Println("Jumlah buku harus antara 1 hingga 7919.")
		return
	}

	TambahkanBuku(&koleksi, jumlahBuku)
	TampilkanFavorit(koleksi)
	SortirBuku(&koleksi)
	TampilkanTop5(koleksi)

	var rating int
	fmt.Print("\nMasukkan rating buku yang ingin dicari: ")
	fmt.Scan(&rating)
	CariBerdasarkanRating(koleksi, rating)
}
