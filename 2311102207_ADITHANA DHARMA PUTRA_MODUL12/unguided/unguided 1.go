package main
import "fmt"

func selectionSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[maxIdx] {
				maxIdx = j
			}
		}
		arr[i], arr[maxIdx] = arr[maxIdx], arr[i]
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	if n <= 0 || n >= 1000 {
		fmt.Println("n harus lebih besar dari 0 dan kurang dari 1000.")
		return
	}

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)

		if m <= 0 || m >= 1000000 {
			fmt.Println("harus lebih besar dari 0 dan kurang dari 1000000.")
			return
		}

		houses207 := make([]int, m)
		fmt.Printf("hasil kerabat dekat ke-%d  ", i+1)
		for j := 0; j < m; j++ {
			fmt.Scan(&houses207[j])
		}

		selectionSort(houses207)
//chill guy
		for j := 0; j < m; j++ {
			if houses207[j] % 2 != 0 {
				fmt.Printf("%d ", houses207[j])
			}
		}
		for j := m-1; j >= 0; j-- {
			if houses207[j] % 2 == 0 {
				fmt.Printf("%d ", houses207[j])
			}
		}

		fmt.Println()
	}
}
