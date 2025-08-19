package learn_pzn_web

import (
	"embed"
	"io/fs"
	"net/http"
	"testing"
)

// TestFileServer demonstrates serving static files from a local directory using Go's http.FileServer.
// It sets up an HTTP server that serves files from the "./resources" directory, accessible under the "/static/" URL path.
// The http.StripPrefix function is used to remove the "/static" prefix from the request URL, allowing clean access to files.
// For example, a request to "/static/image.jpg" serves "./resources/image.jpg".
// The server runs on localhost:8000 and blocks execution until stopped.
// Note: This is not ideal for unit testing due to the blocking nature of ListenAndServe; consider using httptest for testing.
func TestFileServer(t *testing.T) {
	// Create an http.Dir to serve files from the local "./resources" directory
	dir := http.Dir("./resources")
	// Create a file server to handle requests for static files in the directory
	fileServer := http.FileServer(dir)

	// Create a new ServeMux instance for routing
	mux := http.NewServeMux()
	// Register the file server to handle requests under "/static/"
	// http.StripPrefix removes "/static" from the URL to map to the correct file path
	// Example: "/static/image.jpg" maps to "./resources/image.jpg"
	// Note: The commented-out line (mux.Handle("/static/", fileServer)) would serve files with "/static" in the path,
	// e.g., "/static/resources/image.jpg", which is usually not desired
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

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
		// Note: In production, handle errors gracefully (e.g., log or return errors)
		panic(err)
	}
}

//go:embed resources/*
// resources is an embedded filesystem containing all files under the "resources" directory.
// The go:embed directive includes these files in the binary at compile time, eliminating the need
// for a physical directory at runtime.
var resources embed.FS

// TestFileServerGolangEmbed demonstrates serving static files from an embedded filesystem using Go's embed.FS.
// It uses the go:embed directive to include files from the "resources" directory in the binary.
// The fs.Sub function creates a sub-filesystem rooted at "resources" to serve files cleanly.
// The file server is configured to serve files under the "/static/" URL path, with http.StripPrefix
// removing the "/static" prefix for proper file mapping.
// The server runs on localhost:8000 and blocks execution until stopped.
// Note: Like TestFileServer, this is not ideal for unit testing due to blocking; use httptest for testing.
func TestFileServerGolangEmbed(t *testing.T) {
	// Create a sub-filesystem rooted at "resources" from the embedded FS
	// Note: Error from fs.Sub is ignored (_); in production, check for errors
	dir, _ := fs.Sub(resources, "resources")
	// Create a file server to handle requests for files in the embedded filesystem
	fileServer := http.FileServer(http.FS(dir))

	// Create a new ServeMux instance for routing
	mux := http.NewServeMux()
	// Register the file server to handle requests under "/static/"
	// http.StripPrefix removes "/static" from the URL to map to the correct file path
	// Example: "/static/image.jpg" serves the embedded "resources/image.jpg"
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

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
		// Note: In production, handle errors gracefully (e.g., log or return errors)
		panic(err)
	}
}