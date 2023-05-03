package main

import (
	"seems.cloud/badwolf/server/cmd/listener"
	"seems.cloud/badwolf/server/cmd/ui"
)

func main() {
	go ui.HttpServer()
	go listener.Listen()

	select {}
}
