package main

import "fmt"

const NMAX int = 9999

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku = [NMAX]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("[Buku %v]\n", i+1)
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
		fmt.Print("Masukkan tahun buku: ")
		fmt.Scan(&pustaka[i].tahun)
		fmt.Print("Masukkan rating buku: ")
		fmt.Scan(&pustaka[i].rating)
		fmt.Println()
	}
}

func CetakTerFavorit(pustaka DaftarBuku, n int) {
	max := pustaka[0]
	for i := 1; i < n; i++ {
		if pustaka[i].rating > max.rating {
			max = pustaka[i]
		}
	}

	fmt.Println("[Buku Favorit]")
	fmt.Printf("- Judul: %v\n", max.judul)
	fmt.Printf("- Penulis: %v\n", max.penulis)
	fmt.Printf("- Penerbit: %v\n", max.penerbit)
	fmt.Printf("- Tahun: %v\n", max.tahun)
	fmt.Printf("- Rating: %v\n", max.rating)
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	var temp Buku
	for i := 1; i < n; i++ {
		temp = pustaka[i]
		j := i - 1
		for j >= 0 && pustaka[j].rating > temp.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = temp
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	fmt.Printf("[%v Buku Favorit]\n", n)
	for i := 0; i < n && i < 5; i++ {
		fmt.Printf("%v. %v\n", i+1, pustaka[i].judul)
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	low, high := 0, n
	for low <= high {
		mid := (low + high) / 2
		if r == pustaka[mid].rating {
			fmt.Printf("[Buku dengan Rating %v]\n", r)
			fmt.Printf("- Judul: %v\n", pustaka[mid].judul)
			fmt.Printf("- Penulis: %v\n", pustaka[mid].penulis)
			fmt.Printf("- Penerbit: %v\n", pustaka[mid].penerbit)
			fmt.Printf("- Tahun: %v\n", pustaka[mid].tahun)
			fmt.Printf("- Eksemplar: %v\n", pustaka[mid].eksemplar)
			fmt.Printf("- Rating: %v\n", pustaka[mid].rating)
			return
		} else if r > pustaka[mid].rating {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	fmt.Println("Tidak ada buku dengan rating seperti itu.")
}

func main() {
	var pustaka DaftarBuku
	var nPustaka, rBuku int

	fmt.Print("Masukkan jumlah buku: ")
	fmt.Scan(&nPustaka)

	DaftarkanBuku(&pustaka, nPustaka)
	CetakTerFavorit(pustaka, nPustaka)
	UrutBuku(&pustaka, nPustaka)
	Cetak5Terbaru(pustaka, nPustaka)

	fmt.Print("Masukkan rating buku yang ingin dicari: ")
	fmt.Scan(&rBuku)

	CariBuku(pustaka, nPustaka, rBuku)
}
