package main

import (
	"fmt"

	"algo4/algo4"
	"algo4/stdin"
)

// go run cmd/selection/main.go < data/tiny.txt
// go run cmd/selection/main.go < data/words3.txt
func main() {
	items := stdin.ReadAllStrings()
	algo4.SelectionSort(algo4.StringSlice(items))
	if algo4.IsSorted(algo4.StringSlice(items)) {
		fmt.Println("sorted items:", items)
	} else {
		fmt.Println("not sorted")
		fmt.Println(items)
	}
}
