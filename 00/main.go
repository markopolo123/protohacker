package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

const (
	HOST = "0.0.0.0"
	PORT = "3333"
	TYPE = "tcp"
)

func main() {
	// Listen for incoming connections
	l, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes
	defer l.Close()
	fmt.Println("Listening on " + HOST + ":" + PORT)
	for {
		// Listen for an incoming connection
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("Got connection from: ", conn.RemoteAddr())

		// Handle connections in a new goroutine
		go handleRequest(conn)
	}
}

// Handles incoming requests
func handleRequest(conn net.Conn) {
	// close conn later
	defer conn.Close()
	// io.Copy(conn, conn)
	//read
	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error: ", err)
			}
			fmt.Printf("Writing last n bytes: %d\n", n)
			conn.Write(buf[:n])
			break
		}
		fmt.Printf("Writing bytes: %d\n", n)
		conn.Write(buf[:n])
	}
	fmt.Println("Connection closed")
}
