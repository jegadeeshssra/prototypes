package server

import (
	"fmt"
	"strconv"
	"strings"
)

//var DirPath string = "C:\\Users\\018046\\OneDrive - Sify Technologies Limited\\go\\prototypes\\codecrafters-http-server-go\\files\\"

type HTTPReq struct {
	Headers map[string]string
	Url     URL
	Method  string
	Body    []byte
}

type URL struct {
	Original    string
	Parameters  map[string]string
	QueryParams map[string]string
}

func (request HTTPReq) readUrlParams(method string, path string, routes *[]Route) func(HTTPReq) HTTPResponse {
	var matchRouteFunc func(HTTPReq) HTTPResponse
	matchRouteFunc = nil

ROUTELOOP:
	for _, route := range *routes {

		if route.Method != method {
			continue
		}

		uriParts := strings.Split(path, "/")
		routePathParts := strings.Split(route.Path, "/")

		if len(uriParts) != len(routePathParts) {
			continue
		}

		parameters := make(map[string]string, 0)
		for i, part := range routePathParts {
			if strings.HasPrefix(routePathParts[i], "{") && strings.HasSuffix(routePathParts[i], "}") {
				parameters[part[1:len(part)-1]] = uriParts[i]
				continue
			}
			if len(uriParts[i]) == len(part) {
				continue
			}
			continue ROUTELOOP
		}

		request.Url = URL{
			Original:   path,
			Parameters: parameters,
		}

		matchRouteFunc = route.Function
	}
	return matchRouteFunc
}

func (request HTTPReq) ReadAndProcessRequest(requeststr string, routes *[]Route) (func(HTTPReq) HTTPResponse, bool) {
	reqSplit := strings.Split(requeststr, "\r\n\r\n")
	lines := strings.Split(reqSplit[0], "\r\n")
	parts := strings.Split(lines[0], " ")
	// METHOD
	method := strings.TrimSpace(parts[0])
	// URL
	path := strings.TrimSpace(parts[1])
	// HEADERS
	headers := make(map[string]string)
	for i := 1; i < len(lines); i++ {
		headerParts := strings.Split(lines[i], ":")
		if len(headerParts) >= 2 {
			headers[headerParts[0]] = strings.Join(headerParts[1:], "")
		}
	}
	// Content Length to extract the BODY CONTENT
	contentLength, err := strconv.Atoi(headers["Content-Length"])
	if err != nil {
		fmt.Println("Could not convert content length to int, ignoring body")
		contentLength = 0
	}

	// Filling in the request struct
	request.Headers = headers
	request.Body = []byte(reqSplit[1][:contentLength])
	request.Method = method

	var keepAlive bool
	connection := request.Headers["Connection"]
	if connection != "" && connection == "close" {
		// client wants to close the connection after this request/response
		keepAlive = false
	}

	// retrieve the route matched and fill in the parameters
	matchedRouteFunc := request.readUrlParams(method, path, routes)
	if matchedRouteFunc == nil {
		keepAlive = false
		return nil, keepAlive
	}

	return matchedRouteFunc, keepAlive
}
