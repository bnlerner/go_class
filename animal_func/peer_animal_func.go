package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noice      string
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noice)
}

func main() {
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	reader := bufio.NewReader(os.Stdin)
	for true {
		fmt.Println("Please enter in following format: <animal-type> <animal-info>. Example: bird move")
		fmt.Print("> ")
		raw_string_input, _ := reader.ReadString('\n')
		raw_val := strings.TrimSuffix(raw_string_input, "\n")

		args := strings.Split(raw_val, " ")

		switch args[0] {
		case "cow":
			printInfo(cow, args[1])
		case "bird":
			printInfo(bird, args[1])
		case "snake":
			printInfo(snake, args[1])
		default:
			fmt.Println("Invalid animal reference")
		}
	}
}

func printInfo(a Animal, info string) {
	switch info {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	default:
		fmt.Println("Invalid Info Requested")

	}
}
