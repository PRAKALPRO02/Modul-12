package main

import "fmt"

const nMax = 7919

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku []Buku

func DaftarkanBuku(Pustaka *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		var buku Buku
		fmt.Printf("Masukkan data buku ke-%d\n", i+1)
		fmt.Print("ID, Judul, Penulis, Penerbit, Eksemplar: ")
		fmt.Scan(&buku.id, &buku.judul, &buku.penulis, &buku.penerbit, &buku.eksemplar)
		fmt.Print("Tahun, Rating : ")
		fmt.Scan(&buku.tahun, &buku.rating)
		*Pustaka = append(*Pustaka, buku)
	}
}

func CetakTerFavorit(Pustaka DaftarBuku) {
	terFavorit := Pustaka[0]

	for _, buku := range Pustaka {
		if buku.rating > terFavorit.rating {
			terFavorit = buku
		}
	}
	fmt.Println("Buku dengan rating tertinggi ")
	fmt.Printf("Judul: %s\n", terFavorit.judul)
	fmt.Printf("Penulis: %s\n", terFavorit.penulis)
	fmt.Printf("Penerbit: %s\n", terFavorit.penerbit)
	fmt.Printf("Tahun: %d\n", terFavorit.tahun)
	fmt.Printf("Rating: %d\n", terFavorit.rating)
}

func UrutBuku(Pustaka *DaftarBuku) {
	n := len(*Pustaka)
	for i := 1; i < n; i++ {
		cur := (*Pustaka)[i]
		j := i - 1

		for j >= 0 && (*Pustaka)[j].rating < cur.rating {
			(*Pustaka)[j+1] = (*Pustaka)[j]
			j--
		}
		(*Pustaka)[j+1] = cur
	}
}

func Cetak5Terbaru(Pustaka DaftarBuku) {
	if len(Pustaka) < 5 {
		fmt.Println("Data pustaka kurang dari 5 buku.")
		return
	}

	fmt.Println("5 Buku dengan Rating terbaik")
	for i := 0; i < 5; i++ {
		buku := Pustaka[i]
		fmt.Printf("Judul: %s, Rating: %d\n", buku.judul, buku.rating)
	}
}

func CariBuku(Pustaka DaftarBuku, r int) {
	UrutBuku(&Pustaka)

	left, right := 0, len(Pustaka)-1
	for left <= right {
		mid := (left + right) / 2
		if Pustaka[mid].rating == r {
			fmt.Printf("Buku dengan Rating %d, dtemukan\n", r)
			fmt.Printf("Judul: %s\n", Pustaka[mid].judul)
			fmt.Printf("Penulis: %s\n", Pustaka[mid].penulis)
			fmt.Printf("Penerbit: %s\n", Pustaka[mid].penerbit)
			fmt.Printf("Tahun: %d\n", Pustaka[mid].tahun)
			fmt.Printf("Rating: %d\n", Pustaka[mid].rating)
			return
		} else if Pustaka[mid].rating < r {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

}

func main() {
	var Pustaka DaftarBuku
	var n, r int

	fmt.Print("Masukkan Jumlah Buku : ")
	fmt.Scan(&n)

	DaftarkanBuku(&Pustaka, n)
	CetakTerFavorit(Pustaka)
	UrutBuku(&Pustaka)
	Cetak5Terbaru(Pustaka)

	fmt.Print("Masukkan rating yang dicari: ")
	fmt.Scan(&r)
	CariBuku(Pustaka, r)
}
