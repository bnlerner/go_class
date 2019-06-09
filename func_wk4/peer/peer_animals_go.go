package main

/*

new_animals.go



###

The goal of this assignment is to write a GoLang routine that allows users to create a set of animals and then get information about those animals.
Review criteria

This assignment is worth 10 points.

Test the program by running it and issuing three newanimal commands followed by three query commands. Each new animal should involve a different animal type (cow, bird, snake), each with a different name. Each query should involve a different animal and a different property of the animal (eat, move, speak). Give 2 points for each query for which the program provides the correct response.

Examine the code. If the code contains an interface type called Animal, which is a struct containing three fields, all of which are strings, then give another 2 points. If the program contains three types â€“ Cow, Bird, and Snake â€“ which all satisfy the Animal interface, give another 2 points.

####


Write a program which allows the user to create a set of animals and to get information about those animals. Each animal has a name and can be either a cow, bird, or snake. With each command, the user can either create a new animal of one of the three types, or the user can request information about an animal that he/she has already created. Each animal has a unique name, defined by the user. Note that the user can define animals of a chosen type, but the types of animals are restricted to either cow, bird, or snake. The following table contains the three types of animals and their associated data.
Animal	Food eaten	Locomtion method	Spoken sound
cow	grass	walk	moo
bird	worms	fly	peep
snake	mice	slither	hsss

Your program should present the user with a prompt, ">â€, to indicate that the user can type a request. Your program should accept one command at a time from the user, print out a response, and print out a new prompt on a new line. Your program should continue in this loop forever. Every command from the user must be either a "newanimalâ€ command or a "queryâ€ command.

Each "newanimalâ€ command must be a single line containing three strings. The first string is "newanimalâ€. The second string is an arbitrary string which will be the name of the new animal. The third string is the type of the new animal, either "cowâ€, "birdâ€, or "snakeâ€. Your program should process each newanimal command by creating the new animal and printing "Created it!â€ on the screen.

Each "queryâ€ command must be a single line containing 3 strings. The first string is "queryâ€. The second string is the name of the animal. The third string is the name of the information requested about the animal, either "eatâ€, "moveâ€, or "speakâ€. Your program should process each query command by printing out the requested data.

Define an interface type called Animal which describes the methods of an animal. Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values. The Eat() method should print the animalâ€™s food, the Move() method should print the animalâ€™s locomotion, and the Speak() method should print the animalâ€™s spoken sound. Define three types Cow, Bird, and Snake. For each of these three types, define methods Eat(), Move(), and Speak() so that the types Cow, Bird, and Snake all satisfy the Animal interface. When the user creates an animal, create an object of the appropriate type. Your program should call the appropriate method when the user issues a query command.



*/

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*according instructions : Define an interface type called Animal which describes the methods of an animal. Specifically, the Animal interface should contain the methods Eat(), Move(), and Speak(), which take no arguments and return no values
 */

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{ name string }

func (c Cow) Eat() {
	fmt.Println("grass")
}
func (c Cow) Move() {
	fmt.Println("walk")
}
func (c Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct{ name string }

func (c Bird) Eat() {
	fmt.Println("worms")
}
func (c Bird) Move() {
	fmt.Println("fly")
}
func (c Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct{ name string }

func (c Snake) Eat() {
	fmt.Println("mice")
}
func (c Snake) Move() {
	fmt.Println("slither")
}
func (c Snake) Speak() {
	fmt.Println("hsss")
}

func print_animal_fully(a Animal) {
	fmt.Printf("Animal type %T\n", a)
	switch underling := a.(type) {
	case Cow:
		fmt.Println("name :", underling.name)
	case Snake:
		fmt.Println("name :", underling.name)
	case Bird:
		fmt.Println("name :", underling.name)
	default:
		fmt.Printf("Animal sub type unknown %T\n", a)
	}
	a.Speak()
	a.Move()
	a.Eat()
}

func readln_whitespace_str(prompt string) string {
	var ss string
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		ss = scanner.Text()
		//fmt.Printf("got: %q\n", ss)
	} else {
		fmt.Println("error reading input\n")
		os.Exit(1)
	}
	return ss
}

func help() {
	fmt.Println(`
	Every command from the user must be either a 'newanimal' command or a 'query' command  
	 Each 'newanimal' command must be a single line containing three strings. 
	The first string is 'newanimal'. The second string is an arbitrary string which will be the name of the new animal. 
	The third string is the type of the new animal, either 'cow', 'bird', or 'snake'. 
	Your program should process each newanimal command by creating the new animal and printing 'Created it!' on the screen.  
	 Each 'query' command must be a single line containing 3 strings. 
	The first string is 'query'. The second string is the name of the animal. 
	The third string is the name of the information requested about the animal, either 'eat', 'move', or 'speak'. 
	Your program should process each query command by printing out the requested data
	e.g
	newanimal Hellu cow
	newanimal Rock snake
	newanimal Tipu bird
	query Tipu eat
	query Rock move
	query Hellu speak
	`)
}

const debug = false

var animals map[string]Animal

func find_animal(aname string) Animal {
	//a := Bird{name: aname}
	a, ok := animals[aname]
	if !ok {
		fmt.Println("ERROR Could not find animal named (" + aname + ")")
		return a
	}
	if debug {
		print_animal_fully(a)
	}
	return a
}

func store_animal(a Animal, name string) {
	animals[name] = a
	if debug {
		print_animal_fully(a)
		for key, value := range animals {
			fmt.Println("Key:", key, "Value:", value)
		}
	}
}

func main() {

	var ss string = ""
	animals = make(map[string]Animal)

	help()

	for {
		ss = readln_whitespace_str(" > ")
		//fmt.Printf("ss: %s\n", ss)
		s := strings.Split(ss, " ")

		if len(s) < 3 {
			fmt.Printf("\nERROR input %s splitted len=%d cap=%d %v\n", ss, len(s), cap(s), s)
			help()
			continue
		}

		command, aname, req := s[0], s[1], s[2]
		if debug {
			fmt.Printf("command %s name %s req %s \n", command, aname, req)
		}

		switch command {
		case "newanimal":
			animalt := req
			var a Animal
			switch animalt {
			case "cow":
				a = Cow{name: aname}
			case "bird":
				a = Bird{name: aname}
			case "snake":
				a = Snake{name: aname}
			default:
				fmt.Printf("ERROR: unregonized animal (%s)\nPlease try again\n", animalt)
				continue
			}

			store_animal(a, aname)
			fmt.Println("Created it!")

		case "query":

			operstr := req
			a := find_animal(aname)
			if a == nil {
				fmt.Println("ERROR finding animal named (" + aname + ") please try again")
				continue
			}

			switch operstr {
			case "eat":
				a.Eat()
			case "move":
				a.Move()
			case "speak":
				a.Speak()
			default:
				fmt.Printf("ERROR: unregonized operation (%s)\nPlease try again\n", operstr)
				continue
			}
		}
	}
}
