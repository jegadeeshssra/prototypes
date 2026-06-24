package server

import (
	"fmt"
	"strconv"
	"strings"
)

//var DirPath string = "C:\\Users\\018046\\OneDrive - Sify Technologies Limited\\go\\prototypes\\codecrafters-http-server-go\\files\\"

// func GetURLPath(request string) string {
// 	parts := strings.Split(request, " ")
// 	//fmt.Println("Parts - ", parts)
// 	path := parts[1]
// 	return strings.TrimSpace(path)
// }

// func GetUserAgent(request string) (string, error) {
// 	lines := strings.Split(request, "\r\n")
// 	for _, line := range lines {
// 		// checks if the line has the /files/ prefix
// 		if strings.HasPrefix(line, "User-Agent: ") {
// 			subStr := strings.TrimPrefix(strings.TrimSpace(line), "User-Agent: ")
// 			value := strings.TrimSpace(subStr)
// 			return value, nil
// 		}
// 	}
// 	return "", fmt.Errorf("No User-Agent Header in the request")
// }

// func ConnectionHeader(request string) (string, error) {
// 	lines := strings.Split(request, "\r\n")
// 	// loops through each line
// 	for _, line := range lines {
// 		// checks if the line has the /files/ prefix
// 		if strings.HasPrefix(line, "Connection: ") {
// 			subStr := strings.TrimPrefix(strings.TrimSpace(line), "Connection: ")
// 			value := strings.TrimSpace(subStr)
// 			return value, nil
// 		}
// 	}
// 	return "", fmt.Errorf("No Connection Header")
// }

// func GetAcceptEncoding(request string) []string {
// 	lines := strings.Split(request, "\r\n")
// 	// loops through each line
// 	for _, line := range lines {
// 		// checks if the line has the /files/ prefix
// 		if strings.HasPrefix(line, "Accept-Encoding: ") {
// 			subStr := strings.TrimPrefix(strings.TrimSpace(line), "Accept-Encoding: ")
// 			techniques := strings.Split(subStr, ",")
// 			return techniques
// 		}
// 	}
// 	return nil
// }

// func GetFileName(conn net.Conn, request string) string {
// 	path := GetURLPath(request)
// 	if strings.HasPrefix(path, "/files/") {
// 		fileName := strings.TrimPrefix(path, "/files/")
// 		return fileName
// 	} else {
// 		data := "HTTP/1.1 404 Not Found\r\n\r\n"
// 		WriteTCPResponse(conn, data)
// 		return ""
// 	}
// }

// func GetFullPathWithFilename(fileName string) string {
// 	return fmt.Sprintf("%s%s", DirPath, fileName)
// }

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
