package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food, locomotion, noise string
}

func (a Animal) Eat() {
	println(a.food)
}

func (a Animal) Move() {
	println(a.locomotion)
}

func (a Animal) Speak() {
	println(a.noise)
}

func main() {
	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	stdin := bufio.NewReader(os.Stdin)

	for {
		println("Print request (animal property): ")
		r, err := stdin.ReadString('\n')

		if err != nil {
			fmt.Printf("Error happened: %s, repeating", err.Error())
			continue
		}

		req := strings.Split(strings.TrimSuffix(r, "\n"), " ")
		animal := req[0]

		var prop string
		if len(req) < 2 {
			prop = ""
		} else {
			prop = req[1]
		}

		switch animal {
		case "cow":
			printProp(cow, prop)
		case "bird":
			printProp(bird, prop)
		case "snake":
			printProp(snake, prop)
		default:
			fmt.Printf("Unknown animal '%s'\n", animal)
		}
	}
}

func printProp(animal Animal, prop string) {
	switch prop {
	case "food":
		animal.Eat()
	case "locomotion":
		animal.Move()
	case "noise":
		animal.Speak()
	default:
		fmt.Printf("Unknown animal property '%s'\n", prop)
	}
}
