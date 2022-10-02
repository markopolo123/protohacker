package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net"
	"os"
)

const (
	HOST = "0.0.0.0"
	PORT = "3333"
	TYPE = "tcp"
)

type Request struct {
	Method *string  `json:"method"`
	Number *float64 `json:"number"`
}

type Response struct {
	Method string `json:"method"`
	Prime  bool   `json:"prime"`
	Error  string `json:"error"`
}

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

	var reader = bufio.NewReader(conn)

	for {
		bytes, err := reader.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Printf("Error reading: %#v", err)
			}
			return
		}
		fmt.Printf("Message received: %s", string(bytes))
		request, err := parse(bytes)
		if err != nil {
			response := Response{
				Error: "Invalid JSON",
			}
			sendResponse(conn, response)
			return
		}
		response := handleCommand(request)
		sendResponse(conn, response)
	}
}

func parse(input []byte) (Request, error) {
	var r Request

	err := json.Unmarshal(input, &r)

	if err != nil {
		fmt.Printf("Error unmarshaling input: %#v", err)
		return r, err
	} else if r.Method == nil {
		return r, errors.New("Method not provided")
	} else if r.Number == nil {
		return r, errors.New("Number not provided")
	}

	return r, nil
}

func sendResponse(conn net.Conn, response Response) {
	rawResponse, err := json.Marshal(response)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Printf("Responding: %s\n", string(rawResponse))

	_, err = conn.Write(append(rawResponse, '\n'))
	if err != nil {
		fmt.Printf("Error writing: %#v\n", err)
	}
}

func handleCommand(request Request) Response {
	var response Response

	if *request.Method == "isPrime" {
		response = Response{
			Method: "isPrime",
			Prime:  big.NewInt(int64(*request.Number)).ProbablyPrime(0),
		}
	} else {
		response = Response{
			Error: "Unsupported Method",
		}
	}
	return response
}
