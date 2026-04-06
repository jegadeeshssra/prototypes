package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	g.l.Println("Handling Goodbye Request")
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(res, "/goodbye endpoint failed", http.StatusBadRequest)
	}
	fmt.Fprintf(res, "\nReceived DATA : %s", data)
}
