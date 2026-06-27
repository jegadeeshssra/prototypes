package server

import "github.com/codecrafters-io/http-server-starter-go/app/methods"

type Route struct {
	Function func(request HTTPReq) HTTPResponse
	Method   string
	Path     string
}

var routes = []Route{
	Route{
		Function: methods.DefaultPath,
		Method:   "GET",
		Path:     "/",
	},
	Route{
		Function: Echo,
		Method:   "GET",
		Path:     "/echo/{str}",
	},
	Route{
		Function: methods.DefaultPath,
		Method:   "GET",
		Path:     "/file/{filename}",
	},
}
