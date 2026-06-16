package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/app/methods"
	"github.com/codecrafters-io/http-server-starter-go/app/server"
)

func handleConnection(conn net.Conn) {
	buf := make([]byte, 1024) // makes a byte array of 1024 bytes
	bufLen, _ := conn.Read(buf)
	requeststr := string(buf[:bufLen]) // converts the "bufLen" bytes that Read() actually filled

	lines := strings.Split(requeststr, "\r\n")
	fmt.Println("----------------------")
	parts := strings.Split(lines[0], " ")
	method := strings.TrimSpace(parts[0])

	techniques := server.GetAcceptEncoding(requeststr)
	fmt.Println("techniques - ", techniques)
	if techniques != nil {
		for _, each := range techniques {
			technique := strings.TrimSpace(each)
			// fmt.Printf("each: %q\n", each) // To know exactly whats in the each and technique
			// fmt.Printf("technique: %q\n", technique)
			if technique == "gzip" {

				// get the str from PATH
				url := server.GetURLPath(requeststr)
				str := strings.TrimPrefix(url, "/echo/")
				// Gzip the str
				compressed_data, err := server.GzipCompression(str)
				if err != nil {
					fmt.Println("Cannot Compress the data using Gzip - ", err)
					return
				}

				data := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Encoding: %s\r\nContent-Length: %d\r\n\r\n", technique, len(compressed_data))
				var buf bytes.Buffer  // Used dynamic buffer
				buf.WriteString(data) // to write string and binary
				buf.Write(compressed_data)
				fmt.Printf("% X", buf.Bytes()) // To display in hexadecimal

				server.WriteBinaryTCPResponse(conn, buf.Bytes())
				return
			}
		}
	} else {
		server.StatusCode_200(conn, "")
		return
	}

	// Not needed for this task
	if method == "GET" {
		path := server.GetURLPath(requeststr)
		if strings.HasPrefix(path, "/files") {
			methods.RetrieveFiles(conn, requeststr)
		}
		return
	} else if method == "POST" {
		methods.ReadWriteRequestBody(conn, requeststr)
		return
	}
}

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	// fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind port 4221")
		os.Exit(1)
	}
	fmt.Println("HTTP Server is listening on PORT : 4221")
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}
		go handleConnection(conn)
	}
}
