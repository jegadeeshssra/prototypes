package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"product-api/handlers"
	"syscall"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

func main() {
	log_object := log.New(os.Stdout, "Product-API:v1 - ", log.LstdFlags)

	productHandler := handlers.NewProductsHandler(log_object)

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", productHandler.GetProducts)
	getRouter.HandleFunc("/products/{id:[0-9]+}", productHandler.GetSingleProduct)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products/{id:[0-9]+}", productHandler.UpdateProduct)
	putRouter.Use(productHandler.Middleware)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products/", productHandler.AddProduct)
	postRouter.Use(productHandler.Middleware)

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/products/{id:[0-9]+}", productHandler.DeleteProduct)

	// for serving documentation using Redoc
	options := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	docsHandler := middleware.Redoc(options, nil)
	getRouter.Handle("/docs", docsHandler)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	server := &http.Server{
		Addr:         ":8090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log_object.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// Create buffered channel to receive signals
	sigChan := make(chan os.Signal, 1)
	// Register to receive SIGINT and SIGTERM
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	log_object.Println("Waiting for any SIG...")
	// Block until signal received
	sig := <-sigChan
	log_object.Println("Got Signals : ", sig)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)
}
