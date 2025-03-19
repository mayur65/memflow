package main

import (
	"github.com/mayur65/memflow/internal/config"
	"github.com/mayur65/memflow/internal/server"
	"log"
)

func main() {
	log.Print("Starting Memflow ...")
	server.Start(config.ServerPort)
}
