// package server

// // Endpoint - "/"
// func DefaultPath(request HTTPReq) HTTPResponse {
// 	return HTTPResponse{
// 		StatusCode: StatusOK,
// 	}
// }

// // Endpoint - "/echo/{str}"
// func Echo(request HTTPReq) HTTPResponse {
// 	content := request.Url.Parameters["str"]

// 	return HTTPResponse{
// 		StatusCode: StatusOK,
// 		Headers:    map[string]string{"Content-Type": "text/plain"},
// 		Body:       []byte(content),
// 	}
// }

// func UserAgentHeader(conn net.Conn, requeststr string) bool {
// 	// time.Sleep(5 * time.Second)
// 	headerContent, err := server.GetUserAgent(requeststr)
// 	if err != nil {
// 		fmt.Println(err)
// 		data := "HTTP/1.1 404 Header string Not Found\r\n\r\n"
// 		return server.WritePersistentTCPResponse(conn, data)
// 	}
// 	data := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(headerContent), headerContent)
// 	return server.WritePersistentTCPResponse(conn, data)
// }