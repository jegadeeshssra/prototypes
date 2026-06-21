package server

import (
	"fmt"
	"net"
	"strings"
)

var DirPath string = "C:\\Users\\018046\\OneDrive - Sify Technologies Limited\\go\\prototypes\\codecrafters-http-server-go\\files\\"

func GetURLPath(request string) string {
	parts := strings.Split(request, " ")
	//fmt.Println("Parts - ", parts)
	path := parts[1]
	return strings.TrimSpace(path)
}

func GetUserAgent(request string) (string, error) {
	lines := strings.Split(request, "\r\n")
	for _, line := range lines {
		// checks if the line has the /files/ prefix
		if strings.HasPrefix(line, "User-Agent: ") {
			subStr := strings.TrimPrefix(strings.TrimSpace(line), "User-Agent: ")
			value := strings.TrimSpace(subStr)
			return value, nil
		}
	}
	return "", fmt.Errorf("No User-Agent Header in the request")
}

func ConnectionHeader(request string) (string, error) {
	lines := strings.Split(request, "\r\n")
	// loops through each line
	for _, line := range lines {
		// checks if the line has the /files/ prefix
		if strings.HasPrefix(line, "Connection: ") {
			subStr := strings.TrimPrefix(strings.TrimSpace(line), "Connection: ")
			value := strings.TrimSpace(subStr)
			return value, nil
		}
	}
	return "", fmt.Errorf("No Connection Header")
}

func GetAcceptEncoding(request string) []string {
	lines := strings.Split(request, "\r\n")
	// loops through each line
	for _, line := range lines {
		// checks if the line has the /files/ prefix
		if strings.HasPrefix(line, "Accept-Encoding: ") {
			subStr := strings.TrimPrefix(strings.TrimSpace(line), "Accept-Encoding: ")
			techniques := strings.Split(subStr, ",")
			return techniques
		}
	}
	return nil
}

func GetFileName(conn net.Conn, request string) string {
	path := GetURLPath(request)
	if strings.HasPrefix(path, "/files/") {
		fileName := strings.TrimPrefix(path, "/files/")
		return fileName
	} else {
		data := "HTTP/1.1 404 Not Found\r\n\r\n"
		WriteTCPResponse(conn, data)
		return ""
	}
}

// func GetContentLength(request string) string {
// 	lines :=strings.Split(request,"\r\n")
// 	for _ , line := range lines {
// 		cleanLine := strings.TrimSpace(line)
// 		if strings.HasPrefix(cleanLine,"Content-Length:"){

// 		}
// 	}

// }

func GetFullPathWithFilename(fileName string) string {
	return fmt.Sprintf("%s%s", DirPath, fileName)
}
