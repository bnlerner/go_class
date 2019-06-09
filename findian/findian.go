package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func found() {
	fmt.Println("Found!")
}

func notFound() {
	fmt.Println("Not Found!")
}

func main() {
	fmt.Println("Enter a string to check if it contains a i or n!")
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	input := in.Text()

	if (strings.HasPrefix(input, "i") || strings.HasPrefix(input, "I")) && (strings.HasSuffix(input, "n") || strings.HasSuffix(input, "N")) && (strings.Contains(input, "a") || strings.Contains(input, "A")) {
		found()
	} else {
		notFound()
	}
}
