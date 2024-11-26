package main

import (
	"bufio"
	"fmt"
	"os"
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

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	in := bufio.NewReader(os.Stdin)
	for i := 0; i < *n; i++ {
		var buku Buku
		fmt.Printf("_____________Masukkan data buku ke-%d_____________\n", i+1)
		buku.id = i + 1
		fmt.Scanln()
		fmt.Print("Masukkan Judul Buku: ")
		buku.judul, _ = in.ReadString('\n')
		fmt.Print("Masukkan Penulis Buku: ")
		buku.penulis, _ = in.ReadString('\n')
		fmt.Print("Masukkan Penerbit Buku: ")
		buku.penerbit, _ = in.ReadString('\n')
		fmt.Print("Masukkan Jumlah Eksemplar: ")
		fmt.Scan(&buku.eksemplar)
		fmt.Print("Masukkan Tahun Terbit: ")
		fmt.Scan(&buku.tahun)
		fmt.Print("Masukkan Rating Buku: ")
		fmt.Scan(&buku.rating)
		pustaka[i] = buku
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n == 0 {
		fmt.Println("Tidak ada data buku.")
		return
	}

	var maxRatingBuku Buku
	for i := 0; i < n; i++ {
		if pustaka[i].rating > maxRatingBuku.rating {
			maxRatingBuku = pustaka[i]
		}
	}

	fmt.Println("Buku Terfavorit:")
	fmt.Printf("Judul: %s\nPenulis: %s\nPenerbit: %s\nTahun: %d\n", maxRatingBuku.judul, maxRatingBuku.penulis, maxRatingBuku.penerbit, maxRatingBuku.tahun)
}

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

func Cetak5Terbesar(pustaka DaftarBuku, n int) {
	fmt.Println("5 Buku dengan Rating Tertinggi:")
	for i := 0; i < 5 && i < n; i++ {
		buku := pustaka[i]
		fmt.Printf("________Nomor ke-%d_______\n", i+1)
		fmt.Printf("Judul: %s\nRating: %d\n", buku.judul, buku.rating)
	}
}

func CariBuku(pustaka DaftarBuku, n int, rating int) {
	low := 0
	high := n - 1
	found := false

	for low <= high {
		mid := (low + high) / 2

		if pustaka[mid].rating == rating {
			fmt.Println("Buku Ditemukan:")
			fmt.Printf("Judul: %s\nPenulis: %s\nPenerbit: %s\nEksemplar: %d\nTahun: %d\nRating: %d\n",
				pustaka[mid].judul, pustaka[mid].penulis, pustaka[mid].penerbit, pustaka[mid].eksemplar, pustaka[mid].tahun, pustaka[mid].rating)
			found = true
			break
		} else if pustaka[mid].rating > rating {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if !found {
		fmt.Println("Tidak ada buku dengan rating tersebut.")
	}
}

func main() {
	var pustaka DaftarBuku
	var n int
	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(&n)
	DaftarkanBuku(&pustaka, &n)
	UrutBuku(&pustaka, n)
	CetakTerfavorit(pustaka, n)
	Cetak5Terbesar(pustaka, n)
	var rating int
	fmt.Print("Masukkan rating untuk mencari buku: ")
	fmt.Scan(&rating)
	CariBuku(pustaka, n, rating)
}
