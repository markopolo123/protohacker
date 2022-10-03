package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
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
	defer conn.Close()
	prices := make(map[int32]int32)

	for {
		req, err := ioutil.ReadAll(io.LimitReader(conn, 9))
		if err != nil {
			fmt.Errorf("error reading bytes: %w", err)
			return
		}
		parseResult, err := parse(req, prices)
		if err != nil {
			fmt.Errorf("error parsing: %w", err)
		}
		if parseResult == nil {
			continue
		}
		if _, err := conn.Write(*parseResult); err != nil {
			fmt.Errorf("failed to write parse result to connection: %w", err)
		}
	}
}

func parse(bytes []byte, prices map[int32]int32) (*[]byte, error) {
	if len(bytes) < 9 {
		return nil, nil
	}
	messageType := bytes[0]
	messageOne := int32(big.NewInt(0).SetBytes(bytes[1:5]).Int64())
	messageTwo := int32(big.NewInt(0).SetBytes(bytes[5:]).Int64())

	switch messageType {
	case 'I':
		prices[messageOne] = messageTwo
	case 'Q':
		sum := int64(0)
		count := int64(0)
		for time, price := range prices {
			if time >= messageOne && time <= messageTwo {
				sum += int64(price)
				count += 1
			}
		}

		var avg int32
		if count != 0 {
			avg = int32(sum / count)
		}

		var b [4]byte
		binary.BigEndian.PutUint32(b[:], uint32(avg))
		r := b[:]
		return &r, nil
	}

	return nil, nil
}
