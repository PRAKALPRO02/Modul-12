package main

import "fmt"

type arrayDaerah [1000]int

func selectionSort(array * arrayDaerah, n int) {
    var ganjil arrayDaerah
    var genap arrayDaerah
    var jumlahGanjil, jumlahGenap int

    for i:= 0; i < n; i++ {
        if array[i]%2 != 0{
            ganjil[jumlahGanjil] = array[i]
            jumlahGanjil++
        } else {
            genap[jumlahGenap] = array[i]
            jumlahGenap++
        }
    }

    for i:= 0; i < jumlahGanjil; i++ {
        minimal := i 
        for j := i + 1; j < jumlahGanjil; j++{
            if ganjil[minimal] > ganjil[j] {
                minimal = j
            }
        }
        ganjil[i], ganjil[minimal] = ganjil[minimal], ganjil[i]
    }

    for i := 0; i < jumlahGenap; i++ {
        maksimal := i 
        for j := i + 1; j < jumlahGenap; j++ {
            if genap[maksimal] < genap[j] {
                maksimal = j
            }
        }
        genap[i], genap[maksimal] = genap[maksimal], genap[i]
    }

    index := 0
    for i:= 0; i < jumlahGanjil; i++ {
        array[index] = ganjil[i]
        index++
    }

    for i := 0; i < jumlahGenap; i++ {
        array[index] = genap[i]
        index++
    }
}

func main() {
    var n int
    fmt.Print("Masukkan jumlah daerah (n): ")
    fmt.Scan(&n)

    if n <= 0 || n > 1000{
        fmt.Println("n Harus lebih besar dari 0 dan kurang dari 1000!")
        return 
    }

    for i := 0; i < n; i++ {
        var jumlah int
        fmt.Printf("\nMasukkan jumlah kerabat untuk daerah ke-%d: ", i+1)
        fmt.Scan(&jumlah)

        if jumlah <= 0 || jumlah > 1000000 {
            fmt.Println("Jumlah rumah harus lebih besar dari 0 dan kurang dari 1000000!")
            return 
        }
        var rumah arrayDaerah 
        fmt.Printf("Masukkan nomor rumah kerabat daerah ke-%d: ", i+1)

        for j := 0; j < jumlah; j++ {
            fmt.Scan(&rumah[j])
        }
        selectionSort(&rumah, jumlah)

        fmt.Printf("Hasil urutan rumah untuk daerah ke-%d: ", i + 1)
        for k := 0; k < jumlah; k++ {
            fmt.Printf("%d ", rumah[k])
        }
        fmt.Println()
    }
}


