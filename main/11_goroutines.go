package main

import (
	"log"
	"math/rand"
	"runtime"
	"sync"
)

var group1 sync.WaitGroup

func init() {
	println("Init...")
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func task(name string) {
	for count := 0; count < 1000; count++ {
		log.Printf("%v: %v\n", name, rand.Int())
	}
	group.Done()
}

func main11() {
	group.Add(2)
	go task("T1")
	go task("T2")
	group.Wait()
}
