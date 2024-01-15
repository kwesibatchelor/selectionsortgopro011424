package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	piece := createPiece(20)
	fmt.Printf("\nUnsorted:\n%v\n", piece)
	selectionSort(piece)
	fmt.Printf("\nSorted:\n%v\n", piece)
}

func createPiece(size int) []int {
	piece := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		piece[i] = rand.Intn(1999) - 999
	}
	return piece
}

func selectionSort(items []int) {
	var n = len(items)
	for i := 0; i < n; i++ {
		var min = i
		for j := i; j < n; j++ {
			if items[j] < items[min] {
				min = j
			}
		}
		items[i], items[min] = items[min], items[i]
	}
}
