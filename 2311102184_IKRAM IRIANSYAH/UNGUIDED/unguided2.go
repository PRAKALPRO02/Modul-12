//IKRAM IRIANSYAH
//2311102184

package main

import "fmt"

func sort(data []int) {
  n := len(data)
  for i := 0; i < n-1; i++ {
    minIdx := i
    for j := i + 1; j < n; j++ {
      if data[j] < data[minIdx] {
        minIdx = j
      }
    }
    // Tukar elemen
    data[i], data[minIdx] = data[minIdx], data[i]
  }
}

func median(data []int) int {
  n := len(data)
  if n%2 == 0 {
    return (data[n/2-1] + data[n/2]) / 2 
  }
  return data[n/2]
}

func main() {
  var input int
  var data []int
  var groups [][]int

  fmt.Println("Masukkan Data (akhiri dengan -5313) ")
  for {
    fmt.Scan(&input)
    if input == -5313 {
      break
    }
    if input == 0 {
      groups = append(groups, append([]int{}, data...)) 
    } else {
      data = append(data, input)
    }
  }

  for i, group := range groups {
    sort(group)
    median := median(group)
    fmt.Printf("Median %d : %d\n", i+1, median)
  }
}