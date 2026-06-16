package methods

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/app/server"
)

var dirPath string = "C:\\Users\\018046\\OneDrive - Sify Technologies Limited\\go\\prototypes\\codecrafters-http-server-go\\files\\"

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
