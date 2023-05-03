package listener

import (
	"fmt"
	"log"
	"net"
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

		}
	}(conn)

	fmt.Println("New client connected:", conn.RemoteAddr())

	buffer := make([]byte, 256)
	for {
		data, err := conn.Read(buffer)
		println("dsa\n", data)
		if err != nil {
			fmt.Println("Error reading from client: ", err.Error())
			break
		}
	}
}