package main

import (
	"log"

	"github.com/xxidbr9/livekit-server-for-electron/cmd/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
