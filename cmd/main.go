package main

import (
	"fmt"
	"github.com/mayur65/memflow/internal/config"
	"github.com/mayur65/memflow/internal/server"
)

func main() {
	fmt.Println("Starting Memflow ...")
	server.Start(config.ServerPort)
}
