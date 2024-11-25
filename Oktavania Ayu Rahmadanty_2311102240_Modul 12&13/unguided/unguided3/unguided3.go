package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Buku struct {
	ID        int
	Judul     string
	Penulis   string
	Penerbit  string
	Eksemplar int
	Tahun     int
	Rating    int
}

func main() {
	var n int

	// Membaca input dari pengguna
	scanner := bufio.NewScanner(os.Stdin)

	// Memasukkan jumlah buku
	fmt.Print("Masukkan jumlah data buku: ")
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())

	// Membuat slice untuk menyimpan data buku
	buku := make([]Buku, n)

	// Memasukkan data buku
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data buku ke-%d (ID, Judul, Penulis, Penerbit, Eksemplar, Tahun, Rating):\n", i+1)

		// Membaca input per baris
		scanner.Scan()
		id, _ := strconv.Atoi(scanner.Text())

		scanner.Scan()
		judul := scanner.Text()

		scanner.Scan()
		penulis := scanner.Text()

		scanner.Scan()
		penerbit := scanner.Text()

		scanner.Scan()
		eksemplar, _ := strconv.Atoi(scanner.Text())

		scanner.Scan()
		tahun, _ := strconv.Atoi(scanner.Text())

		scanner.Scan()
		rating, _ := strconv.Atoi(scanner.Text())

		// Menyimpan data buku ke dalam slice
		buku[i] = Buku{id, judul, penulis, penerbit, eksemplar, tahun, rating}
	}

	// Menentukan buku terfavorit
	var favorit Buku
	for _, b := range buku {
		if b.Rating > favorit.Rating {
			favorit = b
		}
	}
	fmt.Println("Buku Terfavorit:")
	fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d\n", favorit.Judul, favorit.Penulis, favorit.Penerbit, favorit.Tahun)

	// Mengurutkan buku berdasarkan rating (descending)
	for i := 0; i < len(buku); i++ {
		for j := i + 1; j < len(buku); j++ {
			if buku[j].Rating > buku[i].Rating {
				buku[i], buku[j] = buku[j], buku[i]
			}
		}
	}

	// Menampilkan lima buku dengan rating tertinggi
	fmt.Println("5 Buku dengan Rating Tertinggi:")
	for i := 0; i < 5 && i < len(buku); i++ {
		fmt.Printf("%d. Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d\n", i+1, buku[i].Judul, buku[i].Penulis, buku[i].Penerbit, buku[i].Tahun)
	}

	// Mencari buku berdasarkan rating
	fmt.Print("Masukkan rating untuk mencari buku: ")
	scanner.Scan()
	cariRating, _ := strconv.Atoi(scanner.Text())

	found := false
	for _, b := range buku {
		if b.Rating == cariRating {
			fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d\n", b.Judul, b.Penulis, b.Penerbit, b.Tahun)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("Tidak ada buku dengan rating seperti itu.")
	}
}
