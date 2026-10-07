package main

import (
	"encoding/json"
	"fmt"
	"net"
	"rpc/scheme"
)

func handleConn(conn net.Conn) {
	//close the connection when done handling the request
	defer conn.Close()

	var req scheme.Request
	//this will read bytes until newline is recieved
	//we use json.NewDecoder bec it reads JSON from io.Reader while unmarshall reads json from memory
	err := json.NewDecoder(conn).Decode(&req)
	if err != nil {
		fmt.Println("Error reading request : ", err)
		return
	}
	fmt.Println("Client says: ", req)
	//validate request
	method, err := scheme.ValidateRequest(req)
	if err != nil {
		json.NewEncoder(conn).Encode(scheme.Response{
			Status: 400,
			Value:  nil,
			Error:  err.Error(),
		})
		return
	}

	response := handleReq(method)
	err = json.NewEncoder(conn).Encode(response)
	if err != nil {
		json.NewEncoder(conn).Encode(scheme.Response{
			Status: 400,
			Value:  nil,
			Error:  err.Error(),
		})
		return
	}
}

func main() {
	//opens a TCP listener on a specific port
	listener, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Println("Listening on localhost:3000")

	for {
		//Accept detects when a client connects to the server
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConn(conn)
	}
}
