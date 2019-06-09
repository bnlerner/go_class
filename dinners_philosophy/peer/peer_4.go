package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type host struct {
	active  map[int]*philo
	request chan int
}

type philo struct {
	id             int
	stickL, stickR *sync.Mutex
	notify         chan bool
}

func (h *host) listen(wg *sync.WaitGroup) {
	done := false
	for {
		select {
		case v := <-h.request:
			n := h.next(v)
			if n > -1 {
				h.active[n].notify <- true
			} else {
				if !done {
					done = true
					wg.Done()
				}
				return
			}
		default:
			if len(h.active) == 0 && !done {
				done = true
				wg.Done()
				return
			}
		}
	}
}

func (h *host) next(current int) int {
	for i := 1; i < 6; i++ {
		n := (current + i) % len(h.active)
		if _, ok := h.active[n]; ok {
			//fmt.Println("next id: ", n)
			return n
		}
	}
	return -10
}

func (h *host) finished(id int) {
	delete(h.active, id)
}

func (p *philo) eat() {
	var i int
	for {
		select {
		case <-p.notify:
			p.stickL.Lock()
			p.stickR.Lock()
			i++
			fmt.Printf("%v.\tstarting to eat: \tid:%v\t(%v of 3)\n", atomic.AddInt32(&counter, 1), p.id, i)
			fmt.Printf("%v.\tfinishing eating: \tid:%v\t(%v of 3)\n", atomic.AddInt32(&counter, 1), p.id, i)
			p.stickR.Unlock()
			p.stickL.Unlock()
			if i < 3 {
				hostr.request <- p.id
			} else {
				hostr.finished(p.id)
				hostr.request <- p.id
				return
			}
		default:
		}
	}
}

var hostr = host{}
var philos = [5]philo{}
var counter int32

func init() {
	hostr = host{active: make(map[int]*philo), request: make(chan int, 2)}
	sticks := [...]*sync.Mutex{&sync.Mutex{}, &sync.Mutex{}, &sync.Mutex{}, &sync.Mutex{}, &sync.Mutex{}}
	for i := 0; i < len(philos); i++ {
		philos[i].id = i
		philos[i].stickL = sticks[i]
		philos[i].stickR = sticks[(i+1)%len(philos)]
		philos[i].notify = make(chan bool)
		hostr.active[i] = &philos[i]
	}
	//fmt.Println(sticks)
	//fmt.Println(philos)
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go hostr.listen(&wg)

	for i := 0; i < len(philos); i++ {
		go philos[i].eat()
	}
	hostr.request <- 0
	hostr.request <- 3

	wg.Wait()
	fmt.Println("DONE")
}
