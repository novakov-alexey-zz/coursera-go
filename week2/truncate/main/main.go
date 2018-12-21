package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	println("Print float number: ")
	var f float32
	_, err := fmt.Fscan(stdin, &f)

	if err != nil {
		fmt.Printf("Error happened: %s, repeating", err.Error())
		_, _ = stdin.ReadString('\n')
		main()
	} else {
		i := int32(f)
		fmt.Printf("Int form: %d", i)
	}
}
