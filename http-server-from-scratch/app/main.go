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
	//fmt.Println(requeststr)
	if method == "GET" {
		cleanPath := strings.TrimSpace(path)
		//fmt.Println(cleanPath)
		if cleanPath == "/" {
			return methods.DefaultPath(conn)
		} else if strings.HasPrefix(cleanPath, "/echo/") {
			return methods.EchoPathStr(conn, cleanPath)
		} else if strings.HasPrefix(cleanPath, "/user-agent") {
			return methods.UserAgentHeader(conn, requeststr)
		} else if connHead, _ := server.ConnectionHeader(requeststr); connHead == "close" {
			return server.ClosePersistentConnection(conn)
		}
	}
	data := "HTTP/1.1 404 Not Found\r\n\r\n"
	return server.WritePersistentTCPResponse(conn, data)
}

func handleConnection(conn net.Conn) {
	buf := make([]byte, 1024) // makes a byte array of 1024 bytes

	for {
		bufLen, _ := conn.Read(buf)
		if bufLen > 0 {
			requeststr := string(buf[:bufLen]) // converts the "bufLen" bytes that Read() actually filled
			_ = router(conn, requeststr)
		}
		if bufLen <= 0 {
			break
		}
	}
}

func main() {

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind port 4221")
		os.Exit(1)
	}
	fmt.Println("HTTP Server is listening on PORT : 4221")
	defer l.Close()

	for {
		conn, err := l.Accept() // A net.Conn object representing the established TCP connection
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}
		go handleConnection(conn)
	}
}

routes := make([]Route,0)
routes = append(routes,
	Route{
		Fucntion: methods.DefaultPath,
		Method: "GET",
		Path: "/"
	}),
	Route{
		Fucntion: methods.Ec,
		Method: "GET",
		Path: "/"
	}
