package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

/*
Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
 The program should be written as a loop. Before entering the loop, the program should create an empty
 integer slice of size (length) 3. During each pass through the loop, the program prompts the user to enter
  an integer to be added to the slice. The program adds the integer to the slice, sorts the slice, and
  prints the contents of the slice in sorted order. The slice must grow in size to accommodate any number
  of integers which the user decides to enter. The program should only quit (exiting the loop) when the
  user enters the character ‘X’ instead of an integer.
*/

func main() {

	var input string
	myArray := make([]int, 0, 3)
	fmt.Println("Hi welcome to the int function. Enter a number or X to exit")
	for {
		fmt.Println("enter number:")
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if input == "X" {
			fmt.Println(myArray)
			break
		}
		intInput, _ := strconv.Atoi(input)
		myArray = append(myArray, intInput)
		sort.Ints(myArray)
		fmt.Println(myArray)
	}

}
