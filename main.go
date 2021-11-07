package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"rest-files-app/handlers"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	l := log.New(os.Stdout, "product-images-api", log.LstdFlags)

	// create a multiplexer
	m := mux.NewRouter()

	// create subrouters and handlers
	create := m.Methods(http.MethodPost).Subrouter()
	create.Handle("/", handlers.NewCreateImage(l))

	read := m.Methods(http.MethodGet).Subrouter()
	read.Handle("/", handlers.NewReadImages(l))

	update := m.Methods(http.MethodPut).Subrouter()
	update.Handle("/{id:[0-9]+}", handlers.NewUpdateImage(l))

	delete := m.Methods(http.MethodDelete).Subrouter()
	delete.Handle("/{id:[0-9]+}", handlers.NewDeleteImage(l))

	// create a server
	s := http.Server{
		Addr:         ":8080", // add actual env address
		Handler:      m,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		ErrorLog:     l,
	}

	// start the server
	go func() {
		l.Println("Starting server on port :8080")
		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// block until a signal is received
	sig := <-c
	l.Println("Signal:", sig)

	// gracefully shutdown the server, wait max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
