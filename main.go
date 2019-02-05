package main

import (
	"fmt"
	"net/http"

	configuration "github.com/golang-line-ingestor/configurations"
	handlers "github.com/golang-line-ingestor/handlers"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong!")
	})
	http.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Processing...")
		handlers.Processor()

	})
	fmt.Println(fmt.Sprintf("Server started at: %d", configuration.Port))
	http.ListenAndServe(fmt.Sprintf(":%d", configuration.Port), nil)
}
