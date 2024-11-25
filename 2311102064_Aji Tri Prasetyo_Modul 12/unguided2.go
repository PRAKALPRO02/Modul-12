package main

import "fmt"

func hitungMedian(arr []int) int {
    n := len(arr)
    if n%2 == 1 {
        return arr[n/2]
    }
    return (arr[n/2-1] + arr[n/2]) / 2
}

func insertionSort(arr []int) {
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

func main() {
    var data []int
    var input int

    for {
        fmt.Scan(&input)
        if input == -5313 {
            break
        }

        if input == 0 {
            insertionSort(data)
            median := hitungMedian(data)
            fmt.Println(median)
        } else {
            data = append(data, input)
        }
    }
}