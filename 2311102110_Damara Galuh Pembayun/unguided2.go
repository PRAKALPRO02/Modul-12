package main

import (
        "fmt"
)

func insertionSort(arr []int) {
        // Fungsi untuk mengurutkan array menggunakan metode Insertion Sort
        for i := 1; i < len(arr); i++ {
                key := arr[i]
                j := i - 1
                for j >= 0 && arr[j] > key {
                        arr[j+1] = arr[j]
                        j = j - 1
                }
                arr[j+1] = key
        }
}

func checkDistance(arr []int) string {
        // Fungsi untuk memeriksa apakah jarak antar elemen dalam array konsisten
        distance := arr[1] - arr[0]
        for i := 2; i < len(arr); i++ {
                if arr[i]-arr[i-1] != distance {
                        return "Data berjarak tidak tetap"
                }
        }
        return fmt.Sprintf("Data berjarak %d", distance)
}

func main() {
        // Contoh penggunaan
        var arr []int
        var input int

        fmt.Println("Masukkan bilangan bulat (akhiri dengan bilangan negatif):")
        for {
                _, err := fmt.Scan(&input)
                if err != nil || input < 0 {
                        break
                }
                arr = append(arr, input)
        }

        insertionSort(arr)
        fmt.Println(arr)
        fmt.Println(checkDistance(arr))
}