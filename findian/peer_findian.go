package main 

import (
	"fmt"
	"strings"
	"bufio"
	"os"
	)

func main() {

	var x,user_input string
	var starts_with_a, ends_with_n, contains_a bool

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter a string ")
	user_input, _ = reader.ReadString('\n')
	x = strings.Replace(user_input, "\n", "",-1)

	starts_with_a = strings.HasPrefix(strings.ToLower(x), "i")
	contains_a = strings.Contains(strings.ToLower(x), "a")
	ends_with_n = strings.HasSuffix(strings.ToLower(x), "n")

	if (starts_with_a && contains_a && ends_with_n) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
