package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// our host hands out the forks, the Mutex is used
// to make sure only 1 philosopher can take a fork
// at any given time.
type host struct {
	sync.Mutex
	p     int
	forks []bool
	wg    sync.WaitGroup
}

// NewHost initializes a new host
func newHost(f int) *host {
	// make new host and set # of forks and philosophers
	h := host{
		forks: make([]bool, f),
		p:     f,
	}
	// create a waitgroup
	h.wg.Add(f)
	// set all forks to true
	// true -> host has the fork
	// false -> philosopher has the fork
	for i := 0; i < f; i++ {
		h.forks[i] = true
	}
	return &h
}

// philosopher tries to take both forks and eat.
// give each philosopher a place at the table (i),
// and times to eat (t)
func (h *host) philosopher(i int, t int) {
	// philosopher's forks
	var leftFork, rightFork bool
	// number of meals eaten
	n := 0
	// index of the fork to the left of the philosopher
	fi := i
	if i == 0 {
		fi = h.p
	}
	for n < t {
		// Randomly try to pick either left or right fork
		rand.Seed(time.Now().UnixNano())
		r := rand.Float64()
		if r > 0.5 {
			// left fork
			h.Lock()
			if h.forks[fi-1] {
				h.forks[fi-1] = false
				leftFork = true
			}
			h.Unlock()
		} else {
			// right fork
			h.Lock()
			if h.forks[i] {
				h.forks[i] = false
				rightFork = true
			}
			h.Unlock()
		}
		// If we have both forks, eat a meal
		if leftFork && rightFork {
			fmt.Printf("starting to eat %d (meal #%d)\n", i, n+1)
			n++
			// Take a little time to enjoy the meal..
			rand.Seed(time.Now().UnixNano())
			time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
			fmt.Printf("finishing eating %d\n", i)
			// After eating return both forks
			h.Lock()
			h.forks[fi-1] = true
			h.forks[i] = true
			h.Unlock()
		}
	}
	// Notify the waitgroup that this routine is done
	h.wg.Done()
}

func main() {
	// philosophers (note that we need >3 philosophers)
	p := 5
	// meals
	m := 3
	// tell the host how many philosophers are at the table
	host := newHost(p)
	// make the philosophers try to eat
	for i := 0; i < p; i++ {
		go host.philosopher(i, m)
	}
	// wait until all philosophers are done eating
	host.wg.Wait()
}
