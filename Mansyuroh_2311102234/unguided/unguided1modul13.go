package main

import (
	"fmt"
)

const nMax = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku = [nMax]Buku

// Fungsi untuk mendaftarkan buku
func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("Buku %v: \n", i+1)
		fmt.Print("Masukkan id buku: ")
		fmt.Scan(&pustaka[i].id)
		fmt.Print("Masukkan judul buku: ")
		fmt.Scan(&pustaka[i].judul)
		fmt.Print("Masukkan penulis buku: ")
		fmt.Scan(&pustaka[i].penulis)
		fmt.Print("Masukkan penerbit buku: ")
		fmt.Scan(&pustaka[i].penerbit)
		fmt.Print("Masukkan eksemplar buku: ")
		fmt.Scan(&pustaka[i].eksemplar)
		fmt.Print("Masukkan tahun terbit: ")
		fmt.Scan(&pustaka[i].tahun)
		fmt.Print("Masukkan rating buku: ")
		fmt.Scan(&pustaka[i].rating)
		fmt.Println()
	}
}

// Fungsi untuk mencetak buku dengan rating tertinggi (favorit)
func CetakTerfavorit(pustaka DaftarBuku, n int) {
	max := pustaka[0]
	for i := 1; i < n; i++ {
		if pustaka[i].rating > max.rating {
			max = pustaka[i]
		}
	}
	fmt.Println("Buku Favorit: ")
	fmt.Printf("Judul buku: %v\n", max.judul)
	fmt.Printf("Penulis buku: %v\n", max.penulis)
	fmt.Printf("Penerbit buku: %v\n", max.penerbit)
	fmt.Println()
}

// Fungsi untuk mengurutkan buku berdasarkan rating
func UrutBuku(pustaka DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		key := pustaka[i]
		j := i - 1
		for j >= 0 && pustaka[j].rating > key.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = key
	}
}

// Fungsi untuk mencetak 5 buku terbaru dengan rating tertinggi
func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	// Urutkan berdasarkan rating tertinggi terlebih dahulu
	UrutBuku(pustaka, n)

	// Ambil 5 buku dengan rating tertinggi
	fmt.Println("5 Judul Buku dengan Rating Tertinggi: ")
	limit := 5
	if n < limit {
		limit = n
	}
	for i := 0; i < limit; i++ {
		fmt.Printf("Buku %v: \n", i+1)
		fmt.Printf("Judul buku: %v\n", pustaka[i].judul)
		fmt.Println()
	}
}

// Fungsi untuk mencari buku berdasarkan rating
func CariBuku(pustaka DaftarBuku, n, r int) {
	fmt.Print("Masukkan rating untuk mencari buku: ")
	fmt.Scan(&r)

	low, high := 0, n-1
	found := false
	for low <= high {
		mid := (low + high) / 2
		if pustaka[mid].rating == r {
			// Menampilkan buku dengan rating tersebut
			fmt.Println("Buku dengan rating", r, ":")
			i := mid
			for i >= 0 && pustaka[i].rating == r {
				fmt.Printf("Buku %v: \n", i+1)
				fmt.Printf("Judul buku: %v\n", pustaka[i].judul)
				fmt.Printf("Penulis buku: %v\n", pustaka[i].penulis)
				fmt.Printf("Penerbit buku: %v\n", pustaka[i].penerbit)
				fmt.Printf("Tahun terbit: %v\n", pustaka[i].tahun)
				fmt.Println()
				i--
			}

			i = mid + 1
			for i < n && pustaka[i].rating == r {
				fmt.Printf("Buku %v: \n", i+1)
				fmt.Printf("Judul buku: %v\n", pustaka[i].judul)
				fmt.Printf("Penulis buku: %v\n", pustaka[i].penulis)
				fmt.Printf("Penerbit buku: %v\n", pustaka[i].penerbit)
				fmt.Printf("Tahun terbit: %v\n", pustaka[i].tahun)
				fmt.Println()
				i++
			}
			found = true
			break
		} else if pustaka[mid].rating < r {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if !found {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var pustaka DaftarBuku
	var Npustaka, rating int

	// Input jumlah buku
	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(&Npustaka)

	// Menambahkan buku-buku ke pustaka
	DaftarkanBuku(&pustaka, Npustaka)
	CetakTerfavorit(pustaka, Npustaka)
	UrutBuku(pustaka, Npustaka)
	Cetak5Terbaru(pustaka, Npustaka)

	CariBuku(pustaka, Npustaka, rating)
}
