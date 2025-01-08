package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath" // to work with file paths
	"time"
)

const (
	downloadDir = "download" // Folder to save downloaded files
	port        = ":2122"    // Receiver server port
)

func main() {

	// Create the download directory if it doesn't exist
	if err := os.MkdirAll(downloadDir, os.ModePerm); err != nil {
		fmt.Println("Error creating download directory:", err)
		return
	}

	// Start the receiver server
	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Error starting receiver server:", err)
		return
	}

	defer listener.Close()
	fmt.Println("Receiver server is listening on port", port)

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Generate a unique file name using a timestamp
	fileName := fmt.Sprintf("received_file_%d.txt", time.Now().Unix())
	filePath := filepath.Join(downloadDir, fileName)

	// Create a new file in the download directory
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Save the received data to the file
	_, err = io.Copy(file, conn)
	if err != nil {
		fmt.Println("Error saving file:", err)
		return
	}
	fmt.Println("File received and saved as", filePath)
}