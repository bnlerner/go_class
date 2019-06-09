package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Queue struct {
	Message string
}

func main() {
	// Initializing queue
	myTestQueue := new(Queue)
	// "Sending" messages to queue
	go writeAfterDelay("I am 1st message", 1000, myTestQueue)
	go writeAfterDelay("I am 2nd message", 1000, myTestQueue)
	// Getting message from queue
	time.Sleep(time.Duration(1000) * time.Millisecond)
	fmt.Println("Last message in a queue was ", myTestQueue.Message)
}

func writeAfterDelay(message string, maxDelay int, queue *Queue) {
	randSouce := rand.NewSource(time.Now().UnixNano())
	randomizer := rand.New(randSouce)
	time.Sleep(time.Duration(randomizer.Intn(maxDelay)) * time.Microsecond)
	queue.Message = message
}