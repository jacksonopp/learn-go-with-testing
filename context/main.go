package main

import (
	"fmt"
	"net/http"
)

type Store interface {
	Fetch() string
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		// Get the context from the request
		ctx := r.Context()

		// Make a single chanel that holds the data
		data := make(chan string, 1)

		// Make a go routine that stores the fetched data inside the data chanel
		go func() {
			data <- store.Fetch()
		}()

		// A select that either: (this select races the two options)
		select {
		// Stores the data (from the chanel) inside a new variable D -OR-
		case d := <-data:
			fmt.Fprint(rw, d)
			// Triggers store.Cancel if the context done happens first
		case <-ctx.Done():
			store.Cancel()
		}
	}
}

func main() {}
