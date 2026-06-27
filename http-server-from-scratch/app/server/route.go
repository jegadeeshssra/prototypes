package server

type Route struct {
	Function func(request HTTPReq) HTTPResponse
	Method   string
	Path     string
}

var Routes = []Route{}

// This will take the routes array and append the new route within same loc of routes array.
func AddRoute(routes *[]Route, path string, function func(HTTPReq) HTTPResponse, method string) {
	*routes = append(*routes, Route{
		Function: function,
		Path:     path,
		Method:   method,
	})
}
