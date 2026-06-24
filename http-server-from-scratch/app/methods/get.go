package methods

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/app/server"
)

var dirPath string = "C:\\Users\\018046\\OneDrive - Sify Technologies Limited\\go\\prototypes\\codecrafters-http-server-go\\files\\"

func DefaultPath(conn net.Conn) bool {
	data := "HTTP/1.1 200 OK\r\n\r\n"
	return server.WriteTCPResponse(conn, data)
}

func EchoPathStr(conn net.Conn, cleanPath string) bool {
	str := strings.TrimPrefix(cleanPath, "/echo/")
	if len(str) <= 0 {
		data := "HTTP/1.1 404 String Not Found\r\n\r\n"
		return server.WritePersistentTCPResponse(conn, data)
	}
	data := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(str), str)
	return server.WritePersistentTCPResponse(conn, data)
}

func UserAgentHeader(conn net.Conn, requeststr string) bool {
	// time.Sleep(5 * time.Second)
	headerContent, err := server.GetUserAgent(requeststr)
	if err != nil {
		fmt.Println(err)
		data := "HTTP/1.1 404 Header string Not Found\r\n\r\n"
		return server.WritePersistentTCPResponse(conn, data)
	}
	data := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(headerContent), headerContent)
	return server.WritePersistentTCPResponse(conn, data)
}

func FileExists(fullPath string) bool {
	_, err := os.Stat(fullPath)
	if err == nil {
		return true
	}
	// checks if the error specifically means "file does not exist"
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func RetrieveFiles(conn net.Conn, requestStr string) {

	lines := strings.Split(requestStr, "\r\n")
	parts := strings.Split(lines[0], " ")
	path := parts[1]
	// checks if the PATH contains "/files/"
	if strings.HasPrefix(path, "/files/") {
		fileName := strings.TrimPrefix(path, "/files/")
		fullPath := fmt.Sprintf("%s%s", dirPath, fileName)
		fmt.Println(fullPath)
		// checks if the filename is mentioned && checks the file in the system
		if fileName != "" && FileExists(fullPath) {
			b_contents, err := os.ReadFile(fullPath)
			if err != nil {
				fmt.Println("Error reading the file - ", err)
				return
			}
			contents := string(b_contents)
			data := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: application/octet-stream\r\nContent-Length: %d\r\n\r\n%s", len(contents), contents)
			server.WriteTCPResponse(conn, data)
			return
		} else {
			data := "HTTP/1.1 404 File Not Found\r\n\r\n"
			server.WriteTCPResponse(conn, data)
			return
		}
	} else {
		data := "HTTP/1.1 404 Not Found\r\n\r\n"
		server.WriteTCPResponse(conn, data)
		return
	}
}
