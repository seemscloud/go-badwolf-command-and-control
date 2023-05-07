package protocol

import (
	"crypto/tls"
	"fmt"
	"net"
	"time"
)

const typeChunkSize = 128
const serverAddr = ":11112"
const retryPeriod = 1 * time.Second

func handleClientConn() net.Conn {
	for {
		conn, err := tls.Dial("tcp", serverAddr, tlsNonCertConfig())
		if err != nil {
			fmt.Printf("Failed to connect to server: %v\n", err)
			time.Sleep(retryPeriod)
			continue
		}
		return conn
	}
}

func handleServerConnection(conn net.Conn) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("Error closing connection: %s\n", err.Error())
		}
	}(conn)

	buffer := make([]byte, typeChunkSize)
	for {
		n, err := protoDataReceive(&conn, &buffer)

		if err != nil {
			fmt.Printf("Failed to receive data %v\n", err)
		} else {
			if n > 0 {
				protoDataHandler(buffer[:8], &conn)
			}
		}
	}
}
