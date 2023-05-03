package listener

import (
	"fmt"
	"io"
	"log"
	"net"
	"seems.cloud/badwolf/server/cmd/protocol"
	"seems.cloud/badwolf/server/internal/configs"
	"strconv"
)

func Listen() {
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Failed LoadConfig: %v", err)
	}

	listener, err := net.Listen("tcp", ":"+strconv.Itoa(config.Port))
	if err != nil {
		log.Fatalf("Error listening: %v", err.Error())
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {

		}
	}(listener)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err.Error())
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("Error closing connection: %s\n", err.Error())
		}
	}(conn)

	fmt.Println("New client connected:", conn.RemoteAddr())

	buffer := make([]byte, 256)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("Client disconnected")
				return
			} else {
				fmt.Printf("Error when reading from client: %s\n", err)
			}
		}

		if n > 0 {
			messageHandler(buffer[:3])
		}
	}
}

func messageHandler(bytes []byte) {
	switch string(bytes) {
	case protocol.PingPongType:
		fmt.Printf("Received Ping Message\n")
	default:
		fmt.Printf("Unknown message\n")
	}
}
