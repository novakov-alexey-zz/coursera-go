package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	stdin := bufio.NewReader(os.Stdin)
	m := make(map[string]string)

	for {
		println("Print Name: ")
		name, err := stdin.ReadString('\n')
		if err != nil {
			println("Error happened, repeating")
			continue
		}
		m["name"] = strings.Trim(name, "\n")

		println("Print Address: ")
		address, err := stdin.ReadString('\n')
		if err != nil {
			println("Error happened, repeating")
			continue
		}
		m["address"] = strings.Trim(address, "\n")
		break
	}

	js, err := json.Marshal(m)
	if err != nil {
		fmt.Printf("Failed to encode to json: %s\n", err.Error())
	} else {
		println(string(js[:]))
	}
}
