package web

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// SetCookie is an HTTP handler that creates and sets a cookie named "X-CAMELIA-Name".
// The cookie's value is taken from the "name" query parameter in the request URL.
// The cookie is set with a path of "/" (available to all paths on the domain).
// The handler writes "Success Create Cookie" to the response body and sets the cookie in the response headers.
func SetCookie(w http.ResponseWriter, r *http.Request) {
	// Create a new HTTP cookie
	cookie := new(http.Cookie)
	cookie.Name = "X-CAMELIA-Name" // Set the cookie name
	cookie.Value = r.URL.Query().Get("name") // Set the cookie value from the "name" query parameter
	cookie.Path = "/" // Set the cookie path to root, making it accessible across the site

	// Add the cookie to the response headers
	http.SetCookie(w, cookie)
	// Write a success message to the response body
	fmt.Fprint(w, "Success Create Cookie")
}

// GetCookie is an HTTP handler that retrieves the "X-CAMELIA-Name" cookie from the request.
// If the cookie exists, it writes the cookie's value and details to the response.
// If the cookie is not found, it responds with "No Cookie".
// Note: fmt.Fprintln adds a newline character (\n) to the output.
func GetCookie(w http.ResponseWriter, r *http.Request) {
	// Attempt to retrieve the "X-CAMELIA-Name" cookie from the request
	cookie, err := r.Cookie("X-CAMELIA-Name")
	if err != nil {
		// If the cookie is not found or an error occurs, respond with "No Cookie"
		fmt.Fprint(w, "No Cookie")
	} else {
		// If the cookie exists, write its value and details to the response
		fmt.Fprintln(w, "Cookie value :", cookie.Value)
		fmt.Fprintln(w, "Cookie :", cookie)
	}
}

// TestMuxCookie demonstrates setting up an HTTP server with a multiplexer (mux) to handle cookie-related routes.
// It registers two routes:
// - "/set-cookie": Calls the SetCookie handler to create a cookie
// - "/get-cookie": Calls the GetCookie handler to retrieve the cookie
// The server listens on localhost:8000 and blocks execution until stopped.
// Note: This is not ideal for unit testing due to the blocking nature of ListenAndServe.
// Consider using httptest for non-blocking tests.
func TestMuxCookie(t *testing.T) {
	// Create a new ServeMux instance for routing
	mux := http.NewServeMux()
	// Register the SetCookie handler for the "/set-cookie" path
	mux.HandleFunc("/set-cookie", SetCookie)
	// Register the GetCookie handler for the "/get-cookie" path
	mux.HandleFunc("/get-cookie", GetCookie)

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

// TestSetCookie is a unit test for the SetCookie handler function.
// It simulates a GET request to "/?name=bram" using httptest and verifies that the handler
// sets the "X-CAMELIA-Name" cookie correctly.
// The test captures the response, extracts the cookies, and prints them for manual verification.
// Note: The test does not include assertions to verify the cookie's properties or response body.
func TestSetCookie(t *testing.T) {
	// Create a mock GET request with the query parameter "name=bram"
	request := httptest.NewRequest(http.MethodGet, "https://localhost:8080?name=bram", nil)
	// Create a ResponseRecorder to capture the handler's output (status, headers, cookies, body)
	recorder := httptest.NewRecorder()

	// Call the SetCookie handler with the mock request and recorder
	SetCookie(recorder, request)

	// Get the response from the recorder
	response := recorder.Result()
	// Extract the cookies from the response headers
	cookies := response.Cookies()

	// Print the cookies for manual verification
	fmt.Println(cookies)
}