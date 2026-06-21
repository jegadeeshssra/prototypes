package server

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"net"
	"strings"
)

func StatusCode_200(conn net.Conn, message string) {
	if message != "" {
		fmt.Printf("Status Code - 200 and Message - %s", message)
	}
	data := "HTTP/1.1 200 OK\r\n\r\n"
	WriteTCPResponse(conn, data)
}

func StatusCode_201(conn net.Conn, message string) {
	if message != "" {
		fmt.Printf("Status Code - 201 and Message - %s", message)
	}
	data := "HTTP/1.1 201 Created\r\n\r\n"
	WriteTCPResponse(conn, data)
}

func StatusCode_404(conn net.Conn, message string) {
	if message != "" {
		fmt.Printf("Status Code - 404 and Error - %s", message)
	}
	data := "HTTP/1.1 404 Not Found\r\n\r\n"
	WriteTCPResponse(conn, data)
}

func StatusCode_500(conn net.Conn, message string) {
	if message != "" {
		fmt.Printf("Status Code - 500 and Error - %s", message)
	}
	data := "HTTP/1.1 500 Internal Server Error\r\n\r\n"
	WriteTCPResponse(conn, data)
}

func GzipCompression(data string) ([]byte, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	// The compressor doesn't write to the buffer immediately on every byte. It compresses the data in chunks (blocks) and writes compressed chunks to the buffer once a block is full enough.
	if _, err := gz.Write([]byte(data)); err != nil {
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
	return buf.Bytes(), nil

}

func WriteTCPResponse(conn net.Conn, data string) bool {
	defer conn.Close()
	_, err := conn.Write([]byte(data))
	if err != nil {
		fmt.Println("Error Writing data to the accepted connection ", err.Error())
		return false
	}
	fmt.Println("\n----------------------")
	return true
}

func WritePersistentTCPResponse(conn net.Conn, data string) bool {
	fmt.Println("data - ", data)
	_, err := conn.Write([]byte(data))
	if err != nil {
		fmt.Println("Error Writing data to the accepted connection ", err.Error())
		return false
	}
	fmt.Println("\n----------------------")
	return true
}

func GzipResponse(conn net.Conn, requeststr string) {
	lines := strings.Split(requeststr, "\r\n")
	fmt.Println("----------------------")
	fmt.Println(lines)
	techniques := GetAcceptEncoding(requeststr)
	fmt.Println("techniques - ", techniques)
	if techniques != nil {
		for _, each := range techniques {
			technique := strings.TrimSpace(each)
			// fmt.Printf("each: %q\n", each) // To know exactly whats in the each and technique
			// fmt.Printf("technique: %q\n", technique)
			if technique == "gzip" {

				// get the str from PATH
				url := GetURLPath(requeststr)
				str := strings.TrimPrefix(url, "/echo/")
				// Gzip the str
				compressed_data, err := GzipCompression(str)
				if err != nil {
					fmt.Println("Cannot Compress the data using Gzip - ", err)
					return
				}

				data := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Encoding: %s\r\nContent-Length: %d\r\n\r\n", technique, len(compressed_data))
				var buf bytes.Buffer  // Used dynamic buffer
				buf.WriteString(data) // to write string and binary
				buf.Write(compressed_data)
				fmt.Printf("% X", buf.Bytes()) // To display in hexadecimal

				WriteBinaryTCPResponse(conn, buf.Bytes())
				return
			}
		}
	} else {
		StatusCode_200(conn, "")
		return
	}
}

func WriteBinaryTCPResponse(conn net.Conn, data []byte) bool {
	defer conn.Close()
	_, err := conn.Write(data)
	if err != nil {
		fmt.Println("Error Writing data to the accepted connection ", err.Error())
		return false
	}
	fmt.Println("\n----------------------")
	return true
}
