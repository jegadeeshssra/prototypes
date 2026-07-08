package main

import (
	server "github.com/codecrafters-io/http-server-starter-go/app/server"
	service "github.com/codecrafters-io/http-server-starter-go/app/service"
)

func main() {
	router := server.NewServer()

	router.AddRoute("/", service.DefaultPath, "GET")
	router.AddRoute("/echo/{str}", service.Echo, "GET")
	router.AddRoute("/user-agent", service.UserAgentHeader, "GET")
	router.AddRoute("/files/{filename}", service.RetrieveFiles, "GET")
	router.AddRoute("/files/{filename}", service.ReadWriteRequestBody, "POST")

	router.Use(server.TimingMiddleware)
	router.Use(server.LoggingMiddleware)

	router.Start()
}
