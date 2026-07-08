package server

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"log"
	"strings"
)

// func ClosePersistentConnection(conn net.Conn) bool {
// 	data := "HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nConnection: close\r\n\r\n"
// 	_ = WritePersistentTCPResponse(conn, data)
// 	defer conn.Close()
// 	return true
// }

type HTTPResponse struct {
	Headers    map[string]string
	StatusCode int
	Body       []byte
	Request    *HTTPReq
}

const (
	StatusOK       = 200
	StatusCreated  = 201
	StatusNotFound = 404
)

func StatusText(statusCode int) string {
	switch statusCode {
	case StatusOK:
		return "OK"
	case StatusCreated:
		return "Created"
	case StatusNotFound:
		return "Not Found"
	}
	return ""
}

func (response HTTPResponse) gzipCompression(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	// The compressor doesn't write to the buffer immediately on every byte. It compresses the data in chunks (blocks) and writes compressed chunks to the buffer once a block is full enough.
	if _, err := gz.Write(data); err != nil {
		return nil, err
	}
	// The gzip compression algorithm usually waits to collect a certain amount of data before creating a compressed block. Flush() ensures that all compressed data reaches the buffer before we finalize the file.
	if err := gz.Flush(); err != nil {
		return nil, err
	}
	// It writes the gzip footer to the buffer. This footer contains a CRC-32 checksum of the original, uncompressed data and the original data's length.
	if err := gz.Close(); err != nil {
		return nil, err
	}

	//return base64.StdEncoding.EncodeToString(buf.Bytes()), nil
	// LOOK the reduntant type conversion
	return buf.Bytes(), nil
}

func (response HTTPResponse) Write(request HTTPReq) []byte {
	// GZIP Compression if the req has its headers
	if encodingstr, ok := request.Headers["Accept-Encoding"]; ok {
		encodings := strings.Split(encodingstr, ",")
		for _, encoding := range encodings {
			if strings.TrimSpace(encoding) == "gzip" {
				if len(response.Body) > 0 {
					gzipContent, err := response.gzipCompression(response.Body)
					if err != nil {
						log.Fatal(err)
					}
					response.Body = gzipContent
					break
				}
			}
		}
	}
	// REQUEST LINE
	str := fmt.Sprintf("HTTP/1.1 %d %s\r\n", response.StatusCode, StatusText(response.StatusCode))
	// HEADERS LINE
	for key, value := range response.Headers {
		str += fmt.Sprintf("%s: %s\r\n", key, value)
	}
	str += "\r\n"
	// RESPONSE BODY CONTENT
	if len(response.Body) > 0 {
		str += string(response.Body)
	}
	return []byte(str)
}
