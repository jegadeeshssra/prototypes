package main

import (
	"fmt"
	"net"
	"os"
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
	defer listerner.Close()
	fmt.Println("Listening on %s%s",listener.Addr(),port)

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
	defer conn.close()

}