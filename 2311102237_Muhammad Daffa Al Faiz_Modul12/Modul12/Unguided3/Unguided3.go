package main

import (
	"fmt"
)

const nmax = 7919

type Buku struct {
	ID, Judul, Penulis, Penerbit string
	Eksemplar, Tahun, Ranting    int
}
type DaftarBuku [nmax]Buku

func insertionSort(Pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		key := Pustaka[i]
		j := i - 1
		for j >= 0 && Pustaka[j].Ranting < key.Ranting {
			Pustaka[j+1] = Pustaka[j]
			j--
		}
		Pustaka[j+1] = key
	}
}

func findfavorit(Pustaka DaftarBuku, n int) Buku {
	max := Pustaka[0].Ranting
	favorit := Pustaka[0]
	for i := 1; i < n; i++ {
		if Pustaka[i].Ranting > max {
			max = Pustaka[i].Ranting
			favorit = Pustaka[i]
		}
	}
	return favorit
}

func limaRating(Pustaka *DaftarBuku, n int) []Buku {
	insertionSort(Pustaka, n)
	limit := 5
	if n < 5 {
		limit = n
	}
	return Pustaka[:limit]
}

func binarySearch(Pustaka DaftarBuku, n, target int) *Buku {
	insertionSort(&Pustaka, n)
	low, high := 0, n-1

	for low <= high {
		mid := (low + high) / 2
		if Pustaka[mid].Ranting == target {
			return &Pustaka[mid]
		} else if Pustaka[mid].Ranting < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return nil
}

func main() {
	var Buku DaftarBuku
	var n, cari int

	fmt.Print("Masukan Banyak Buku: ")
	fmt.Scan(&n)

	fmt.Println("Masukan (ID, Judul, Penulis, Penerbit, Eksemplar, Tahun, Rating)")
	for i := 0; i < n; i++ {
		fmt.Print("Masukan: ")
		fmt.Scan(&Buku[i].ID, &Buku[i].Judul, &Buku[i].Penulis, &Buku[i].Penerbit, &Buku[i].Eksemplar, &Buku[i].Tahun, &Buku[i].Ranting)
	}

	fmt.Print("Masukan Rating Buku Yang Anda Cari: ")
	fmt.Scan(&cari)

	favorit := findfavorit(Buku, n)
	fmt.Printf("Buku Terfavorit: %s - %s (Rating: %d)\n", favorit.Judul, favorit.Penulis, favorit.Ranting)

	topRated := limaRating(&Buku, n)
	fmt.Println("Lima Rating Tertinggi:")
	for i, buku := range topRated {
		fmt.Printf("%d. %s (Rating: %d)\n", i+1, buku.Judul, buku.Ranting)
	}

	temukan := binarySearch(Buku, n, cari)
	if temukan != nil {
		fmt.Printf("Buku dengan Rating %d ditemukan:\n", cari)
		fmt.Printf("ID: %s, Judul: %s, Penulis: %s\n", temukan.ID, temukan.Judul, temukan.Penulis)
	} else {
		fmt.Printf("Buku dengan Rating %d tidak ditemukan\n", cari)
	}
}
