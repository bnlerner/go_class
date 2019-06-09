package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Write a program which reads information from a file and represents it in a slice of structs. Assume that there is a text file
which contains a series of names. Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name, and lname for the last name.
Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file. Your program will successively read each line of the text
file and create a struct which contains the first and last names found in the file. Each struct created will be added to a
slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in
the file. After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.

*/

type Name struct {
	fname string
	lname string
}

type lNames struct {
	Items []Name
}

func (Items *lNames) addName(item Name) []Name {
	Items.Items = append(Items.Items, item)
	return Items.Items
}

func main() {

	fmt.Printf("please enter file name to read:")
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	fileName := in.Text()

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("error reading file")
		os.Exit(1)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	oneName := []Name{}
	listNames := lNames{oneName}

	var line string
	var iterName Name

	for {
		line, err = reader.ReadString('\n')
		words := strings.Fields(line)
		iterName = Name{
			fname: words[0],
			lname: words[1],
		}
		listNames.addName(iterName)
		if err != nil {
			break
		}
	}

	//starting loop
	for _, fullName := range listNames.Items {
		fmt.Print("first name: ", fullName.fname, "    ")
		fmt.Print("last name: ", fullName.lname, "\n")
	}
}
