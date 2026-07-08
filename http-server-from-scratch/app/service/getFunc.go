package server

import (
	"fmt"
	"os"

	http "github.com/codecrafters-io/http-server-starter-go/app/http"
)

// Endpoint - "/"
func DefaultPath(request http.HTTPReq) http.HTTPResponse {
	return http.HTTPResponse{
		StatusCode: http.StatusOK,
	}
}

// Endpoint - "/echo/{str}"
func Echo(request http.HTTPReq) http.HTTPResponse {
	content := request.Url.Parameters["str"]

	return http.HTTPResponse{
		StatusCode: http.StatusOK,
		Headers:    map[string]string{"Content-Type": "text/plain"},
		Body:       []byte(content),
	}
}

func UserAgentHeader(request http.HTTPReq) http.HTTPResponse {
	content := request.Headers["User-Agent"]
	return http.HTTPResponse{
		StatusCode: http.StatusOK,
		Headers:    map[string]string{"Content-Type": "text/plain"},
		Body:       []byte(content),
	}
}

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

func RetrieveFiles(request http.HTTPReq) http.HTTPResponse {

	filename := request.Url.Parameters["filename"]
	fullPath := fmt.Sprintf("%s%s", dirPath, filename)
	//fmt.Println(fullPath)
	// checks if the filename is mentioned && checks the file in the system
	if !FileExists(fullPath) {
		return http.HTTPResponse{
			StatusCode: http.StatusNotFound,
		}
	} else {
		b_contents, err := os.ReadFile(fullPath)
		if err != nil {
			fmt.Println("Error reading the file - ", err)
			return http.HTTPResponse{
				StatusCode: http.StatusNotFound,
			}
		}
		content := string(b_contents)
		return http.HTTPResponse{
			StatusCode: http.StatusOK,
			Headers:    map[string]string{"Content-Type": "text/plain"},
			Body:       []byte(content),
		}
	}
}
