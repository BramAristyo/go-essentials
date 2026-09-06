package web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

// ServeFile is an HTTP handler that serves a specific file ("./resources/index.css") for all requests.
// It uses http.ServeFile to directly serve the file from the local filesystem.
// The handler responds with the file's contents and appropriate headers (e.g., Content-Type, Content-Length).
// Note: This handler serves the same file regardless of the request URL, which may not be ideal for a production file server.
func ServeFile(w http.ResponseWriter, r *http.Request) {
	// Serve the "./resources/index.css" file directly
	// http.ServeFile sets appropriate headers (e.g., Content-Type: text/css) and streams the file content
	http.ServeFile(w, r, "./resources/index.css")
}

// TestServeFile demonstrates setting up an HTTP server to test the ServeFile handler.
// It configures a server to listen on localhost:8000 and uses ServeFile as the handler for all requests.
// The server blocks execution until stopped, making it unsuitable for unit testing.
// Note: For unit testing, consider using httptest to simulate requests without starting a real server.
func TestServeFile(t *testing.T) {
	// Initialize an HTTP server with the ServeFile handler and address
	server := http.Server{
		Addr:    "localhost:8000",
		Handler: http.HandlerFunc(ServeFile), // Convert ServeFile to http.HandlerFunc
	}

	// Start the server and listen for incoming requests
	// ListenAndServe blocks execution until the server stops
	err := server.ListenAndServe()
	if err != nil {
		// Panic if the server fails to start (e.g., port already in use)
		// Note: In production, handle errors gracefully (e.g., log or return errors)
		panic(err)
	}
}

//go:embed resources/index.css
// resourcesCss is a string containing the contents of the "resources/index.css" file.
// The go:embed directive embeds the file's contents into the binary at compile time.
var resourcesCss string

// ServeFileEmbed is an HTTP handler that serves the embedded contents of "resources/index.css".
// It writes the contents of the embedded file (stored in resourcesCss) directly to the response.
// Note: Unlike http.ServeFile, this does not automatically set headers like Content-Type or Content-Length.
func ServeFileEmbed(w http.ResponseWriter, r *http.Request) {
	// Write the embedded CSS file contents to the response
	// Note: fmt.Fprint does not add a newline and does not set Content-Type (defaults to text/plain)
	fmt.Fprint(w, resourcesCss)
}

// TestServeFileEmbed demonstrates setting up an HTTP server to test the ServeFileEmbed handler.
// It configures a server to listen on localhost:8000 and uses ServeFileEmbed as the handler for all requests.
// The server blocks execution until stopped, making it unsuitable for unit testing.
// Note: For unit testing, consider using httptest to simulate requests without starting a real server.
func TestServeFileEmbed(t *testing.T) {
	// Initialize an HTTP server with the ServeFileEmbed handler and address
	server := http.Server{
		Addr:    "localhost:8000",
		Handler: http.HandlerFunc(ServeFileEmbed), // Convert ServeFileEmbed to http.HandlerFunc
	}

	// Start the server and listen for incoming requests
	// ListenAndServe blocks execution until the server stops
	err := server.ListenAndServe()
	if err != nil {
		// Panic if the server fails to start (e.g., port already in use)
		// Note: In production, handle errors gracefully (e.g., log or return errors)
		panic(err)
	}
}