package main

import (
	"fmt"
	"math/rand"
	"sync"
)

/*
Implement the dining philosopher’s problem with the following constraints/modifications.

There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
The host allows no more than 2 philosophers to eat concurrently.
Each philosopher is numbered, 1 through 5.
When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.
*/

var wg sync.WaitGroup

type ChopS struct{ sync.Mutex }

type Philosopher struct {
	leftC, rightC *ChopS
	num           int
}

func (p Philosopher) eat(EatTimes int, ch chan bool, outchan chan int) {
	var ranNum int

	for i := 0; i < EatTimes; i++ {

		ranNum = rand.Intn(100)
		if ranNum%2 == 0 {
			p.leftC.Lock()
			p.rightC.Lock()
		} else {
			p.rightC.Lock()
			p.leftC.Lock()
		}
		//fmt.Printf("philosopher %d obtained lock and waiting for go ahead from master thread\n", p.num)
		<-ch
		//fmt.Printf("philosopher %d obtained go ahead\n", p.num)

		fmt.Printf("starting to eat %d\n", p.num)
		outchan <- 1
		fmt.Printf("finishing eating %d\n", p.num)

		p.leftC.Unlock()
		p.rightC.Unlock()

	}
	//fmt.Printf("philosopher %d finished eating\n", p.num)
	wg.Done()
}

func MasterThread(MaxEating int, DoneEating chan int, OkToEat chan bool) {

	for i := 0; i < MaxEating; i++ {
		//fmt.Printf("creating eater number %d\n", i+1)
		OkToEat <- true
		//fmt.Println("sent signal to eater")
	}
	for {
		select {
		case <-DoneEating:
			//fmt.Println("received signal from eater. sending new")
			OkToEat <- true
		}
	}
}

func main() {
	CSticks := make([]*ChopS, 5)
	philos := make([]*Philosopher, 5)

	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	var InChan chan int
	var OutChan chan bool
	EatNum := 2

	OutChan = make(chan bool, EatNum)
	InChan = make(chan int)

	for i := 0; i < 5; i++ {
		//fmt.Printf("creating philosopher %d\n", i+1)
		wg.Add(1)
		philos[i] = &Philosopher{CSticks[i], CSticks[(i+1)%5], i + 1}
		go philos[i].eat(3, OutChan, InChan)
	}
	go MasterThread(EatNum, InChan, OutChan)
	wg.Wait()
}
