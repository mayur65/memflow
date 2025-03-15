package main

import (
	"fmt"
	"github.com/mayur65/memflow/config"
	"github.com/mayur65/memflow/server"
)

func main() {
	fmt.Println("Starting Memflow ...")
	server.Start(config.ServerPort)
}
