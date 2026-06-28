package main

import (
	"github.com/codecrafters-io/http-server-starter-go/app/server"
)

func main() {
	router := server.NewServer()

	router.AddRoute("/", server.DefaultPath, "GET")
	router.AddRoute("/echo/{str}", server.Echo, "GET")
	router.AddRoute("/user-agent", server.UserAgentHeader, "GET")
	router.AddRoute("/files/{filename}", server.RetrieveFiles, "GET")
	router.AddRoute("/files/{filename}", server.ReadWriteRequestBody, "POST")

	router.Start()

}
