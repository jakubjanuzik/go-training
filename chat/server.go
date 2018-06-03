package chat

import (
	"bufio"
	"log"
	"math/rand"
	"net"
	"strconv"
)

type Server struct {
	Address string
}

func (server Server) Start() {
	mainRoom := server.createRoom("main")
	listener, err := net.Listen("tcp", server.Address)
	if err != nil {
		log.Fatal(err)
	}

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		} else {
			mainRoom.addClient(server.createClient(mainRoom, connection, strconv.Itoa(rand.Int())))
		}
	}
}

func (server Server) createClient(room *room, connection net.Conn, userName string) *client {
	newClient := client{
		name:       userName,
		room:       room,
		reader:     bufio.NewReader(connection),
		writer:     bufio.NewWriter(connection),
		connection: connection,
	}
	go newClient.read()
	return &newClient
}

func (server Server) createRoom(name string) *room {
	newRoom := room{
		name:     name,
		outgoing: make(chan message),
		joining:  make(chan *client),
		shutdown: make(chan struct{}),
		clients:  make([]*client, 0),
	}

	newRoom.listen()

	return &newRoom
}
