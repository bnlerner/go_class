package main

import (
	"fmt"
	"sync"
)

var c = make(chan int, 2)

var mu sync.RWMutex
var everyoneAte int
var timesEaten = make(map[int]int, 5)

type chopstick struct {
	sync.Mutex
}

type philosopher struct {
	leftCs  *chopstick
	rightCs *chopstick
}

func srv(c chan int, count int) {

	for {
		select {
		case msg := <-c:
			c <- msg
		default:
			fmt.Println("no message received")
		}

		if everyoneAte == count {
			break
		}
	}
}

func (p philosopher) eat(index int) {

	c <- 1
	// ask for permition from server

	if timesEaten[index] < 3 {
		p.leftCs.Lock()
		p.rightCs.Lock()

		fmt.Printf("Starting to eat %v\n", index)
		fmt.Printf("Finishing eating %v\n", index)
		mu.Lock()
		timesEaten[index]++

		if timesEaten[index] == 3 {
			everyoneAte++
		}
		mu.Unlock()
		p.rightCs.Unlock()
		p.leftCs.Unlock()

	}
	<-c
}

func main() {
	count := 5
	chopsticks := make([]*chopstick, count)
	// start SERVER(limit 2)
	go srv(c, count)
	for i := 0; i < count; i++ {
		chopsticks[i] = &chopstick{}
	}

	philosophers := make([]*philosopher, count)
	for i := 0; i < count; i++ {
		philosophers[i] = &philosopher{
			leftCs:  chopsticks[i],
			rightCs: chopsticks[(i+1)%count],
		}
	}
	for {
		mu.RLock()
		if everyoneAte == count {
			//fmt.Println("time eaten === ", timesEaten)
			return
		}
		for i := 0; i < count; i++ {
			if timesEaten[i] == 3 {
				continue
			}
			go philosophers[i].eat(i + 1)
		}
		mu.RUnlock()
	}

}
