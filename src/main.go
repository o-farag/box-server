package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	serverPort int = 8080
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlePost).Methods("POST")
	r.HandleFunc("/", handleGet).Methods("GET")

		// Create the CORS middleware handler
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // Replace "*" with your specific allowed origins
		handlers.AllowedMethods([]string{"GET", "POST"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)

	// Wrap your router with the CORS middleware
	handler := corsHandler(r)

	srv := &http.Server{
		Addr:    ":"+strconv.Itoa(serverPort),
		Handler: handler,
	}
	fmt.Printf("Listening on port %d...", serverPort)
	srv.ListenAndServe()
	
}

func handleGet(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "get\n")
}

func handlePost(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer req.Body.Close()

	// Convert the body bytes to a string
	bodyStr := string(body)

	// Print the request body
	fmt.Fprintf(w, "Hello %s\n", bodyStr)
}
