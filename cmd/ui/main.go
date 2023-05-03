package ui

import (
	"log"
	"net/http"
	"seems.cloud/badwolf/server/cmd/ui/handlers"
	"seems.cloud/badwolf/server/internal/configs"
	"strconv"
)

func HttpServer() {
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Failed LoadConfig: %v", err)
	}

	listenAddr := config.WebInterface.Host + ":" + strconv.Itoa(config.WebInterface.Port)
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    listenAddr,
		Handler: mux,
	}

	mux.HandleFunc("/api/hosts", handlers.HostsHandler)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
