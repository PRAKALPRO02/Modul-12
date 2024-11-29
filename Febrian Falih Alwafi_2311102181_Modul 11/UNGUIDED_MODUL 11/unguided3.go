package main

import "fmt"


const nMax = 7919


type Buku struct {
	ID, Judul, Penulis, Penerbit string
	Eksemplar, Tahun, Ranting    int
}


type DaftarBuku struct {
	Pustaka_2311102181  [nMax]Buku
	nPustaka int 
}


func sortBukuByRating(pustaka *DaftarBuku) {
	
	for i := 1; i < pustaka.nPustaka; i++ {
		key := pustaka.Pustaka_2311102181[i]
		j := i - 1

		
		for j >= 0 && pustaka.Pustaka_2311102181[j].Ranting < key.Ranting {
			pustaka.Pustaka_2311102181[j+1] = pustaka.Pustaka_2311102181[j]
			j--
		}
		pustaka.Pustaka_2311102181[j+1] = key
	}
}


func printFavoritBuku(pustaka DaftarBuku) {
	if pustaka.nPustaka == 0 {
		fmt.Println("Tidak ada buku.")
		return
	}

	
	favorit := pustaka.Pustaka_2311102181[0]
	for _, b := range pustaka.Pustaka_2311102181[:pustaka.nPustaka] {
		if b.Ranting > favorit.Ranting {
			favorit = b
		}
	}
	fmt.Printf("Buku Terfavorit: %s %s %s %s Tahun: %d\n", favorit.ID, favorit.Judul, favorit.Penulis, favorit.Penerbit, favorit.Tahun)
}


func printTop5Buku(pustaka DaftarBuku) {
	if pustaka.nPustaka < 5 {
		fmt.Println("Jumlah buku kurang dari 5, menampilkan semua buku.")
	}
	fmt.Println("5 Buku dengan Rating Tertinggi:")
	for i := 0; i < pustaka.nPustaka && i < 5; i++ {
		fmt.Printf("%s ", pustaka.Pustaka_2311102181[i].Judul)
	}
	fmt.Println()
}


func searchBukuByRating(pustaka DaftarBuku, rating int) *Buku {
	low, high := 0, pustaka.nPustaka-1
	for low <= high {
		mid := (low + high) / 2
		if pustaka.Pustaka_2311102181[mid].Ranting == rating {
			return &pustaka.Pustaka_2311102181[mid] 
		} else if pustaka.Pustaka_2311102181[mid].Ranting < rating {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return nil 
}

func main() {
	var pustaka DaftarBuku
	var n, cariRating int

	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(&n)

	fmt.Println("Masukkan data buku (ID, Judul, Penulis, Penerbit, Eksemplar, Tahun, Rating):")
	for i := 0; i < n; i++ {
		var b Buku
		fmt.Print("Masukkan : ")
		fmt.Scan(&b.ID, &b.Judul, &b.Penulis, &b.Penerbit, &b.Eksemplar, &b.Tahun, &b.Ranting)
		pustaka.Pustaka_2311102181[i] = b
	}

	
	pustaka.nPustaka = n

	fmt.Print("Masukkan rating buku yang Anda cari: ")
	fmt.Scan(&cariRating)

	sortBukuByRating(&pustaka)

	printFavoritBuku(pustaka)
	printTop5Buku(pustaka)

	temukan := searchBukuByRating(pustaka, cariRating)
	if temukan != nil {
		fmt.Printf("Buku dengan rating %d ditemukan: %s %s %s %s %d %d %d\n", cariRating, temukan.ID, temukan.Judul, temukan.Penulis, temukan.Penerbit, temukan.Eksemplar, temukan.Tahun, temukan.Ranting)
	} else {
		fmt.Printf("Buku dengan rating %d tidak ditemukan.\n", cariRating)
	}
}
