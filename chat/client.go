package chat

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

type client struct {
	name       string
	room       *room
	reader     *bufio.Reader
	writer     *bufio.Writer
	connection net.Conn
}

func (client *client) read() {
	for {
		line, err := client.reader.ReadString('\n')

		if err != nil {
			if err == io.EOF {
				client.room.broadcast(message{fmt.Sprintf("ROOM: %v, Client Disconnected: ", client.room.name), client})
				log.Printf("Client disconnected")
				return
			} else {
				log.Printf("Exception: %v", err)
			}
		} else {
			client.room.outgoing <- message{line, client}
		}

	}
}

func (client *client) write(message message) {
	messageText := fmt.Sprintf("%s: %s", message.sender.name, message.text)
	client.writer.WriteString(messageText)
	client.writer.Flush()
}

func (client *client) disconnect() {
	client.connection.Close()
	fmt.Printf("Disconnecting %v", client.name)
}
