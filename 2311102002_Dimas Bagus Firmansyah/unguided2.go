package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Baca jumlah data
	fmt.Println("Masukkan data (akhiri dengan -5313):")
	scanner.Scan()
	input := scanner.Text()

	// Parsing input menjadi slice
	data := strings.Fields(input)
	numbers := []int{}

	for _, v := range data {
		num, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Input tidak valid:", v)
			return
		}
		if num == -5313 {
			break
		}
		numbers = append(numbers, num)
	}

	// Proses data
	result := []int{}
	currentNumbers := []int{}

	for _, num := range numbers {
		if num == 0 {
			// Urutkan data yang sudah dibaca
			sort.Ints(currentNumbers)

			// Hitung median
			length := len(currentNumbers)
			if length%2 == 1 {
				// Jika jumlah data ganjil
				median := currentNumbers[length/2]
				result = append(result, median)
			} else {
				// Jika jumlah data genap
				median := (currentNumbers[length/2-1] + currentNumbers[length/2]) / 2
				result = append(result, median)
			}
		} else {
			// Tambahkan data baru
			currentNumbers = append(currentNumbers, num)
		}
	}

	// Cetak hasil
	fmt.Println("Keluaran:")
	for _, res := range result {
		fmt.Println(res)
	}
}
