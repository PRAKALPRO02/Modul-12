package main

import (
        "fmt"
        "sort"
)

func findMedian(data []int) float64 {
        // Fungsi ini menghitung median dari sebuah slice of int
        n := len(data)
        if n == 0 {
                return 0 // Jika data kosong, kembalikan 0
        }
        sort.Ints(data) // Urutkan data secara ascending
        if n%2 == 0 {
                // Jika jumlah data genap, ambil rata-rata dari dua nilai tengah
                return float64(data[n/2-1]+data[n/2]) / 2
        } else {
                // Jika jumlah data ganjil, ambil nilai tengah
                return float64(data[n/2])
        }
}

func main() {
        var data []int
        var input int
        for {
                fmt.Scan(&input)
                if input == -5313 {
                        break
                }
                if input == 0 {
                        median := findMedian(data)
                        fmt.Println(median)
                        data = []int{} // Kosongkan slice data untuk input berikutnya
                } else {
                        data = append(data, input)
                }
        }
}