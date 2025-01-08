package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

const (
	port = ":2121" // Sender server port
)

func main() {
	// Check if the file path is provided as an argument (file to be sent)
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run sender.go <file_path>")
		return
	}

	filePath := os.Args[1] // Get the file path 

	// Start the sender server
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Error starting sender server:", err)
		return
	}

	defer listener.Close()
	fmt.Println("Sender server is listening on port", port)

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn, filePath)
	}
}

func handleConnection(conn net.Conn, filePath string) {
	defer conn.Close()

	// Open the file to be sent
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Send the file
	_, err = io.Copy(conn, file)
	if err != nil {
		fmt.Println("Error sending file:", err)
		return
	}
	fmt.Println("File sent successfully:", filePath)
}