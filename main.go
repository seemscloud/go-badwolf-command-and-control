package main

import (
	"os"
	"seems.cloud/badwolf/server/cmd/protocol"
	"seems.cloud/badwolf/server/cmd/ui"
)

func main() {
	if len(os.Args) != 2 {
		go protocol.Client()

		go ui.Server()
		go protocol.Server()
	} else {
		switch os.Args[1] {
		case "client":
			go protocol.Client()
		case "server":
			go ui.Server()
			go protocol.Server()
		}
	}
	//
	//data := "ls -lh"
	//
	//transfer := definitions.ProtoDataTransfer{
	//	Type:    definitions.ProtoLinuxToExecType,
	//	Length:  [5]byte{5, 0, 0, 0, 0},
	//	Payload: []byte(data),
	//}
	//
	//definitions.ProtoDataTransferChannel <- transfer

	select {}
}
