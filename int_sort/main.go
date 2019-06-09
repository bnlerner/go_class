/*
Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is sorted
by a different goroutine. Each partition should be of approximately equal size. Then the main goroutine should merge the
4 sorted subarrays into one large sorted array.

The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should print
the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.
*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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
func BubbleSort(slc []int, c chan []int) {
	fmt.Printf("goroutine sorting %#v\n", slc)
	for i := len(slc) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			Swap(slc, j)
		}
	}
	c <- slc
}

// ConvertToIntArray blah
func ConvertToIntArray(numberSlice []string) []int {

	slc := make([]int, 0)
	for _, i := range numberSlice {
		a, _ := strconv.Atoi(i)
		slc = append(slc, a)
	}
	return slc
}

func LaunchWorkers(slc []int, workers float64) []int {

	var lenSlice = float64(len(slc))
	var firstNum int64
	var secondNum int64
	var finalSlice []int
	c := make(chan []int, int(workers))

	for i := 0; i < int(workers); i++ {
		firstNum = int64(math.Floor(lenSlice * float64(i) / workers))
		secondNum = int64(math.Floor(lenSlice * (float64(i) + 1.0) / workers))
		newSlice := slc[firstNum:secondNum]
		go BubbleSort(newSlice, c)
		sortedSlice := <-c
		finalSlice = append(finalSlice, sortedSlice...)
	}
	go BubbleSort(finalSlice, c)
	finalSlice = <-c
	return finalSlice
}

func main() {

	fmt.Print("Enter array of integers to sort\n example 2 3 4 52 3 43 5 3 \n")

	var input string
	var numWorkers = 4.0

	in := bufio.NewScanner(os.Stdin)
	in.Scan()

	input = in.Text()
	numbers := strings.Fields(input)
	slc := ConvertToIntArray(numbers)
	slc = LaunchWorkers(slc, numWorkers)

	fmt.Println("final slice:", slc)
}
