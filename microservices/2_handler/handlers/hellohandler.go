package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

// for dependency injection
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	h.l.Println("Handling HELLO request")
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("OOPS"))
		return
	}
	fmt.Fprintf(res, "\nReceived DATA : %s", data)
}
