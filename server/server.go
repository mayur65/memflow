package server

import (
	"fmt"
	"net"
)

func Start(address string) {
	listener, _ := net.Listen("tcp", address)

	fmt.Println("Memflow started on ... " + address)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: " + err.Error())
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading: " + err.Error())
	}

	command := string(buffer[:n])
	fmt.Println(command)
	response := "Received command"
	_, _ = conn.Write([]byte(response))
}
