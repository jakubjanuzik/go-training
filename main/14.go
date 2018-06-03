package main

import (
	"math/rand"
	"time"
)

const capacity = 10

var storage = make(chan int, capacity)

func producer() {
	for counter := 0; counter < 10000; counter++ {
		storage <- rand.Int()
		time.Sleep(5 * time.Millisecond)
	}
	time.Sleep(5000)
	close(storage)
}

func main() {
	go producer()
	for value := range storage {
		println(value)
	}
}
