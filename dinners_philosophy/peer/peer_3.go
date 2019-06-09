package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var c chan int
var d chan int

type ChopS struct {
	sync.Mutex
}

type Philo struct {
	leftCS, rightCS *ChopS
	number          int
}

func (p Philo) eat() {
	var k int
	i := 0
	for {
		k = <-c
		if k == p.number {
			fmt.Println("starting to eat", p.number)
			fmt.Println("finishing eating", p.number)
			i++
			d <- p.number
			if i == 3 {
				break
			}
		} else {
			c <- k
		}
	}
	wg.Done()
}

func host() {
	for i := 0; i < 3; i++ {
		// philos 1 and 3 eating concurrently
		c <- 1
		c <- 3
		<-d
		<-d
		// philos 5 and 2 eating concurrently
		c <- 5
		c <- 2
		<-d
		<-d
		c <- 4
		<-d
	}
	wg.Done()
}

func main() {
	c = make(chan int)
	d = make(chan int)
	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}
	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5], i + 1}
	}
	wg.Add(6)
	go host()
	for i := 0; i < 5; i++ {
		go philos[i].eat()
	}
	wg.Wait()
}
