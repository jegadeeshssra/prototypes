package server

import (
	"fmt"
	"net"
	"os"
)

type Route struct {
	Function func(request HTTPReq) HTTPResponse
	Method   string
	Path     string
}

type Server struct {
	Routes []Route
}

func NewServer() Server {
	return Server{
		Routes: make([]Route, 0),
	}
}

// This will take the routes array and append the new route within same loc of routes array.
func (server Server) AddRoute(path string, function func(HTTPReq) HTTPResponse, method string) {
	server.Routes = append(server.Routes, Route{
		Function: function,
		Path:     path,
		Method:   method,
	})
}

func (Server Server) writeConnection(conn net.Conn, data []byte) {
	defer conn.Close()
	_, err := conn.Write([]byte(data))
	if err != nil {
		fmt.Println("Error Writing data to the accepted connection ", err.Error())
	}
	fmt.Println("\n----------------------")
}

func (server Server) handleConnection(conn net.Conn) {
	buf := make([]byte, 1024) // makes a byte array of 1024 bytes

	for {
		bufLen, _ := conn.Read(buf)
		if bufLen > 0 {
			requeststr := string(buf[:bufLen]) // converts the "bufLen" bytes that Read() actually filled
			req := HTTPReq{}
			res := req.ReadAndProcessRequest(requeststr, &server.Routes)
			server.writeConnection(conn, res)
		}
		if bufLen <= 0 {
			break
		}
	}
}

func (server Server) Start() {
	listener, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind port 4221")
		os.Exit(1)
	}
	fmt.Println("HTTP Server is listening on PORT : 4221")
	defer listener.Close()

	for {
		conn, err := listener.Accept() // A net.Conn object representing the established TCP connection
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}
		go server.handleConnection(conn)
	}
}
