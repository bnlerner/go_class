package main

import (
	"fmt"
	"sync"
	"time"
)

var eatGroup sync.WaitGroup

type chopStick struct{ sync.Mutex }

type philosopher struct {
	id                            int
	leftChopStick, rightChopStick *chopStick
}

func (p philosopher) eat() {
	for j := 0; j < 3; j++ {
		p.leftChopStick.Lock()
		p.rightChopStick.Lock()

		logAction("starting to eat", p.id)
		time.Sleep(time.Second)

		p.rightChopStick.Unlock()
		p.leftChopStick.Unlock()

		logAction("finishing eating", p.id)
		time.Sleep(time.Second)
	}
	eatGroup.Done()
}

func logAction(action string, id int) {
	fmt.Printf(" %s #%d\n", action, id+1)
}

func main() {
	// How many philosophers and chopStick
	count := 5

	// Create chopSticks
	chopSticks := make([]*chopStick, count)
	for i := 0; i < count; i++ {
		chopSticks[i] = new(chopStick)
	}

	// Create philosopher, assign them 2 chopSticks and send them to the dining table
	philosophers := make([]*philosopher, count)

	for i := 0; i < count; i++ {
		philosophers[i] = &philosopher{
			i,
			chopSticks[i],
			chopSticks[(i+1)%count],
		}

		eatGroup.Add(1)

		go philosophers[i].eat()
	}
	eatGroup.Wait()
}
