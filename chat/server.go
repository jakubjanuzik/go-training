package chat

import (
	"bufio"
	"log"
	"math/rand"
	"net"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Server struct {
	Address string
}

func (server Server) Start() {

	mainRoom := server.createRoom("main", server.InitDb())
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

func (server Server) createRoom(name string, db *gorm.DB) *room {
	newRoom := room{
		name:     name,
		outgoing: make(chan message),
		joining:  make(chan *client),
		shutdown: make(chan struct{}),
		clients:  make([]*client, 0),
		db:       db,
	}

	newRoom.listen()

	return &newRoom
}

func (server Server) InitDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", "chat.db")
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&chatMessage{})
	return db
}
