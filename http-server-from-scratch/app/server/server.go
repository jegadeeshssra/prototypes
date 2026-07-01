package server

import (
	"fmt"
	"net"
	"os"
	"time"
)

type Route struct {
	Function func(request HTTPReq) HTTPResponse
	Method   string
	Path     string
}

type handler func(request HTTPReq) HTTPResponse

type Server struct {
	Routes      []Route
	Middlewares []func(handler) handler
	Semaphore   chan struct{}
}

func NewServer() Server {
	return Server{
		Routes:      make([]Route, 0),
		Middlewares: make([]func(handler) handler, 0),
		Semaphore:   make(chan struct{}, 1000),
	}
}

func (server Server) Use(middleware func(handler) handler) {
	server.Middlewares = append(server.Middlewares, middleware)
}

// This will take the routes array and append the new route within same loc of routes array.
func (server Server) AddRoute(path string, function func(HTTPReq) HTTPResponse, method string) {
	server.Routes = append(server.Routes, Route{
		Function: function,
		Path:     path,
		Method:   method,
	})
}

func (server Server) writeConnection(conn net.Conn, data []byte, keepAlive bool) {
	if keepAlive == false {
		defer conn.Close()
	}
	_, err := conn.Write([]byte(data))
	if err != nil {
		fmt.Println("Error Writing data to the accepted connection ", err.Error())
	}
	fmt.Println("\n----------------------")
}

func (server Server) handleConnection(conn net.Conn) {

	defer func() { <-server.Semaphore }()

	// waits idle for 30ms expecting request data
	conn.SetReadDeadline(time.Now().Add(1000 * time.Millisecond))

	buf := make([]byte, 1024) // makes a byte array of 1024 bytes

	for {
		bufLen, err := conn.Read(buf)
		if os.IsTimeout(err) {
			conn.Close()
			return
		}
		if bufLen > 0 {
			requeststr := string(buf[:bufLen]) // converts the "bufLen" bytes that Read() actually filled
			req := HTTPReq{}
			reqFunc, keepAlive := req.ReadAndProcessRequest(requeststr, &server.Routes)

			// reqFunction goes through
			for i := len(server.Middlewares) - 1; i > 0; i-- {
				reqFunc = server.Middlewares[i](reqFunc)
			}

			resDataBytes := reqFunc(req).Write(req)

			conn.SetWriteDeadline(time.Now().Add(300 * time.Millisecond))
			server.writeConnection(conn, resDataBytes, keepAlive)
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

		select {
		//acquire a slot , blocks if it is full
		case server.Semaphore <- struct{}{}:

			go server.handleConnection(conn)
		default:
			conn.Write([]byte("HTTP/1.1 503 Service Unavailable\r\nContent-Length: 15\r\n\r\nserver too busy\n"))
			conn.Close()
		}
	}
}
