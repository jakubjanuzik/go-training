package chat

import (
	"net"

	"github.com/jinzhu/gorm"
)

type room struct {
	name     string
	outgoing chan message
	joining  chan *client
	shutdown chan struct{}
	clients  []*client
	listener net.Listener
	db       *gorm.DB
}

func (room *room) broadcast(message message) {
	for _, client := range room.clients {
		if client.connection != message.sender.connection {
			client.write(message)
		}
	}
}

func (room *room) addClient(client *client) {
	room.clients = append(room.clients, client)
	client.room = room
}

func (room *room) listen() {
	go func() {
		for {
			select {
			case message := <-room.outgoing:
				room.broadcast(message)
				room.db.Create(&chatMessage{Text: message.text, Sender: message.sender.name})
			}
		}
	}()
}
