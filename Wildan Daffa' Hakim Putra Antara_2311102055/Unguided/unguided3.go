package main

import "fmt"

const nMax = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	fmt.Println("Format masukan (id, judul, penulis, penerbit, eksamplar, tahun, rating)")
	for i := 0; i < n; i++ {
		fmt.Print("Buku ke ", i+1, " : ")
		fmt.Scanln(&pustaka[i].id, &pustaka[i].judul, &pustaka[i].penulis, &pustaka[i].penerbit, &pustaka[i].eksemplar, &pustaka[i].tahun, &pustaka[i].rating)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	var indexFav = 0
	for i := 1; i < n; i++ {
		if pustaka[i].rating > pustaka[indexFav].rating {
			indexFav = i
		}
	}

	fmt.Printf("Buku terfavorit sekarang: : %s %s %s %s %d %d %d\n",
		pustaka[indexFav].id, pustaka[indexFav].judul, pustaka[indexFav].penerbit, pustaka[indexFav].penulis,
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
	fmt.Print(n, " buku rating tertinggi :")
	for i := 0; i < n; i++ {
		fmt.Print(" ", pustaka[i].judul)
	}
	fmt.Println()
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
		fmt.Printf("Buku dengan rating %d: %s %s %s %s %d %d %d\n",
			r, pustaka[med].id, pustaka[med].judul, pustaka[med].penerbit, pustaka[med].penulis,
			pustaka[med].eksemplar, pustaka[med].rating, pustaka[med].tahun)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var pustaka DaftarBuku
	var nPustaka, chooseRating int
	fmt.Print("Masukkan banyaknya buku : ")
	fmt.Scan(&nPustaka)
	DaftarkanBuku(&pustaka, nPustaka)
	fmt.Print("Masukkan rating buku yang dicari : ")
	fmt.Scan(&chooseRating)
	CetakTerfavorit(pustaka, nPustaka)
	UrutBuku(&pustaka, nPustaka)
	Cetak5Terbaru(pustaka, nPustaka)
	CariBuku(pustaka, nPustaka, chooseRating)
}
