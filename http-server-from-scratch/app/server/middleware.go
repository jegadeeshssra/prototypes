package server

import (
	"fmt"
	"time"

	http "github.com/codecrafters-io/http-server-starter-go/app/http"
)

func LoggingMiddleware(next handler) handler {
	return func(req http.HTTPReq) http.HTTPResponse {
		fmt.Println("Receiving call on ", req.Url.Original)
		resp := next(req)
		fmt.Println("Received call on ", req.Url.Original)
		return resp
	}
}

func TimingMiddleware(next handler) handler {
	return func(req http.HTTPReq) http.HTTPResponse {
		start := time.Now()
		resp := next(req)
		duration := time.Since(start)
		fmt.Printf("%s %s - %d (%v)\n", req.Method, req.Url.Original, resp.StatusCode, duration)
		return resp
	}
}
