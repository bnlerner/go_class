package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	fmt.Print("Enter integers in one line seperated by spaces. Number of integers should be > 4 for smooth partitioning  \n")
	sl := make([]int, 0, 10)
	reader := bufio.NewReader(os.Stdin)
	raw_string_input, _ := reader.ReadString('\n')
	raw_val := strings.Split(strings.TrimSpace(strings.TrimSuffix(raw_string_input, "\n")), " ")

	if len(raw_val) < 4 {
		fmt.Println("sequence length should be >= 4")
		return
	}

	for _, v := range raw_val {
		int_val, _ := strconv.ParseInt(v, 10, 32)
		sl = append(sl, int(int_val))
	}

	// partition capacity
	pc := make([]int, 4, 4)
	extra := len(raw_val) % 4
	initial_cap := len(raw_val) / 4
	for indx, _ := range pc {
		pc[indx] = initial_cap
		if extra > 0 {
			pc[indx] += 1
			extra -= 1
		}
	}
	p1 := sl[0:pc[0]]

	p2_end := pc[0] + pc[1]
	p2 := sl[pc[0]:p2_end]

	p3_end := p2_end + pc[2]
	p3 := sl[p2_end:p3_end]

	p4_end := p3_end + pc[3]
	p4 := sl[p3_end:p4_end]

	fmt.Println("\nPartition Details: ")
	fmt.Println("partition 1", "length:", len(p1), p1)
	fmt.Println("partition 2", "length:", len(p2), p2)
	fmt.Println("partition 3", "length:", len(p3), p3)
	fmt.Println("partition 4", "length:", len(p4), p4)

	var wg sync.WaitGroup
	wg.Add(4)
	go SortParition(p1, &wg)
	go SortParition(p2, &wg)
	go SortParition(p3, &wg)
	go SortParition(p4, &wg)

	wg.Wait()

	// doing merge in pair of 2, since merging 4 is more complex implementation wise
	res := Merge(p1, p2)
	res = Merge(res, p3)
	res = Merge(res, p4)

	fmt.Println("\nFinal Result:")
	fmt.Println(res)
}

func SortParition(sl []int, wg *sync.WaitGroup) {
	defer wg.Done()
	sort.Ints(sl)
}

func Merge(left, right []int) []int {
	size, i, j := len(left)+len(right), 0, 0
	slice := make([]int, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i] < right[j] {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}
