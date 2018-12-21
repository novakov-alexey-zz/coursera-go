package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := readArray()
	fmt.Printf("Input: %v\n", input)

	BubbleSort(input)
	fmt.Printf("Sorted: %v", input)
}

func readArray() []int {
	stdin := bufio.NewReader(os.Stdin)
	n := 0
	input := make([]int, 0)
	max := 10

	for n < max {
		fmt.Printf("Input integer (%d/%d):\n", n+1, max)
		var i int
		_, err := fmt.Fscan(stdin, &i)

		if err != nil {
			fmt.Printf("Error happened: %s, repeating", err.Error())
			continue
		}
		input = append(input, i)
		n++
	}
	return input
}

func BubbleSort(input []int) {
	swapped := true
	j := 0

	for swapped {
		swapped = false
		j++
		for i := 0; i < len(input)-j; i++ {

			if input[i] > input[i+1] {
				Swap(input, i)
				swapped = true
			}
		}
	}
}

func Swap(input []int, i int) {
	a := input[i]
	input[i] = input[i+1]
	input[i+1] = a
}
