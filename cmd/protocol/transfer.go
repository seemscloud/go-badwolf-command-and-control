package protocol

import (
	"errors"
	"fmt"
	"io"
	"net"
	"syscall"
)

func protoDataReceive(conn *net.Conn, buffer *[]byte) (int, error) {
	n, err := (*conn).Read(*buffer)
	if err != nil {
		if err == io.EOF {
			return -1, fmt.Errorf("failed to receive data, broken pipe %v", err)
		} else {
			return -1, fmt.Errorf("failed to receive data %v", err)
		}
	}

	return n, nil
}

func protoDataSend(conn *net.Conn, data []byte) {
	_, err := (*conn).Write(data)
	if err != nil {
		fmt.Println("Error writing to server:", err)

		if errors.Is(err, syscall.EPIPE) {
			fmt.Println("Broken pipe error")
		} else if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			fmt.Println("Timeout error")
		} else {
			fmt.Println("Other error")
			return
		}
		*conn = handleClientConn()
	}
}
