package main
import "fmt"

const nmax = 7919

type Buku struct {
	ID, Judul, Penulis, Penerbit string
	Eksemplar, Tahun, Ranting    int
}
type DaftarBuku [nmax]Buku

func insertionSort(Pustaka207 *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		key := Pustaka207[i]
		j := i - 1

		for j >= 0 && Pustaka207[j].Ranting < key.Ranting {
			Pustaka207[j+1] = Pustaka207[j]
			j--
		}
		Pustaka207[j+1] = key
	}
}

func CariFavorit(Pustaka207 DaftarBuku, n int) {
	max := Pustaka207[0].Ranting
	favorit := 0
	for i := 0; i < n; i++ {
		if Pustaka207[i].Ranting > max {
			max = Pustaka207[i].Ranting
			favorit = i
		}
	}
	fmt.Printf("Buku Terfavorit Adalah : %s %s %s %s %v \n", Pustaka207[favorit].ID, Pustaka207[favorit].Judul, Pustaka207[favorit].Penulis, Pustaka207[favorit].Penerbit, Pustaka207[favorit].Tahun)
}

func Rating(Pustaka207 DaftarBuku, n int) {
	insertionSort(&Pustaka207, n)
	fmt.Print("Lima Rating Tertinggi : ")
	for i := 0; i < n; i++ {
		fmt.Print(Pustaka207[i].Judul, " ")

	}
	fmt.Println()
}

func BinarySearch(Pustaka207 DaftarBuku, n, target int) int {
	low, high := 0, n-1

	for low <= high {
		mid := (low + high) / 2

		if Pustaka207[mid].Ranting == target {
			return mid 
		} else if Pustaka207[mid].Ranting < target {
			low = mid + 1 
		} else {
			high = mid - 1 
		}
	}

	return -1 
}
func main() {
	var Buku DaftarBuku
	var n, Cari int
	fmt.Print("Masukan Banyak Buku : ")
	fmt.Scan(&n)
	fmt.Println("Masukan (ID, Judul, Penulis, Penerbit, Eksemplar, Tahun, Rating)")
	for i := 0; i < n; i++ {
		fmt.Print("Masukan : ")
		fmt.Scan(&Buku[i].ID, &Buku[i].Judul, &Buku[i].Penulis, &Buku[i].Penerbit, &Buku[i].Eksemplar, &Buku[i].Tahun, &Buku[i].Ranting)
	}

	fmt.Print("Masukan Rating Buku Yang Anda Cari : ")
	fmt.Scan(&Cari)

	CariFavorit(Buku, n)
	Rating(Buku, n)
	Temukan := BinarySearch(Buku, n, Cari)
	if Temukan != -1 {
		fmt.Printf("Buku dengan Rating %v : %s %s %s %s %v %v %v\n", Cari, Buku[Temukan].ID, Buku[Temukan].Judul, Buku[Temukan].Penulis, Buku[Temukan].Penerbit, Buku[Temukan].Eksemplar, Buku[Temukan].Tahun, Buku[Temukan].Ranting)
	} else {
		fmt.Printf("Buku dengan Reting %v tidak ditemukan",Cari)
	}

}