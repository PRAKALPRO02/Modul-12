package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var numbers []int
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Masukkan angka (akhiri dengan -5313):")
	for {
		scanner.Scan()
		input := scanner.Text()

		// Konversi input menjadi angka
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Masukkan angka valid.")
			continue
		}

		// Jika input adalah -5313, hentikan program
		if num == -5313 {
			break
		}

		// Jika input adalah 0, hitung dan cetak median
		if num == 0 {
			if len(numbers) == 0 {
				fmt.Println("Belum ada angka untuk dihitung median.")
				continue
			}
			// Urutkan data
			sort.Ints(numbers)

			// Hitung median
			mid := len(numbers) / 2
			var median int
			if len(numbers)%2 == 0 {
				median = (numbers[mid-1] + numbers[mid]) / 2
			} else {
				median = numbers[mid]
			}
			fmt.Println("Median saat ini:", median)
		} else {
			// Tambahkan angka ke dalam array
			numbers = append(numbers, num)
		}
	}

	fmt.Println("Program selesai.")
}
