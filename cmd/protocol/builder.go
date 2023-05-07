package protocol

import (
	"fmt"
	"net"
	"seems.cloud/badwolf/server/cmd/helpers"
	"seems.cloud/badwolf/server/cmd/protocol/definitions"
)

func protoDataBuilder(data []byte) []byte {
	var message definitions.ProtoDataTransfer

	copy(message.Type[:], definitions.ProtoDataTransferType)
	copy(message.Checksum[:], helpers.Sha512Checksum(data))
	copy(message.Payload[:], data)

	return protoDataBytes(message)
}

func protoDataBytes(p definitions.ProtoDataTransfer) []byte {
	var buf []byte
	buf = append(buf, p.Type[:]...)
	buf = append(buf, p.Payload[:]...)
	return buf
}

func protoDataHandler(bytes []byte, conn *net.Conn) {
	switch string(bytes) {
	case definitions.ProtoDataTransferType:
		fmt.Printf("Data Transfer\n")
	case definitions.ProtoLinuxToExecType:
		fmt.Printf("Linux To Exec\n")
	default:
		fmt.Printf("Unknown message\n")
	}
}
