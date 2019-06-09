package main

import "fmt"

func main() {
	var val float64

	fmt.Println("Enter a float to be truncated!")
	_, err := fmt.Scan(&val)
	if err == nil {
		fmt.Println(int32(val))
	} else {
		fmt.Println(err)
	}
}
