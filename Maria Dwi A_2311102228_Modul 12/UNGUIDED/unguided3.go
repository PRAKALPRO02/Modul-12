package main

import "fmt"

type Buku struct {
	ID string
	Judul string
	Penulis string
	Penerbit string
	Eksemplar int
	Tahun int
	Rating int
}

const nMax int = 7919

type DaftarBuku [nMax]Buku



func cetakTerfavorit(pustaka DaftarBuku, n int) {
	if n== 0 {
		fmt.Println("Tidak ada buku di pustaka")
		return 
	}

	ratingTertinggi := 0

	for i := 1; i < n ; i++ {
		if pustaka[i].Rating > pustaka[ratingTertinggi].Rating {
			ratingTertinggi = i
		}
	}
	buku := pustaka[ratingTertinggi]
	fmt.Printf("Tervaforit: %s, %s, %s, %d\n", buku.Judul, buku.Penulis, buku.Penerbit, buku.Tahun)


}

func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		key := pustaka[i]
		j := i - 1

		for j >= 0 && pustaka[j].Rating < key.Rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = key
	}
}


func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	batas := 5

	if n < 5 {
		batas = 5
	}
	fmt.Println("5 Judul Buku dengan Rating Tertinggi: ")
	for i:= 0; i < batas; i++{
		fmt.Printf("%s, %s, %s, %d\n", pustaka[i].Judul, pustaka[i].Penulis, pustaka[i].Penerbit, pustaka[i].Tahun)
	}
}


func CariBuku(pustaka DaftarBuku, n int, rating int) {
	kiri := 0
	kanan := n-1

	for kiri <= kanan {
		tengah := (kiri + kanan) /2
		if pustaka[tengah].Rating == rating {
			buku := pustaka[tengah]
			fmt.Printf("Ditemukan : %s, %s, %s, %d\n", buku.Judul, buku.Penulis, buku.Penerbit, buku.Tahun)
			return
		} else if pustaka[tengah].Rating < rating {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}
	fmt.Println("Tidak ada buku dengan rating seperti itu. ")
}


func main() {
	var pustaka DaftarBuku
	var n int

	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data untuk buku ke-%d:\n", i + 1)

		fmt.Print("ID: ")
		fmt.Scanln(&pustaka[i].ID)

		fmt.Print("Judul: ")
		fmt.Scanln("%q\n", &pustaka[i].Judul)

		fmt.Print("Penulis: ")
		fmt.Scanln("%q\n",&pustaka[i].Penulis)

		fmt.Print("Penerbit: ")
		fmt.Scanln("%q\n",&pustaka[i].Penerbit)

		fmt.Print("Eksemplar: ")
		fmt.Scanln(&pustaka[i].Eksemplar)

		fmt.Print("Tahun: ")
		fmt.Scanln(&pustaka[i].Tahun)

		fmt.Print("Rating: ")
		fmt.Scanln(&pustaka[i].Rating)
	}

	var cariRating int
	fmt.Print("Masukkan rating buku yang ingin dicari: ")
	fmt.Scanln(&cariRating)

	fmt.Println("\nHasil : ")
	cetakTerfavorit(pustaka, n)
	UrutBuku(&pustaka,n)
	Cetak5Terbaru(pustaka, n)
	CariBuku(pustaka, n, cariRating)


}