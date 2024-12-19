//IKRAM IRIANSYAH
//2311102184

package main

import (
	"fmt"
)

type Buku struct {
	ID        string
	Judul     string
	Penulis   string
	Penerbit  string
	Eksemplar int
	Tahun     int
	Rating    int
}

type DaftarBuku struct {
	Pustaka  []Buku
	NPustaka int
}

func TambahBuku(pustaka *DaftarBuku) {
	var n int
	fmt.Print("Masukkan Jumlah Buku : ")
	fmt.Scan(&n)

	pustaka.NPustaka = n
	pustaka.Pustaka = make([]Buku, n)

	for i := 0; i < n; i++ {
		fmt.Printf("\nMasukkan Data Buku %d! \n", i+1)
		fmt.Print("ID : ")
		fmt.Scan(&pustaka.Pustaka[i].ID)
		fmt.Print("Judul : ")
		fmt.Scan(&pustaka.Pustaka[i].Judul)
		fmt.Print("Penulis : ")
		fmt.Scan(&pustaka.Pustaka[i].Penulis)
		fmt.Print("Penerbit : ")
		fmt.Scan(&pustaka.Pustaka[i].Penerbit)
		fmt.Print("Eksemplar : ")
		fmt.Scan(&pustaka.Pustaka[i].Eksemplar)
		fmt.Print("Tahun : ")
		fmt.Scan(&pustaka.Pustaka[i].Tahun)
		fmt.Print("Rating : ")
		fmt.Scan(&pustaka.Pustaka[i].Rating)
	}
}

func CetakBuku(pustaka DaftarBuku) {
	fmt.Println("\n= DATA BUKU = ")
	fmt.Printf("%-5s %-25s %-20s %-20s %-10s %-10s %-5s\n", "ID", "JUDUL", "PENULIS", "PENERBIT", "EKSEMPLAR", "TAHUN", "RATING")
	for _, buku := range pustaka.Pustaka {
		fmt.Printf("%-5s %-25s %-20s %-20s %-10d %-10d %-5d\n",
			buku.ID, buku.Judul, buku.Penulis, buku.Penerbit, buku.Eksemplar, buku.Tahun, buku.Rating)
	}
}

func UrutkanBuku(pustaka *DaftarBuku) {
	n := pustaka.NPustaka
	for i := 1; i < n; i++ {
		key := pustaka.Pustaka[i]
		j := i - 1
		for j >= 0 && pustaka.Pustaka[j].Rating < key.Rating {
			pustaka.Pustaka[j+1] = pustaka.Pustaka[j]
			j--
		}
		pustaka.Pustaka[j+1] = key
	}
}

func CetakBukuTeratas(pustaka DaftarBuku) {
	fmt.Println("\n= 5 BUKU TERATAS BERDASARKAN RATING = ")
	fmt.Printf("%-25s %-5s\n", "JUDUL", "RATING")
	for i := 0; i < 5 && i < pustaka.NPustaka; i++ {
		fmt.Printf("%-25s %-5d\n", pustaka.Pustaka[i].Judul, pustaka.Pustaka[i].Rating)
	}
}

func CariBuku(pustaka DaftarBuku) {
	var rating int
	fmt.Print("\nMasukan Rating :")
	fmt.Scan(&rating)

	left, right := 0, pustaka.NPustaka-1
	found := false
	for left <= right {
		mid := (left + right) / 2
		if pustaka.Pustaka[mid].Rating == rating {
			found = true
			buku := pustaka.Pustaka[mid]
			fmt.Printf("\n= BUKU DENGAN RATING %d = \n", rating)
			fmt.Printf("%-5s %-25s %-20s %-20s %-10s %-10s %-5s\n", "ID", "JUDUL", "PENULIS", "PENERBIT", "EKSEMPLAR", "TAHUN", "RATING")
			fmt.Printf("%-5s %-25s %-20s %-20s %-10d %-10d %-5d\n",
				buku.ID, buku.Judul, buku.Penulis, buku.Penerbit, buku.Eksemplar, buku.Tahun, buku.Rating)
			break
		} else if pustaka.Pustaka[mid].Rating < rating {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if !found {
		fmt.Println("\nTidak ada buku dengan rating tersebut.")
	}
}

func main() {
	var pustaka DaftarBuku

	TambahBuku(&pustaka)

	UrutkanBuku(&pustaka)

	CetakBuku(pustaka)

	CetakBukuTeratas(pustaka)

	CariBuku(pustaka)
}
