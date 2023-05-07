package main

import (
	"fmt"
	"os"
	"seems.cloud/badwolf/server/cmd/protocol"
	"seems.cloud/badwolf/server/cmd/protocol/definitions"
	"seems.cloud/badwolf/server/cmd/ui"
)

func main() {
	if len(os.Args) < 2 {
		go protocol.ClientHandler()
		go ui.HttpServer()
		go protocol.ServerHandler()
	} else {
		switch os.Args[1] {
		case "client":
			go protocol.ClientHandler()
		case "command2":
			go ui.HttpServer()
			go protocol.ServerHandler()
		}
	}

	definitions.ProtoLinuxToExecChannel <- fmt.Sprintf("ls -lh")

	select {}
}
