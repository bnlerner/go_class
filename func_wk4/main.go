/*
create a set of animals to get information about those
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Animal type
type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{ name string }
type Snake struct{ name string }
type Bird struct{ name string }

func (c Cow) Eat()   { fmt.Println("Grass") }
func (c Cow) Move()  { fmt.Println("Walk") }
func (c Cow) Speak() { fmt.Println("Moo") }

func (s Snake) Eat()   { fmt.Println("Mice") }
func (s Snake) Move()  { fmt.Println("Slither") }
func (s Snake) Speak() { fmt.Println("Hsss") }

func (b Bird) Eat()   { fmt.Println("Worms") }
func (b Bird) Move()  { fmt.Println("Fly") }
func (b Bird) Speak() { fmt.Println("Peep") }

func SatisfyQuery(AllAnimals map[string]Animal, an, command string) {

	a, ok := AllAnimals[an]

	if !ok {
		fmt.Println("animal not found")
	}
	SatisfyCommand(a, command)

	/*
		for _, a := range AllAnimals {
			c, _ := a.(Cow)
			s, _ := a.(Snake)
			b, _ := a.(Bird)
			if c.name == an {
				SatisfyCommand(c, command)
			} else if s.name == an {
				SatisfyCommand(s, command)
			} else if b.name == an {
				SatisfyCommand(b, command)
			} else {
				continue
			}
		}
	*/
}

func SatisfyCommand(a Animal, command string) {
	switch strings.ToLower(command) {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	default:
		fmt.Println("Unknown command")
	}
}

func CreateAnimal(AllA *map[string]Animal, AName, AnType string) {

	var a Animal
	switch strings.ToLower(AnType) {
	case "cow":
		a = Cow{AName}
	case "snake":
		a = Snake{AName}
	case "bird":
		a = Bird{AName}
	default:
		fmt.Println("unknown animal type")
	}
	(*AllA)[AName] = a
	fmt.Println("Created it!")
}

func main() {
	var input string
	animals := make(map[string]Animal)

	fmt.Println("Please enter a command [newanimal | query] then arguments. Either the name and animal type for new animal or animal type and action for query")
	for {
		fmt.Print(">")
		in := bufio.NewScanner(os.Stdin)
		in.Scan()
		input = in.Text()
		commands := strings.Fields(input)

		if len(commands) != 3 {
			fmt.Println("issue with entry")
			continue
		}

		switch commands[0] {
		case "newanimal":
			CreateAnimal(&animals, commands[1], commands[2])
		case "query":
			SatisfyQuery(animals, commands[1], commands[2])
		default:
			fmt.Print("Unknown type")
		}
	}
}
