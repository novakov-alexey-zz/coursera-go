package main

import (
	"fmt"
	"sort"
)

func main() {
	a := readArray()

	batchSize := 4
	batches := partition(a, batchSize)
	batchesCount := len(batches)
	fmt.Printf("batches: %v\n", batches)

	var ch = make(chan []int, batchSize)
	mapSorting(batchesCount, batches, ch)
	reducedArray := reduceSorting(batchesCount, ch)

	sortArray(reducedArray, ch)
	fullySorted := <-ch

	fmt.Printf("Result: %v", fullySorted)
}

func reduceSorting(batchesCount int, ch chan []int) []int {
	reducedArray := make([]int, 0)
	for i := 0; i < batchesCount; i++ {
		var sorted = <-ch
		reducedArray = append(reducedArray, sorted...)
	}
	return reducedArray
}

func mapSorting(batchesCount int, batches [][]int, ch chan []int) {
	for i := 0; i < batchesCount; i++ {
		go sortArray(batches[i], ch)
	}
}

func sortArray(a []int, c chan []int) {
	sort.Ints(a)
	c <- a
}

func readArray() []int {
	var n, i int
	fmt.Println("Enter the number of inputs")
	_, _ = fmt.Scanln(&n)

	fmt.Println("Enter the inputs")
	a := make([]int, n)
	for i = 0; i < n; i++ {
		_, _ = fmt.Scanln(&a[i])
	}

	return a
}

func partition(a []int, batchSize int) [][]int {
	var batches [][]int

	for batchSize < len(a) {
		a, batches = a[batchSize:], append(batches, a[0:batchSize:batchSize])
	}
	return append(batches, a)
}
