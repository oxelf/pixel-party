package main

import (
	"log"

	"github.com/oxelf/pixel-party/internal/server"
)

func main() {
    srv := server.New()
    if err := srv.Start(); err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}
