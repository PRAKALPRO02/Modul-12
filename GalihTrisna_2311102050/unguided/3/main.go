package main

import "fmt"

const nMax = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	fmt.Println("Masukkan informasi buku (id, judul, penulis, penerbit, eksamplar, tahun, rating):")
	for i := 0; i < n; i++ {
		fmt.Printf("Buku ke-%d: ", i+1)
		fmt.Scan(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis, &pustaka[i].penerbit, &pustaka[i].eksemplar, &pustaka[i].tahun, &pustaka[i].rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	var indexFav = 0
	for i := 1; i < n; i++ {
		if pustaka[i].rating > pustaka[indexFav].rating {
			indexFav = i
		}
	}

	fmt.Printf("Buku terfavorit saat ini: %s | %s | %s | %s | Eksemplar: %d | Rating: %d | Tahun: %d\n",
		pustaka[indexFav].id, pustaka[indexFav].judul, pustaka[indexFav].penulis, pustaka[indexFav].penerbit,
		pustaka[indexFav].eksemplar, pustaka[indexFav].rating, pustaka[indexFav].tahun)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		var key Buku = pustaka[i]
		var j int = i - 1

		for j >= 0 && pustaka[j].rating < key.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = key
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	if n > 5 {
		n = 5
	}
	fmt.Printf("5 Buku dengan rating tertinggi:\n")
	for i := 0; i < n; i++ {
		fmt.Printf("%d. %s (Rating: %d)\n", i+1, pustaka[i].judul, pustaka[i].rating)
	}
}

func CariBuku(pustaka DaftarBuku, n, r int) {
	kr := 0
	kn := n - 1
	var med int
	var found bool = false

	for kr <= kn && !found {
		med = (kr + kn) / 2

		if pustaka[med].rating > r {
			kr = med + 1
		} else if pustaka[med].rating < r {
			kn = med - 1
		} else {
			found = true
		}
	}
	if found {
		fmt.Printf("Buku dengan rating %d ditemukan: %s | %s | %s | %s | Eksemplar: %d | Rating: %d | Tahun: %d\n",
			r, pustaka[med].id, pustaka[med].judul, pustaka[med].penulis, pustaka[med].penerbit,
			pustaka[med].eksemplar, pustaka[med].rating, pustaka[med].tahun)
	} else {
		fmt.Printf("Tidak ada buku dengan rating %d.\n", r)
	}
}

func main() {
	var pustaka DaftarBuku
	var nPustaka, chooseRating int
	fmt.Print("Masukkan jumlah buku yang akan didaftarkan: ")
	fmt.Scan(&nPustaka)
	DaftarkanBuku(&pustaka, nPustaka)
	fmt.Print("Masukkan rating buku yang ingin dicari: ")
	fmt.Scan(&chooseRating)
	fmt.Println("\nMenampilkan buku dengan rating tertinggi...")
	CetakTerfavorit(pustaka, nPustaka)
	fmt.Println("\nMengurutkan buku berdasarkan rating...")
	UrutBuku(&pustaka, nPustaka)
	fmt.Println("\nMenampilkan 5 buku dengan rating tertinggi:")
	Cetak5Terbaru(pustaka, nPustaka)
	fmt.Printf("\nMencari buku dengan rating %d...\n", chooseRating)
	CariBuku(pustaka, nPustaka, chooseRating)
}
