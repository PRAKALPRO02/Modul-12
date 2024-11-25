package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const nMax = 7019

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

type DaftarBuku [nMax]Buku

var pustaka DaftarBuku

func daftarkanBuku(pustaka *DaftarBuku, n int) {
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < n; i++ {
		scanner.Scan()
		data := strings.Split(scanner.Text(), ",")
		for j := range data {
			data[j] = strings.TrimSpace(data[j])
		}
		pustaka[i] = Buku{
			id:       data[0],
			judul:    data[1],
			penulis:  data[2],
			penerbit: data[3],
			eksemplar: func() int {
				val, _ := strconv.Atoi(data[4])
				return val
			}(),
			tahun: func() int {
				val, err := strconv.Atoi(data[5])
				if err != nil {
					fmt.Println("Error parsing year:", err)
				}
				return val
			}(),
			rating: func() int {
				val, _ := strconv.Atoi(data[6])
				return val
			}(),
		}
	}
}

func cetakTerfavorit(pustaka *DaftarBuku, n int) {
	var maxRating = -1
	var favoritIndeks = -1
	for i := 0; i < n; i++ {
		if pustaka[i].rating > maxRating {
			maxRating = pustaka[i].rating
			favoritIndeks = i
		}
	}
	fmt.Printf("\nBuku Terfavorit: %s, %s, %s, %d\n", pustaka[favoritIndeks].judul, pustaka[favoritIndeks].penulis, pustaka[favoritIndeks].penerbit, pustaka[favoritIndeks].tahun)
}

func urutBuku(pustaka *DaftarBuku, n int) {
	sort.Slice(pustaka[:n], func(i, j int) bool {
		return pustaka[i].rating > pustaka[j].rating
	})
}

func cetak5Terbaru(pustaka *DaftarBuku, n int) {
	batas := 5
	if n < batas {
		batas = n
	}
	fmt.Println("\n5 Judul Buku dengan rating tertinggi:")
	for i := 0; i < batas; i++ {
		fmt.Printf("%d. %s (%d)\n", i+1, pustaka[i].judul, pustaka[i].rating)
	}

}

func cariBuku(pustaka *DaftarBuku, n int, rating int) {
	ketemu := false
	no := 1
	for i := 0; i < n; i++ {
		if pustaka[i].rating == rating {
			fmt.Printf("%d. %s, %s, %s, %d, %d, %d\n", no, pustaka[i].judul, pustaka[i].penulis, pustaka[i].penerbit, pustaka[i].eksemplar, pustaka[i].tahun, pustaka[i].rating)
			ketemu = true
			no++
		}
	}
	if !ketemu {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan jumlah buku: ")
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Masukkan data buku (id, judul, penulis, penerbit, eksemplar, tahun, rating):")
	daftarkanBuku(&pustaka, N)

	urutBuku(&pustaka, N)
	cetakTerfavorit(&pustaka, N)
	cetak5Terbaru(&pustaka, N)

	fmt.Print("\nMasukkan rating buku yang dicari: ")
	scanner.Scan()

	ratingCari, _ := strconv.Atoi(scanner.Text())

	cariBuku(&pustaka, N, ratingCari)
}
