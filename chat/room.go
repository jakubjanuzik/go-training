package chat

import (
	"net"
)

type room struct {
	name     string
	outgoing chan message
	joining  chan *client
	shutdown chan struct{}
	clients  []*client
	listener net.Listener
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
				// if strings.HasPrefix(message.text, "#") {
				// 	roomName := strings.Split(message.text, "#")
				// 	room := room.server.createRoom(strings.Join(roomName, ""))
				// 	room.addClient(message.sender)
				// } else {
				// 	room.broadcast(message)
				// }
				room.broadcast(message)
			}
		}
	}()
}
