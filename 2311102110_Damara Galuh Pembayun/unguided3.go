package main

import "fmt"

const maxBuku = 7919

// Struktur data untuk mewakili sebuah buku
type Buku struct {
        id       string
        judul    string
        penulis  string
        penerbit string
        eksemplar int
        tahun     int
        rating    int
}

// Array untuk menyimpan daftar buku
type DaftarBuku [maxBuku]Buku

func main() {
        var pustaka DaftarBuku
        var n int // Jumlah buku

        // Input jumlah buku dan data buku
        fmt.Print("Masukkan jumlah buku: ")
        fmt.Scanln(&n)
        for i := 0; i < n; i++ {
                fmt.Printf("Buku ke-%d\n", i+1)
                fmt.Print("ID: ")
                fmt.Scanln(&pustaka[i].id)
                // ... (input data lainnya)
        }

        // Mengurutkan buku berdasarkan rating (descending)
        UrutBuku(&pustaka, n)

        // Menampilkan buku dengan rating tertinggi
        fmt.Println("\nBuku Terfavorit:")
        CetakTerfavorit(pustaka[n-1])

        // Menampilkan 5 buku dengan rating tertinggi
        fmt.Println("\n5 Buku dengan Rating Tertinggi:")
        Cetak5Terbaru(pustaka, n)

        // Mencari buku berdasarkan rating
        var ratingCari int
        fmt.Print("Masukkan rating buku yang ingin dicari: ")
        fmt.Scanln(&ratingCari)
        CariBuku(pustaka, n, ratingCari)
}

// Mengurutkan buku berdasarkan rating menggunakan Insertion Sort (descending)
func UrutBuku(pustaka *DaftarBuku, n int) {
        for i := 1; i < n; i++ {
                buku := (*pustaka)[i]
                j := i - 1
                for j >= 0 && (*pustaka)[j].rating < buku.rating {
                        (*pustaka)[j+1] = (*pustaka)[j]
                        j--
                }
                (*pustaka)[j+1] = buku
        }
}

// Menampilkan data buku
func CetakBuku(buku Buku) {
        fmt.Printf("ID: %s, Judul: %s, Penulis: %s, Penerbit: %s, Tahun: %d, Rating: %d\n",
                buku.id, buku.judul, buku.penulis, buku.penerbit, buku.tahun, buku.rating)
}

// Menampilkan buku dengan rating tertinggi
func CetakTerfavorit(buku Buku) {
        CetakBuku(buku)
}

// Menampilkan 5 buku dengan rating tertinggi
func Cetak5Terbaru(pustaka DaftarBuku, n int) {
        for i := 0; i < 5 && i < n; i++ {
                CetakBuku(pustaka[i])
        }
}

// Mencari buku berdasarkan rating menggunakan pencarian biner
func CariBuku(pustaka DaftarBuku, n int, rating int) {
        left, right := 0, n-1
        for left <= right {
                mid := (left + right) / 2
                if pustaka[mid].rating == rating {
                        CetakBuku(pustaka[mid])
                        return
                } else if pustaka[mid].rating < rating {
                        left = mid + 1
                } else {
                        right = mid - 1
                }
        }
        fmt.Printf("Tidak ada buku dengan rating %d\n", rating)
}