package main

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/app/methods"
	"github.com/codecrafters-io/http-server-starter-go/app/server"
)

func router(conn net.Conn, requeststr string) bool {
	lines := strings.Split(requeststr, "\r\n")
	parts := strings.Split(lines[0], " ")
	method := strings.TrimSpace(parts[0])
	path := server.GetURLPath(requeststr)
	if method == "GET" {
		cleanPath := strings.TrimSpace(path)
		if strings.HasPrefix(cleanPath, "/echo/") {
			return methods.EchoPathStr(conn, cleanPath)
		} else if strings.HasPrefix(cleanPath, "/user-agent") {
			return methods.UserAgentHeader(conn, requeststr)
		} else if connHead, _ := server.ConnectionHeader(requeststr); connHead == "close" {
			defer conn.Close()
		}
	}
	data := "HTTP/1.1 404 Not Found\r\n\r\n"
	return server.WritePersistentTCPResponse(conn, data)
}

func handleConnection(conn net.Conn) {
	buf := make([]byte, 1024) // makes a byte array of 1024 bytes

	for {
		bufLen, _ := conn.Read(buf)
		if bufLen <= 0 {
			break
		}
		requeststr := string(buf[:bufLen]) // converts the "bufLen" bytes that Read() actually filled

		_ = router(conn, requeststr)

		// connHeader, err := server.ConnectionHeader(requeststr)
		// if err != nil {
		// 	conn.Close()
		// 	return
		// }
		// if connHeader == "Closed" {
		// 	conn.Close()
		// }
	}

	defer conn.Close()
	// lines := strings.Split(requeststr, "\r\n")
	// parts := strings.Split(lines[0], " ")
	// method := strings.TrimSpace(parts[0])

	// // // Not needed for this task
	// if method == "GET" {
	// 	path := server.GetURLPath(requeststr)
	// 	if strings.HasPrefix(path, "/files") {
	// 		methods.RetrieveFiles(conn, requeststr)
	// 	}
	// 	return
	// } else if method == "POST" {
	// 	methods.ReadWriteRequestBody(conn, requeststr)
	// 	return
	// }
}

// func HandleConnection(conn net.Conn){{
// 	buf := make([]byte,4096)
// 	bufLen , _ := conn.Read(buf)
// 	requestStr := buf[:bufLen]

// }}

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
