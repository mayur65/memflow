package server

import (
	"fmt"
	"github.com/mayur65/memflow/internal/protocol"
	"github.com/mayur65/memflow/internal/storage"
	"net"
)

func Start(address string) {
	listener, _ := net.Listen("tcp", address)

	fmt.Println("Memflow started on ... " + address)

	db := storage.InitDB()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: " + err.Error())
		}
		go handleClient(conn, db)
	}
}

func handleClient(conn net.Conn, db *storage.DB) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	//Make this a stream later?
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading: " + err.Error())
	}

	command := string(buffer[:n])
	fmt.Println(command)

	cmd, _ := protocol.ParseCommand(command)

	var response = ""

	if cmd.Name == "GET" {
		response = db.Get(cmd.Args[0])
	} else if cmd.Name == "SET" {
		response = db.Set(cmd.Args[0], cmd.Args[1])
	} else {
		response = "Unknown command: " + cmd.Name
	}
	_, _ = conn.Write([]byte(response))
}
