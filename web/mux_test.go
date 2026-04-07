package web

import (
	"fmt"
	"net/http"
	"testing"
)

// MUX / ROUTING
// TestMux demonstrates the use of an HTTP multiplexer (mux) to handle different routes.
// It sets up an http.ServeMux to define multiple routes with specific handler functions:
// - "/" (root): Responds with "Hello World"
// - "/hi": Responds with "Hi"
// - "/images/": Responds with "images" for any URL starting with "/images/"
// - "/images/thumbnails/": Responds with "thumbnails" for any URL starting with "/images/thumbnails/"
// Note: The ServeMux uses longest-prefix matching, so "/images/thumbnails/" takes precedence over "/images/" for matching URLs.
// The server listens on localhost:8000 and blocks execution until stopped or an error occurs.
// For unit testing, consider using httptest to avoid binding to a real port.
func TestMux(t *testing.T) {
	// Create a new ServeMux instance for routing
	mux := http.NewServeMux()

	// Register handler for the root path ("/")
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello World")
	})

	// Register handler for the "/hi" path
	mux.HandleFunc("/hi", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hi")
	})

	// Register handler for any path starting with "/images/"
	mux.HandleFunc("/images/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "images")
	})

	// Register handler for any path starting with "/images/thumbnails/"
	// Note: This takes precedence over "/images/" due to longer prefix
	mux.HandleFunc("/images/thumbnails/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "thumbnails")
	})

	// Initialize an HTTP server with the configured mux and address
	server := http.Server{
		Addr:    "localhost:8000",
		Handler: mux, // Use the custom mux for routing
	}

	// Start the server and listen for incoming requests
	// ListenAndServe blocks execution until the server stops
	err := server.ListenAndServe()
	if err != nil {
		// Panic if the server fails to start (e.g., port already in use)
		panic(err)
	}
}

// READ REQUEST
// TestRequest demonstrates how to inspect and display detailed HTTP request information.
// The handler extracts and responds with various request properties, including:
// - HTTP method (e.g., GET, POST)
// - Request URI and URL components (path, query parameters)
// - Protocol version (e.g., HTTP/1.1)
// - Host and client details (remote address, user agent)
// - Content metadata (length, content type)
// - All HTTP headers sent by the client
// This is useful for debugging or logging request details.
// The server listens on localhost:8000 and blocks execution.
// Note: For unit testing, consider using httptest to simulate requests and avoid running a real server.
func TestRequest(t *testing.T) {
	// Define a handler function to process and display request details
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		// Write request details to the response
		fmt.Fprintln(writer, "Method:", request.Method)
		fmt.Fprintln(writer, "URI:", request.RequestURI)
		fmt.Fprintln(writer, "URL Path:", request.URL.Path)
		fmt.Fprintln(writer, "URL Query:", request.URL.RawQuery)
		fmt.Fprintln(writer, "Protocol:", request.Proto)
		fmt.Fprintln(writer, "Host:", request.Host)
		fmt.Fprintln(writer, "Remote Address:", request.RemoteAddr)
		fmt.Fprintln(writer, "User Agent:", request.UserAgent())
		fmt.Fprintln(writer, "Content Length:", request.ContentLength)
		fmt.Fprintln(writer, "Content Type:", request.Header.Get("Content-Type"))

		// Print all HTTP headers
		fmt.Fprintln(writer, "\nHeaders:")
		for key, values := range request.Header {
			for _, value := range values {
				fmt.Fprintf(writer, "%s: %s\n", key, value)
			}
		}
	}

	// Initialize an HTTP server with the handler and address
	server := http.Server{
		Addr:    "localhost:8000",
		Handler: handler, // Use the custom handler to process requests
	}

	// Start the server and listen for incoming requests
	// ListenAndServe blocks execution until the server stops
	err := server.ListenAndServe()
	if err != nil {
		// Panic if the server fails to start (e.g., address already in use)
		panic(err)
	}
}