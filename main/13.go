package main

func main13() {
	channel := make(chan int)
	go func() {
		println(<-channel)
	}()

	channel <- 1
}
