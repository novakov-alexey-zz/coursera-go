package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	sli := make([]Name, 0)
	filePath := "coursera-go/read/main/file.txt"

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Printf("Read line: %s\n", text)
		sli = append(sli, readName(text))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, n := range sli {
		fmt.Printf("First Name: %s, Last Name: %s \n", n.fname, n.lname)
	}
}

func readName(text string) Name {
	split := strings.Split(text, " ")
	return Name{fname: split[0], lname: split[1]}
}
