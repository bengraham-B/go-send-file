package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	filePath := "send.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("1.Error opening file:", err)
		return
	}
	defer file.Close()

	conn, err := net.Dial("tcp", "192.168.101.23:3043")
	if err != nil {
		fmt.Println("2.Error connecting:", err)
		return
	}
	defer conn.Close()

	// Send file name first
	_, _ = fmt.Fprintf(conn, "%s\n", filePath)

	// Send file content
	_, err = io.Copy(conn, file)
	if err != nil {
		fmt.Println("3.Error sending file:", err)
		return
	}

	fmt.Println("File sent successfully!")
}
