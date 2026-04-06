package main

import (
	"fmt"
	"net"
	"os"
	"bufio"
	"io"
)

func main(){

	// go run main.go 8080 
	if(len(os.Args) < 2){
		fmt.Println("Usage : go run main.go <port>")
		os.Exit(1) // Exits the program with success(1)
	}

	port := fmt.Sprintf(":%s",os.Args[1]) //returns the formatted string without printing it, allowing you to store or manipulate the result

	listener , err := net.Listen("tcp",port)
	if err != nil {
		fmt.Println("Failed to create a listerner : ",err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Printf("Listening on %s\n",listener.Addr())

	for {
		conn , err := listener.Accept()
		if err != nil{
			fmt.Println("Failed to accept the connection, err:",err)
		}

		// create a goroutine for each connection
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn){
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		data , err := reader.ReadBytes('\n')
	if err != nil {
		if err != io.EOF {
			fmt.Println("failed to read data, err : ",err)
			}
		return
		}
	fmt.Printf("Request : %s", data)
	line := fmt.Sprintf("ECHO : %s", data) // data type is converted from byte to string
	fmt.Printf("Response : %s",line)

	_ , write_err := conn.Write([]byte(line)) // This converts the string(line) into its UTF-8 byte representation.
	if write_err != nil {
		fmt.Println("failed to write data, err: ",write_err)
		return
		}
	}
}