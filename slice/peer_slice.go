package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	intSlice := make([]int, 0, 3)

	for {
		var temp string
		fmt.Print("Please enter an inter or exit by entering 'X':")
		_, _ = fmt.Scan(&temp)
		if string(temp) == "X" {
			break
		} else {
			number, _ := strconv.Atoi(temp)
			//fmt.Print(temp,reflect.TypeOf(temp))
			intSlice = append(intSlice, number)
			sort.Ints(intSlice)
			fmt.Print(intSlice)
		}

	}
}
