package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	listener, err := net.Listen("tcp", ":3043") // replace port with the desired port number
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting connection:", err)
		return
	}
	defer conn.Close()

	// Receive file name first
	fileNameBuffer := make([]byte, 4096)
	_, err = conn.Read(fileNameBuffer)
	if err != nil {
		fmt.Println("Error receiving file name:", err)
		return
	}
	fileName := string(fileNameBuffer)

	// Create the file
	receivedFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer receivedFile.Close()

	// Receive file content
	_, err = io.Copy(receivedFile, conn)
	if err != nil {
		fmt.Println("Error receiving file content:", err)
		return
	}

	fmt.Println("File received successfully!")
}
