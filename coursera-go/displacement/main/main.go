package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, v, s float64
	readInLoop(a)
	readInLoop(v)
	readInLoop(s)

	fn := GenDisplaceFn(a, v, s)

	var t float64
	readInLoop(t)

	fmt.Printf("result: %f", fn(t))
}

func readInLoop(i float64) {
	stdin := bufio.NewReader(os.Stdin)
	repeat := true

	for repeat {
		println("enter a number:")
		_, err := fmt.Fscan(stdin, &i)

		if err != nil {
			fmt.Printf("Error happened: %s\n", err.Error())
			repeat = true
		} else {
			repeat = false
		}
	}
}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 1/2*a*t*t + v0*t + s0
	}
}
