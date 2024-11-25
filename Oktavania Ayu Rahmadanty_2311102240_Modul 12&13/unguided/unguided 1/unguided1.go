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

	fmt.Println("Masukkan jumlah daerah:")
	scanner.Scan()
	numRegions, _ := strconv.Atoi(scanner.Text())

	results := make([][]int, numRegions)

	for i := 0; i < numRegions; i++ {
		fmt.Printf("Masukkan bilangan untuk daerah %d:\n", i+1)
		scanner.Scan()
		input := scanner.Text()

		// Parsing input
		nums := parseInput(input)

		// Pisahkan ganjil dan genap
		oddNums, evenNums := separateOddEven(nums)

		// Urutkan sesuai ketentuan
		sort.Sort(sort.Reverse(sort.IntSlice(oddNums))) // Ganjil: descending
		sort.Ints(evenNums)                             // Genap: ascending

		// Gabungkan hasil
		results[i] = append(oddNums, evenNums...)
	}

	// Cetak hasil
	fmt.Println("\nKeluaran:")
	for i, result := range results {
		fmt.Printf("Daerah %d: %v\n", i+1, result)
	}
}

func parseInput(input string) []int {
	parts := strings.Fields(input)
	nums := make([]int, len(parts))
	for i, part := range parts {
		nums[i], _ = strconv.Atoi(part)
	}
	return nums
}

func separateOddEven(nums []int) ([]int, []int) {
	var oddNums, evenNums []int
	for _, num := range nums {
		if num%2 == 0 {
			evenNums = append(evenNums, num)
		} else {
			oddNums = append(oddNums, num)
		}
	}
	return oddNums, evenNums
}
