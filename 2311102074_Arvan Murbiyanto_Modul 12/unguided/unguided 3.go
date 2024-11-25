package main

import "fmt"

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

type DaftarBuku struct {
	pustaka []Buku
}

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		var buku Buku
		fmt.Printf("Masukkan data untuk buku ke-%d (id, judul, penulis, penerbit, eksemplar, tahun, rating):\n", i+1)
		fmt.Scan(&buku.id, &buku.judul, &buku.penulis, &buku.penerbit, &buku.eksemplar, &buku.tahun, &buku.rating)
		if len(pustaka.pustaka) < nMax {
			pustaka.pustaka = append(pustaka.pustaka, buku)
		}
	}
}

func CetakFavorit(pustaka DaftarBuku) {
	if len(pustaka.pustaka) == 0 {
		fmt.Println("Pustaka kosong.")
		return
	}
	terfavorit := pustaka.pustaka[0]
	for _, buku := range pustaka.pustaka {
		if buku.rating > terfavorit.rating {
			terfavorit = buku
		}
	}
	fmt.Println("\nBuku dengan rating tertinggi:")
	fmt.Printf("ID: %d, Judul: %s, Penulis: %s, Penerbit: %s, Eksemplar: %d, Tahun: %d, Rating: %d\n",
		terfavorit.id, terfavorit.judul, terfavorit.penulis, terfavorit.penerbit, terfavorit.eksemplar, terfavorit.tahun, terfavorit.rating)
}

func UrutkanBuku(pustaka *DaftarBuku) {
	for i := 1; i < len(pustaka.pustaka); i++ {
		key := pustaka.pustaka[i]
		j := i - 1
		for j >= 0 && pustaka.pustaka[j].rating < key.rating {
			pustaka.pustaka[j+1] = pustaka.pustaka[j]
			j--
		}
		pustaka.pustaka[j+1] = key
	}
}

func Cetak5Terbaik(pustaka DaftarBuku) {
	fmt.Println("\nLima buku dengan rating tertinggi:")
	for i := 0; i < 5 && i < len(pustaka.pustaka); i++ {
		buku := pustaka.pustaka[i]
		fmt.Printf("ID: %d, Judul: %s, Penulis: %s, Penerbit: %s, Eksemplar: %d, Tahun: %d, Rating: %d\n",
			buku.id, buku.judul, buku.penulis, buku.penerbit, buku.eksemplar, buku.tahun, buku.rating)
	}
}

func CariBuku(pustaka DaftarBuku, r int) {
	ditemukan := false
	for _, buku := range pustaka.pustaka {
		if buku.rating == r {
			fmt.Printf("\nDitemukan buku dengan rating %d:\n", r)
			fmt.Printf("ID: %d, Judul: %s, Penulis: %s, Penerbit: %s, Eksemplar: %d, Tahun: %d, Rating: %d\n",
				buku.id, buku.judul, buku.penulis, buku.penerbit, buku.eksemplar, buku.tahun, buku.rating)
			ditemukan = true
		}
	}
	if !ditemukan {
		fmt.Println("\nTidak ada buku dengan rating tersebut.")
	}
}

func main() {
	var pustaka DaftarBuku
	var n int

	fmt.Print("Masukkan jumlah buku di perpustakaan: ")
	fmt.Scan(&n)
	if n <= 0 || n > nMax {
		fmt.Println("Jumlah buku harus antara 1 hingga 7919.")
		return
	}

	DaftarkanBuku(&pustaka, n)
	CetakFavorit(pustaka)
	UrutkanBuku(&pustaka)
	Cetak5Terbaik(pustaka)

	var rating int
	fmt.Print("\nMasukkan rating buku yang ingin dicari: ")
	fmt.Scan(&rating)
	CariBuku(pustaka, rating)
}
