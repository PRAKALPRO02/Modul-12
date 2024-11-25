package main

import (
	"fmt"
)

// Definisi konstanta
const nMax = 7919

// Definisi struct Buku
type Buku struct {
	id       int
	judul    string
	penulis  string
	penerbit string
	tahun    int
	exemplar int
	rating   int
}

// Definisi struct DaftarBuku
type DaftarBuku struct {
	buku  [nMax]Buku
	nBuku int
}

// Fungsi untuk mendaftarkan buku
func DaftarkanBuku(perpus *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		var id, tahun, exemplar, rating int
		var judul, penulis, penerbit string

		fmt.Printf("Masukkan data buku ke-%d:\n", i+1)
		fmt.Print("ID: ")
		fmt.Scan(&id)
		fmt.Print("Judul: ")
		fmt.Scan(&judul)
		fmt.Print("Penulis: ")
		fmt.Scan(&penulis)
		fmt.Print("Penerbit: ")
		fmt.Scan(&penerbit)
		fmt.Print("Tahun: ")
		fmt.Scan(&tahun)
		fmt.Print("Exemplar: ")
		fmt.Scan(&exemplar)
		fmt.Print("Rating: ")
		fmt.Scan(&rating)

		perpus.buku[i] = Buku{id, judul, penulis, penerbit, tahun, exemplar, rating}
	}
	perpus.nBuku = n
}

// Fungsi untuk mencetak buku dengan rating tertinggi
func CetakTerfavorit(perpus *DaftarBuku) {
	if perpus.nBuku == 0 {
		fmt.Println("Tidak ada buku dalam perpustakaan.")
		return
	}

	// Cari buku dengan rating tertinggi
	maxRating := -1
	var terbaik Buku
	for i := 0; i < perpus.nBuku; i++ {
		if perpus.buku[i].rating > maxRating {
			maxRating = perpus.buku[i].rating
			terbaik = perpus.buku[i]
		}
	}

	fmt.Println("Buku dengan rating tertinggi:")
	fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n",
		terbaik.judul, terbaik.penulis, terbaik.penerbit, terbaik.tahun, terbaik.rating)
}

// Fungsi untuk mengurutkan buku berdasarkan rating menggunakan Insertion Sort
func UrutkanBuku(perpus *DaftarBuku) {
	for i := 1; i < perpus.nBuku; i++ {
		key := perpus.buku[i]
		j := i - 1

		for j >= 0 && perpus.buku[j].rating < key.rating {
			perpus.buku[j+1] = perpus.buku[j]
			j--
		}
		perpus.buku[j+1] = key
	}
}

// Fungsi untuk mencari buku berdasarkan rating
func CariBuku(perpus *DaftarBuku, targetRating int) {
	fmt.Printf("Mencari buku dengan rating: %d\n", targetRating)
	ditemukan := false
	for i := 0; i < perpus.nBuku; i++ {
		if perpus.buku[i].rating == targetRating {
			fmt.Printf("Buku ditemukan: Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d\n",
				perpus.buku[i].judul, perpus.buku[i].penulis, perpus.buku[i].penerbit, perpus.buku[i].tahun)
			ditemukan = true
		}
	}

	if !ditemukan {
		fmt.Println("Tidak ada buku dengan rating tersebut.")
	}
}

func main() {
	var perpustakaan DaftarBuku
	var n int

	fmt.Print("Masukkan jumlah buku yang akan didaftarkan: ")
	fmt.Scan(&n)

	DaftarkanBuku(&perpustakaan, n)
	fmt.Println()

	fmt.Println("Mengurutkan buku berdasarkan rating...")
	UrutkanBuku(&perpustakaan)
	fmt.Println("Daftar buku setelah diurutkan berdasarkan rating:")
	for i := 0; i < perpustakaan.nBuku; i++ {
		fmt.Printf("Judul: %s, Rating: %d\n", perpustakaan.buku[i].judul, perpustakaan.buku[i].rating)
	}
	fmt.Println()

	CetakTerfavorit(&perpustakaan)

	var targetRating int
	fmt.Print("Masukkan rating buku yang ingin dicari: ")
	fmt.Scan(&targetRating)
	CariBuku(&perpustakaan, targetRating)
}
