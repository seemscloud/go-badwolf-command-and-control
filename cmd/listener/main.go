package listener

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"path/filepath"
	"seems.cloud/badwolf/server/cmd/protocol"
	"seems.cloud/badwolf/server/internal/configs"
	"strconv"
)

const typeChunkSize = 8

func getTLSConfig() (*tls.Config, error) {
	cwd, err := os.Executable()
	if err != nil {
		return nil, fmt.Errorf("Failed to get current working directory: %v\n", err)
	}

	execPwd := filepath.Dir(cwd)

	cert, err := tls.LoadX509KeyPair(execPwd+"/configs/ca.crt.pem", execPwd+"/configs/ca.key.pem")
	if err != nil {
		return nil, fmt.Errorf("Failed to load certificate %v\n", err)
	}

	config := tls.Config{
		Certificates: []tls.Certificate{cert},
	}

	return &config, nil
}

func Listen() {
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Failed LoadConfig: %v", err)
	}

	certConfig, err := getTLSConfig()
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

	buffer := make([]byte, typeChunkSize)
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
			messageHandler(buffer[:8], &conn)
		}
	}
}

func messageHandler(bytes []byte, conn *net.Conn) {
	switch string(bytes) {
	case protocol.PingPongType:
		for i := 1; i <= 256; i++ {
			fmt.Println(i)
		}
	default:
		fmt.Printf("Unknown message\n")
	}
}
