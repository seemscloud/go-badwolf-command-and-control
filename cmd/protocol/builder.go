package protocol

import (
	"fmt"
	"net"
	"seems.cloud/badwolf/server/cmd/protocol/definitions"
)

func dataBuilder(dataType string, data []byte) []byte {
	var message definitions.ProtoDataTransfer

	copy(message.Type[:], dataType)
	copy(message.Payload[:], data)

	return dataBytes(message)
}

func dataBytes(p definitions.ProtoDataTransfer) []byte {
	var buf []byte
	buf = append(buf, p.Type[:]...)
	buf = append(buf, p.Payload[:]...)
	return buf
}

func dataHandler(bytes []byte, conn *net.Conn) {
	switch string(bytes) {
	case definitions.ProtoDataTransferType:
		fmt.Printf("Data Transfer\n")
	case definitions.ProtoLinuxToExecType:
		fmt.Printf("Linux To Exec\n")
	default:
		fmt.Printf("Unknown message\n")
	}
}
