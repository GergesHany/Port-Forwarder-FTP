package main

import (
	"fmt"
	"io"
	"net"
)

const (
	senderAddress   = "localhost:2121"   // Sender server address
	receiverAddress = "localhost:2122"   // Receiver server address
)

func main() {
	// Connect to the sender server
	senderConn, err := net.Dial("tcp", senderAddress)
	if err != nil {
		fmt.Println("Error connecting to sender:", err)
		return
	}
	defer senderConn.Close()
	fmt.Println("Connected to sender server")

	// Connect to the receiver server
	receiverConn, err := net.Dial("tcp", receiverAddress)
	if err != nil {
		fmt.Println("Error connecting to receiver:", err)
		return
	}
	defer receiverConn.Close()
	fmt.Println("Connected to receiver server")

	// Forward the file from sender to receiver
	_, err = io.Copy(receiverConn, senderConn)
	if err != nil {
		fmt.Println("Error forwarding file:", err)
		return
	}
	fmt.Println("File forwarded successfully")
}