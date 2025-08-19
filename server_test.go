package learn_pzn_web

import (
	"fmt"
	"net/http"
	"testing"
)

// SERVER
// TestServer demonstrates a minimal HTTP server setup.
// It creates a server that listens on localhost:8000 without a specific handler.
// Note: Since no handler is defined, the server uses the default HTTP multiplexer,
// which responds with a 404 for all requests. This is a basic example and may not
// be suitable for production. Consider adding a handler for specific routes.
func TestServer(t *testing.T) {
	// Initialize an HTTP server with the address set to localhost:8000
	server := http.Server{
		Addr: "localhost:8000",
	}

	// Start the server and listen for incoming requests
	// ListenAndServe blocks execution until the server is stopped or an error occurs
	err := server.ListenAndServe()
	if err != nil {
		// Panic if the server fails to start (e.g., port already in use)
		panic(err)
	}
}

// BASIC HANDLER & SERVER
// TestHandler demonstrates a basic HTTP server with a custom handler function.
// The handler responds with "Hello world!" for all incoming requests, regardless
// of the URL path or HTTP method. The server listens on localhost:8000.
// Note: ListenAndServe() blocks execution, so this is not ideal for unit tests.
// For testing, consider using the httptest package to create a non-blocking server.
func TestHandler(t *testing.T) {
	// Define a handler function that writes "Hello world!" to the response
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		// Logic for handling HTTP requests goes here
		// Currently, it simply writes "Hello world!" to the response body
		fmt.Fprint(writer, "Hello world!")
	}

	// Initialize an HTTP server with the specified address and handler
	server := http.Server{
		Addr:    "localhost:8000",
		Handler: handler, // Assign the custom handler to process requests
	}

	// Start the server and listen for incoming requests
	err := server.ListenAndServe()
	if err != nil {
		// Panic if the server fails to start (e.g., address already in use)
		panic(err)
	}
}

// JSONHandler is a handler function that responds with a JSON object.
// It sets the Content-Type header to "application/json", sets the HTTP status
// code to 201 (Created), and sends a simple JSON response: {"ok":true}.
// This can be used for APIs that need to return JSON data.
func JSONHandler(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to indicate the response is JSON
	w.Header().Set("Content-Type", "application/json")
	// Set the HTTP status code to 201 (Created)
	w.WriteHeader(http.StatusCreated)
	// Write the JSON response body
	w.Write([]byte(`{"ok":true}`))
}

// TestJSONHandler demonstrates an HTTP server with a JSON handler.
// It uses the JSONHandler function to respond to requests with a JSON object.
// The server listens on localhost:8000 and processes all requests with JSONHandler.
// Note: Similar to TestHandler, ListenAndServe() blocks execution, so consider
// using httptest for non-blocking test scenarios.
func TestJSONHandler(t *testing.T) {
	// Assign the JSONHandler function to handle HTTP requests
	var handler http.HandlerFunc = JSONHandler

	// Initialize an HTTP server with the specified address and JSON handler
	server := http.Server{
		Addr:    "localhost:8000",
		Handler: handler, // Assign JSONHandler to process requests
	}

	// Start the server and listen for incoming requests
	err := server.ListenAndServe()
	if err != nil {
		// Panic if the server fails to start (e.g., port conflict)
		panic(err)
	}
}