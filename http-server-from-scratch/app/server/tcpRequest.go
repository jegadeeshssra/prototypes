package server

import (
	"fmt"
	"strconv"
	"strings"
)

//var DirPath string = "C:\\Users\\018046\\OneDrive - Sify Technologies Limited\\go\\prototypes\\codecrafters-http-server-go\\files\\"

type HTTPReq struct {
	Headers map[string]string
	URL     string
	Method  string
	Body    []byte
}

func ReadRequest(requeststr string) HTTPReq {
	reqSplit := strings.Split(requeststr, "\r\n\r\n")
	lines := strings.Split(reqSplit[0], "\r\n")
	parts := strings.Split(lines[0], " ")
	// METHOD
	method := strings.TrimSpace(parts[0])
	// URL
	path := strings.TrimSpace(parts[1])
	// HEADERS
	headers := make(map[string]string)
	for i := 1; i < len(lines); i++ {
		headerParts := strings.Split(lines[i], ":")
		if len(headerParts) >= 2 {
			headers[headerParts[0]] = strings.Join(headerParts[1:], "")
		}
	}
	// Content Length to extract the BODY CONTENT
	contentLength, err := strconv.Atoi(headers["Content-Length"])
	if err != nil {
		fmt.Println("Could not convert content length to int, ignoring body")
		contentLength = 0
	}

	request := HTTPReq{
		Headers: headers,
		URL:     path,
		Method:  method,
		Body:    []byte(reqSplit[1][:contentLength]),
	}

	return request
}
