package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"product-api/microservices/handlers"
	"syscall"
	"time"
)

func main() {
	// Creates a log_object with a prefix , flag and destination(stdout)
	log_object := log.New(os.Stdout, "Product-API:", log.LstdFlags)

	// Custom HANDLER
	hh := handlers.NewHello(log_object) // This will return an Hello instance with given logger ptr
	gh := handlers.NewGoodbye(log_object)

	// CUSTOM ServeMux
	mux := http.NewServeMux()
	// Manually registering	 the custom handlers to the diff paths
	mux.Handle("/", hh)
	mux.Handle("/goodbye", gh)

	server := &http.Server{
		Addr:         ":8090",
		Handler:      mux,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	// start the server in the background for non-blocking the main program
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log_object.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}() // () - calls the anonymous go routine function immediately

	// Create buffered channel to receive signals
	sigChan := make(chan os.Signal, 1)
	// Register to receive SIGINT and SIGTERM
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	log_object.Println("Waiting for any SIG...")
	// Block until signal received
	sig := <-sigChan
	log_object.Println("Got Signals : ", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)

}
