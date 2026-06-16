package methods

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/app/server"
)

func ReadWriteRequestBody(conn net.Conn, requestStr string) {
	lines := strings.Split(requestStr, "\r\n")
	lastLine := lines[len(lines)-1]
	// loops through each line
	for _, line := range lines {
		// checks if the line has the /files/ prefix
		if strings.HasPrefix(line, "Content-Length: ") {
			subStr := strings.TrimPrefix(strings.TrimSpace(line), "Content-Length: ")
			contentLen, err := strconv.Atoi(strings.TrimSpace(subStr))
			if err != nil {
				fmt.Println("\nError converting str to int - ", err)
				server.StatusCode_404(conn, "")
				return
			}
			body := strings.TrimSpace(lastLine)
			// checks if the content length is same as it is mentioned
			if len(body) == contentLen {
				contents := []byte(body)
				filename := server.GetFileName(conn, requestStr)
				fullPath := server.GetFullPathWithFilename(filename)
				// checks if the file already exists
				if !FileExists(fullPath) {
					err := os.WriteFile(fullPath, contents, 0644)
					if err != nil {
						fmt.Println("\nNot able to create file nor write contents into the file")
						server.StatusCode_404(conn, "")
						return
					}
					server.StatusCode_201(conn, "File Created")
					return
				} else {
					fmt.Println("\nFile already exists")
					server.StatusCode_404(conn, "")
					return
				}
			}
		} 
	}
	fmt.Println("\n\nRequest does not contain Content-Length: Header")
	server.StatusCode_404(conn,"")
	return
}
