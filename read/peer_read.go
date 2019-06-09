package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	type Person struct {
		fname string
		lname string
	}
	var people []Person
	var filename string

	fmt.Printf("Please enter the name of the text file: ")
	fmt.Scan(&filename)

	dat, _ := ioutil.ReadFile(filename)

	lines := strings.Split(string(dat), "\n")

	for _, line := range lines {
		name := strings.Split(string(line), " ")
		oneperson := Person{name[0], name[1]}
		people = append(people, oneperson)
	}

	for _, oneperson := range people {
		fmt.Println(oneperson.fname, oneperson.lname)
	}
}
