package main
import "fmt"

func selectionSort(arr207 []int) {
	n := len(arr207)
	for i := 0; i < n-1; i++ {
		maxIdx := i
		for j := i + 1; j < n; j++ {
			if arr207[j] < arr207[maxIdx] {
				maxIdx = j
			}
		}
		arr207[i], arr207[maxIdx] = arr207[maxIdx], arr207[i]
	}
}
func main() {
	var arr207 []int
	var slice []int
	var n int

	var median int
	for {
		fmt.Scan(&n)
		arr207 = append(arr207, n)
		if n == -5313 {
			break
		}
	}
	panjang := len(arr207)

	for i := 0; i < panjang; i++ {
		if arr207[i] == 0 {
			selectionSort(slice)
			if len(slice)%2 == 0 {

				median = (slice[(len(slice)/2)-1] + slice[len(slice)/2]) / 2

			} else {
				median = slice[len(slice)/2]
				fmt.Print()
			}
			fmt.Println(median)
		} else {
			slice = append(slice, arr207[i])
		}
	}

}