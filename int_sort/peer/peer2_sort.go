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
	var input string
	s := bufio.NewScanner(os.Stdin)
	fmt.Print(">")
	if s.Scan() {
		input = s.Text()
		chars := strings.Fields(input)
		vals := []int{}
		for _, c := range chars {
			v, err := strconv.Atoi(c)
			if err != nil {
				panic(err)
			}
			vals = append(vals, v)
		}

		c := make(chan []int)
		var wg sync.WaitGroup
		q := len(vals) / 4
		for i := 0; i < len(vals)-1; i += q {
			chunk := vals[i:min(i+q, len(vals))]
			wg.Add(1)
			go worker(&wg, c, chunk)
		}

		go func() {
			wg.Wait()
			close(c)
		}()

		res := []int{}
		for i := range c {
			res = append(res, i...)
		}

		sort.Ints(res)
		fmt.Print("\n", res, "\ndone\n")
	}
}

func worker(wg *sync.WaitGroup, c chan []int, i []int) {
	defer wg.Done()
	fmt.Println("sorting >", i)
	sort.Ints(i)
	c <- i
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
