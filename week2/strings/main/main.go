package main

import (
	"bufio"
	"fmt"
	"os"
	. "strings"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)

	println("Enter string with 'i', 'a', 'n' : ")
	s, err := stdin.ReadString('\n')

	if err != nil {
		fmt.Printf("Error happened: %s, repeating.", err.Error())
		main()
	} else {
		ss := ToLower(TrimSpace(s))
		i := HasPrefix(ss, "i")
		n := HasSuffix(ss, "n")
		a := len(ss) > 2 && Contains(ss[1:len(ss)-1], "a")

		if i && a && n {
			println("Found!")
		} else {
			println("Not Found!")
		}
	}
}
