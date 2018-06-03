package main

import "github.com/jakubjanuzik/chat"

func main() {
	// chat.Server{":6000", make(map[string]chat.Room)}.Start()
	chat.Server{":6000"}.Start()
	// chat.CreateRoom("go")
}
