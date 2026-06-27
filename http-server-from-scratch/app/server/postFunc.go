package server

import (
	"fmt"
	"os"
)

func ReadWriteRequestBody(request HTTPReq) HTTPResponse {

	contents := request.Body
	filename := request.Url.Parameters["filename"]
	fullPath := fmt.Sprintf("%s%s", dirPath, filename)
	// checks if the file already exists
	if !FileExists(fullPath) {
		err := os.WriteFile(fullPath, contents, 0644)
		if err != nil {
			fmt.Println("\nNot able to create file nor write contents into the file")
			return HTTPResponse{
				StatusCode: StatusNotFound,
			}
		}
		return HTTPResponse{
			StatusCode: StatusCreated,
		}
	} else {
		fmt.Println("\nFile already exists")
		return HTTPResponse{
			StatusCode: StatusNotFound,
		}
	}
}
