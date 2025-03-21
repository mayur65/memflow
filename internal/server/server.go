package server

import (
	"github.com/mayur65/memflow/internal/protocol"
	"github.com/mayur65/memflow/internal/storage"
	"log"
	"net"
	"os"
	"time"
)

func Start(address string) {

	logFile, err := os.OpenFile("output.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	defer logFile.Close()

	// Set log output to file
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile) // Adds timestamps & file names

	listener, _ := net.Listen("tcp", address)

	log.Print("Memflow started on ... " + address)
	defer listener.Close()

	db := storage.InitDB()

	err = db.LoadRDB("save.rdb")

	if err != nil {
		log.Print(err)
		log.Print("Initializing new db, no saved RDB to load.")
	} else {
		log.Print("Loaded saved RDB.")
	}

	go periodicSave(db)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Error accepting connection: " + err.Error())
		}
		go handleClient(conn, db)
		go db.PeriodicCleaning()
	}
}

func handleClient(conn net.Conn, db *storage.DB) {
	defer conn.Close()

	buffer := make([]byte, 1024)

	//Make this a stream later?
	n, err := conn.Read(buffer)
	if err != nil {
		log.Print("Error reading: " + err.Error())
	}

	command := string(buffer[:n])
	log.Print(command)

	cmd, _ := protocol.ParseCommand(command)

	response := executeCommand(&cmd, db)
	_, _ = conn.Write([]byte(response + "\n"))
}

func executeCommand(cmd *protocol.Command, db *storage.DB) string {

	var response string

	switch cmd.Name {
	case "GET":
		response = db.Get(cmd.Args[0])
	case "SET":
		response = db.Set(cmd.Args[0], cmd.Args[1])
	case "DEL":
		response = db.Delete(cmd.Args[0])
	default:
		response = "Unknown command: " + cmd.Name
	}

	return response
}

func periodicSave(db *storage.DB) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		err := db.SaveRDB("save.rdb")
		if err != nil {
			log.Printf("Error saving RDB: %v", err)
		}
	}

}
