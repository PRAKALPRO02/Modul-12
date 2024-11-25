package main

import (
	"fmt"
	"math"
)

type Buku struct {
	id        int
	judul     string
	penulis   string
	penerbit  string
	eksemplar int
	tahun     int
	rating    int
}

func insertionSortBooks(pustaka []Buku) {
	n := len(pustaka)
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

func findHighestRatedBook(pustaka []Buku) Buku {
	highest := pustaka[0]
	for _, buku := range pustaka {
		if buku.rating > highest.rating {
			highest = buku
		}
	}
	return highest
}

func printTop5Books(pustaka []Buku) {
	fmt.Println("Lima buku dengan rating tertinggi:")
	for i := 0; i < int(math.Min(5, float64(len(pustaka)))); i++ {
		buku := pustaka[i]
		fmt.Printf("ID: %d, Judul: %s, Penulis: %s, Penerbit: %s, Eksemplar: %d, Tahun: %d, Rating: %d\n",
			buku.id, buku.judul, buku.penulis, buku.penerbit, buku.eksemplar, buku.tahun, buku.rating)
	}
}

func searchBooksByRating(pustaka []Buku, rating int) {
	found := false
	for _, buku := range pustaka {
		if buku.rating == rating {
			found = true
			fmt.Printf("ID: %d, Judul: %s, Penulis: %s, Penerbit: %s, Eksemplar: %d, Tahun: %d, Rating: %d\n",
				buku.id, buku.judul, buku.penulis, buku.penerbit, buku.eksemplar, buku.tahun, buku.rating)
		}
	}
	if !found {
		fmt.Println("Tidak ada buku dengan rating tersebut.")
	}
}

func main() {
	var pustaka []Buku
	var n int

	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Jumlah buku harus lebih dari 0.")
		return
	}

	for i := 0; i < n; i++ {
		var buku Buku
		fmt.Printf("Masukkan data untuk buku ke-%d (id, judul, penulis, penerbit, eksemplar, tahun, rating):\n", i+1)
		fmt.Scan(&buku.id, &buku.judul, &buku.penulis, &buku.penerbit, &buku.eksemplar, &buku.tahun, &buku.rating)
		pustaka = append(pustaka, buku)
	}

	highest := findHighestRatedBook(pustaka)
	fmt.Println("\nBuku dengan rating tertinggi:")
	fmt.Printf("ID: %d, Judul: %s, Penulis: %s, Penerbit: %s, Eksemplar: %d, Tahun: %d, Rating: %d\n",
		highest.id, highest.judul, highest.penulis, highest.penerbit, highest.eksemplar, highest.tahun, highest.rating)

	insertionSortBooks(pustaka)

	printTop5Books(pustaka)

	var rating int
	fmt.Print("\nMasukkan rating buku yang ingin dicari: ")
	fmt.Scan(&rating)
	searchBooksByRating(pustaka, rating)
}
