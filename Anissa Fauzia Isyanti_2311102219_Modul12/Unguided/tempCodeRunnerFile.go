package main

import "fmt"

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

type DaftarBuku struct {
	pustaka []Buku
}

func DaftarkanBuku(pustaka *DaftarBuku, buku Buku) {
	if len(pustaka.pustaka) < nMax {
		pustaka.pustaka = append(pustaka.pustaka, buku)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) Buku {
	if n == 0 {
		fmt.Println("Tidak ada buku dalam pustaka.")
		return Buku{}
	}

	terfavorit := pustaka.pustaka[0]
	for i := 1; i < n; i++ {
		if pustaka.pustaka[i].rating > terfavorit.rating {
			terfavorit = pustaka.pustaka[i]
		}
	}

	return terfavorit
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		key := pustaka.pustaka[i]
		j := i - 1
		for j >= 0 && pustaka.pustaka[j].rating < key.rating {
			pustaka.pustaka[j+1] = pustaka.pustaka[j]
			j--
		}
		pustaka.pustaka[j+1] = key
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	if n == 0 {
		fmt.Println("Tidak ada buku dalam pustaka.")
		return
	}

	fmt.Println("5 Buku Dengan Rating Tertinggi:")
	for i := 0; i < 5 && i < n; i++ {
		buku := pustaka.pustaka[i]
		fmt.Printf("Judul   : %s\nRating  : %d\n", buku.judul, buku.rating)
	}
}

func CariBuku(pustaka DaftarBuku, n, r int) {
	UrutBuku(&pustaka, n)

	low, high := 0, n-1
	for low <= high {
		mid := (low + high) / 2
		if pustaka.pustaka[mid].rating == r {
			buku := pustaka.pustaka[mid]
			fmt.Printf("Buku ditemukan:\nJudul       : %s\nPenulis     : %s\nPenerbit    : %s\nTahun       : %d\nEksemplar   : %d\nRating      : %d\n",
				buku.judul, buku.penulis, buku.penerbit, buku.tahun, buku.eksemplar, buku.rating)
			return
		} else if pustaka.pustaka[mid].rating < r {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	fmt.Println("Tidak ada buku dengan rating seperti itu.")
}

func main() {
	var pustaka DaftarBuku
	var n int

	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		var id, eksemplar, tahun, rating int
		var judul, penulis, penerbit string

		fmt.Printf("Masukkan data buku ke-%d:\n", i+1)
		fmt.Print("ID           : ")
		fmt.Scanln(&id)
		fmt.Print("Judul        : ")
		fmt.Scanln(&judul)
		fmt.Print("Penulis      : ")
		fmt.Scanln(&penulis)
		fmt.Print("Penerbit     : ")
		fmt.Scanln(&penerbit)
		fmt.Print("Eksemplar    : ")
		fmt.Scanln(&eksemplar)
		fmt.Print("Tahun        : ")
		fmt.Scanln(&tahun)
		fmt.Print("Rating       : ")
		fmt.Scanln(&rating)

		buku := Buku{id, judul, penulis, penerbit, eksemplar, tahun, rating}
		DaftarkanBuku(&pustaka, buku)
		fmt.Println()
	}

	terfavorit := CetakTerfavorit(pustaka, n)
	fmt.Println("Buku Terfavorit:")
	fmt.Printf("Judul       : %s\nPenulis     : %s\nPenerbit    : %s\nTahun       : %d\nRating      : %d\n",
		terfavorit.judul, terfavorit.penulis, terfavorit.penerbit, terfavorit.tahun, terfavorit.rating)

	UrutBuku(&pustaka, n)
	Cetak5Terbaru(pustaka, n)

	var targetRating int
	fmt.Print("\nMasukkan rating buku yang ingin dicari: ")
	fmt.Scanln(&targetRating)
	CariBuku(pustaka, n, targetRating)
}
