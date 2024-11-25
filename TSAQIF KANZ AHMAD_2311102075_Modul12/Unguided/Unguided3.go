package main

import (
	"fmt"
	"sort"
)

type Buku struct {
	Judul, Penulis, Penerbit string
	Tahun, Rating            int
}

func DaftarkanBuku(n int) []Buku {
	pustaka := make([]Buku, n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data buku ke-%d (Judul Penulis Penerbit Tahun Rating):\n", i+1)
		fmt.Scan(&pustaka[i].Judul, &pustaka[i].Penulis, &pustaka[i].Penerbit, &pustaka[i].Tahun, &pustaka[i].Rating)
	}
	return pustaka
}

func CetakTefavorit(pustaka []Buku) {
	if len(pustaka) == 0 {
		fmt.Println("Tidak ada data buku.")
		return
	}
	tertinggi := pustaka[0]
	for _, buku := range pustaka {
		if buku.Rating > tertinggi.Rating {
			tertinggi = buku
		}
	}
	fmt.Println("Buku dengan rating tertinggi:")
	fmt.Printf("%s oleh %s, diterbitkan oleh %s (%d) - Rating: %d\n", tertinggi.Judul, tertinggi.Penulis, tertinggi.Penerbit, tertinggi.Tahun, tertinggi.Rating)
}

func UrutkanBuku(pustaka []Buku) {
	sort.Slice(pustaka, func(i, j int) bool {
		return pustaka[i].Rating > pustaka[j].Rating
	})
}

func Cetak5Terbaik(pustaka []Buku) {
	fmt.Println("5 Buku dengan rating tertinggi:")
	for i := 0; i < len(pustaka) && i < 5; i++ {
		fmt.Printf("%s oleh %s (%d) - Rating: %d\n", pustaka[i].Judul, pustaka[i].Penulis, pustaka[i].Tahun, pustaka[i].Rating)
	}
}

func CariBuku(pustaka []Buku, rating int) {
	for _, buku := range pustaka {
		if buku.Rating == rating {
			fmt.Println("Buku ditemukan:")
			fmt.Printf("%s oleh %s, diterbitkan oleh %s (%d) - Rating: %d\n", buku.Judul, buku.Penulis, buku.Penerbit, buku.Tahun, buku.Rating)
			return
		}
	}
	fmt.Println("Tidak ada buku dengan rating seperti itu.")
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Jumlah buku harus lebih besar dari nol.")
		return
	}
	pustaka := DaftarkanBuku(n)

	UrutkanBuku(pustaka)
	CetakTefavorit(pustaka)
	Cetak5Terbaik(pustaka)

	var rating int
	fmt.Print("Masukkan rating untuk mencari buku: ")
	fmt.Scan(&rating)
	CariBuku(pustaka, rating)
}
