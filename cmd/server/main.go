package main

import (
	"flag"
	"log"

	"github.com/juster/antlr-playgrounds/internal/server"
)

func main() {
	port := flag.String("port", "8080", "Server port")
	flag.Parse()

	srv := server.New(*port)
	if err := srv.Start(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
