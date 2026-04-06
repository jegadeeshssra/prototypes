package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		log.Println("This is the reply for the path '/' ")

		data, err := ioutil.ReadAll(req.Body)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte("OOPS"))
			// (2 lines above) OR below
			// http.Error(res,"OOPS",http.StatusBadRequest)
			return
		}
		//res.WriteHeader(http.StatusAccepted)
		//res.Write([]byte("received"))
		// or
		fmt.Fprintf(res, "\nReceived DATA : %s", data) //automatic formatting
		//log.Printf("Type : %T", data) // Type : []uint8
		//log.Printf("DATA : %s", data)
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		log.Println(("This is home"))
	})

	http.ListenAndServe(":8080", nil)
}
