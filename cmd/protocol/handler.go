package protocol

import (
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"seems.cloud/badwolf/server/internal/configs"
	"strconv"
	"time"
)

func ServerHandler() {
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Failed LoadConfig: %v", err)
	}

	certConfig, err := tlsCertConfig()
	if err != nil {
		log.Fatalf("Failed to load certificate %v\n", err)
	}

	listener, err := tls.Listen("tcp", ":"+strconv.Itoa(config.Port), certConfig)
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
		go handleServerConnection(conn)
	}
}

func ClientHandler() {
	conn := handleClientConn()
	for {
		messageBytes := protoDataBuilder([]byte("asdf"))
		protoDataSend(&conn, messageBytes)

		time.Sleep(time.Second)
	}
}
