package main

import (
	"fmt"
)

// Swap performs a check on adjacent numbers and swaps them
func Swap(slc []int, index int) {
	if slc[index] > slc[index+1] {
		first := slc[index]
		slc[index] = slc[index+1]
		slc[index+1] = first
	}
}

// BubbleSort Performs the bubble sort algo
func BubbleSort(slc []int) {
	for i := len(slc) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			Swap(slc, j)
		}
	}
}

func main() {

	fmt.Print("Enter 10 integers\n")

	var input int
	var slc []int
	for i := 0; i < 10; i++ {
		fmt.Printf("enter number %d:", i+1)
		fmt.Scan(&input)
		slc = append(slc, input)
	}

	BubbleSort(slc)

	fmt.Println("bubble sort result")
	fmt.Println(slc)
}
