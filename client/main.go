package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"reflect"
	"rpc/scheme"
	"strconv"
	"strings"
)

func main() {

	//create a reader to get input from user
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Method: ")
	method, _ := reader.ReadString('\n')
	method = strings.TrimSpace(method)
	info, exists := scheme.Methods[method]
	if !exists {
		fmt.Println("Unknown method")
		return
	}

	fmt.Print("Arguments: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	args := strings.Fields(input)
	args_length := len(args)
	if args_length == 0 {
		fmt.Println("No arguments provided")
		return
	}
	if info.Count != -1 && args_length != info.Count {
		fmt.Println("Invalid number of arguments")
		return
	}
	inputArgs := make([]json.RawMessage, len(args))

	//create an []json.RawMessage with json.Marshal
	for i, raw := range args {
		switch info.ArgType.Kind() {
		case reflect.Int:
			value, err := strconv.Atoi(raw)
			if err != nil {
				fmt.Println("Invalid integer:", raw)
				return
			}
			data, err := json.Marshal(value)
			if err != nil {
				fmt.Println("Error Encoding argument: ", err)
				return
			}
			inputArgs[i] = data
		case reflect.Float64:
			value, err := strconv.ParseFloat(raw, 64)
			if err != nil {
				fmt.Println("Invalid float:", raw)
				return
			}
			data, err := json.Marshal(value)
			if err != nil {
				fmt.Println("Error encoding argument: ", err)
				return
			}
			inputArgs[i] = data

		case reflect.String:
			data, err := json.Marshal(raw)
			if err != nil {
				fmt.Println("Error encoding argument: ", err)
				return
			}
			inputArgs[i] = data
		}
	}

	//create request
	req := scheme.Request{
		Method: method,
		Args:   inputArgs,
	}

	//create connection to the server
	conn, err := net.Dial("tcp", "localhost:3000")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	//sends a line of text
	fmt.Println("Sending request : ", req)
	//send request
	err = json.NewEncoder(conn).Encode(req)
	if err != nil {
		fmt.Println("\nFailed to send request : ", err)
		return
	}
	var response scheme.Response

	//decode the response
	err = json.NewDecoder(conn).Decode(&response)
	if err != nil {
		fmt.Println("\nFailed to read response :", err)
		return
	}
	fmt.Printf("\nResponse status : %v \nResponse Value : %v\nResponse Error : %v\n", response.Status, response.Value, response.Error)

}
