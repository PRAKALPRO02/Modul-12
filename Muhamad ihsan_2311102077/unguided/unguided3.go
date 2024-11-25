package main

import "fmt"

type Buku struct {
	ID        int
	Judul     string
	Penulis   string
	Penerbit  string
	Eksemplar int
	Tahun     int
	Rating    int
}

type DaftarBuku struct {
	Pustaka []Buku
}

func DaftarkanBuku(n int) DaftarBuku {
	var pustaka DaftarBuku
	for i := 0; i < n; i++ {
		var buku Buku
		fmt.Scan(&buku.ID, &buku.Judul, &buku.Penulis, &buku.Penerbit, &buku.Eksemplar, &buku.Tahun, &buku.Rating)
		pustaka.Pustaka = append(pustaka.Pustaka, buku)
	}
	return pustaka
}

func CetakTerfavorit(pustaka DaftarBuku) {
	if len(pustaka.Pustaka) == 0 {
		fmt.Println("Tidak ada buku di perpustakaan.")
		return
	}
	terfavorit := pustaka.Pustaka[0]
	for i := 1; i < len(pustaka.Pustaka); i++ {
		if pustaka.Pustaka[i].Rating > terfavorit.Rating {
			terfavorit = pustaka.Pustaka[i]
		}
	}
	fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n",
		terfavorit.Judul, terfavorit.Penulis, terfavorit.Penerbit, terfavorit.Tahun, terfavorit.Rating)
}

func UrutBuku(pustaka *DaftarBuku) {
	n := len(pustaka.Pustaka)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if pustaka.Pustaka[j].Rating < pustaka.Pustaka[j+1].Rating {
				pustaka.Pustaka[j], pustaka.Pustaka[j+1] = pustaka.Pustaka[j+1], pustaka.Pustaka[j]
			}
		}
	}
}

func Cetak5Terbaru(pustaka DaftarBuku) {
	if len(pustaka.Pustaka) == 0 {
		fmt.Println("Tidak ada buku di perpustakaan.")
		return
	}
	limit := 5
	if len(pustaka.Pustaka) < 5 {
		limit = len(pustaka.Pustaka)
	}
	for i := 0; i < limit; i++ {
		buku := pustaka.Pustaka[i]
		fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n",
			buku.Judul, buku.Penulis, buku.Penerbit, buku.Tahun, buku.Rating)
	}
}

func CariBuku(pustaka DaftarBuku, rating int) {
	UrutBuku(&pustaka)
	low, high := 0, len(pustaka.Pustaka)-1
	found := false
	for low <= high {
		mid := (low + high) / 2
		if pustaka.Pustaka[mid].Rating == rating {
			found = true
			buku := pustaka.Pustaka[mid]
			fmt.Printf("Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Eksemplar: %d, Rating: %d\n",
				buku.Judul, buku.Penulis, buku.Penerbit, buku.Tahun, buku.Eksemplar, buku.Rating)
			break
		} else if pustaka.Pustaka[mid].Rating < rating {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if !found {
		fmt.Println("Tidak ada buku dengan rating tersebut.")
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	pustaka := DaftarkanBuku(n)
	CetakTerfavorit(pustaka)
	UrutBuku(&pustaka)
	Cetak5Terbaru(pustaka)
	var rating int
	fmt.Scan(&rating)
	CariBuku(pustaka, rating)
}
