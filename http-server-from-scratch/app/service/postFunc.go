package server

import (
	"fmt"
	"os"

	http "github.com/codecrafters-io/http-server-starter-go/app/http"
)

func ReadWriteRequestBody(request http.HTTPReq) http.HTTPResponse {

	contents := request.Body
	filename := request.Url.Parameters["filename"]
	fullPath := fmt.Sprintf("%s%s", dirPath, filename)
	// checks if the file already exists
	if !FileExists(fullPath) {
		err := os.WriteFile(fullPath, contents, 0644)
		if err != nil {
			fmt.Println("\nNot able to create file nor write contents into the file")
			return http.HTTPResponse{
				StatusCode: http.StatusNotFound,
			}
		}
		return http.HTTPResponse{
			StatusCode: http.StatusCreated,
		}
	} else {
		fmt.Println("\nFile already exists")
		return http.HTTPResponse{
			StatusCode: http.StatusNotFound,
		}
	}
}
