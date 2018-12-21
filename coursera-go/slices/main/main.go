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
	stdin := bufio.NewReader(os.Stdin)
	sli := make([]int, 3)

	for {
		println("Print int number: ")
		s, err := stdin.ReadString('\n')

		if err != nil {
			fmt.Printf("Error happened: %s, repeating", err.Error())
			continue
		}

		ss := strings.TrimSuffix(s, "\n")
		if ss == "X" {
			break
		} else {
			i, err := strconv.Atoi(ss)
			if err != nil {
				fmt.Printf("%s is not int\n", err.Error())
				continue
			}

			sli = append(sli, i)

			sort.Slice(sli, func(i, j int) bool {
				return sli[i] < sli[j]
			})

			fmt.Println(sli)
		}
	}
}
