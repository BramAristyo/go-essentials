package learn_pzn_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// QUERY PARAMETER
// SayHello is an HTTP handler function that processes a single query parameter "name".
// It checks if the "name" query parameter equals "bram" and responds with "Hello Bram!".
// For any other value (or if the parameter is missing), it responds with "Unknown".
// The response is written to the ResponseWriter with a newline character.
func SayHello(w http.ResponseWriter, r *http.Request) {
	// Extract the "name" query parameter from the URL
	name := r.URL.Query().Get("name")

	// Check if the name is "bram" and respond accordingly
	if name == "bram" {
		fmt.Fprintln(w, "Hello Bram!")
	} else {
		fmt.Fprintln(w, "Unknown")
	}
}

// TestSayHello tests the SayHello handler function using a mock HTTP request.
// It simulates a GET request to "/say-hello" with the query parameter "name=bram".
// The httptest package is used to create a request and capture the response without
// starting a real server, making it suitable for unit testing.
// The response body is read and printed to the console for verification.
// Note: Error handling for io.ReadAll is ignored here (_); in production tests,
// consider checking for errors to ensure robustness.
func TestSayHello(t *testing.T) {
	// Create a mock GET request with the query parameter "name=bram"
	request := httptest.NewRequest("GET", "http://localhost:8080/say-hello?name=bram", nil)
	// Create a response recorder to capture the handler's output
	recorder := httptest.NewRecorder()

	// Call the SayHello handler with the mock request and recorder
	SayHello(recorder, request)

	// Get the response from the recorder
	response := recorder.Result()
	// Read the response body
	body, _ := io.ReadAll(response.Body)

	// Print the response body to verify the output (e.g., "Hello Bram!")
	fmt.Println(string(body))
}

// MULTIPLE QUERY PARAMETER
// MultipleParameter is an HTTP handler function that processes two query parameters:
// "first_name" and "last_name". It extracts both parameters from the URL and writes
// them to the response, separated by a space, followed by a newline.
// If a parameter is missing, an empty string is returned for that parameter.
func MultipleParameter(w http.ResponseWriter, r *http.Request) {
	// Extract "first_name" and "last_name" query parameters from the URL
	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")

	// Write the first and last names to the response, separated by a space
	fmt.Fprintln(w, firstName, lastName)
}

// TestMultipleParameter tests the MultipleParameter handler function with multiple query parameters.
// It simulates a GET request to "/say-hello" with query parameters "first_name=bram" and
// "last_name=aristyo". The httptest package is used to create a mock request and capture
// the response. The response body is read and printed to the console for verification.
// Note: Like TestSayHello, error handling for io.ReadAll is ignored here (_).
// For production tests, consider adding error checking and assertions to verify the response.
func TestMultipleParameter(t *testing.T) {
	// Create a mock GET request with query parameters "first_name=bram" and "last_name=aristyo"
	request := httptest.NewRequest("GET", "http://localhost:8080/say-hello?first_name=bram&last_name=aristyo", nil)
	// Create a response recorder to capture the handler's output
	recorder := httptest.NewRecorder()

	// Call the MultipleParameter handler with the mock request and recorder
	MultipleParameter(recorder, request)

	// Get the response from the recorder
	response := recorder.Result()
	// Read the response body
	body, _ := io.ReadAll(response.Body)

	// Print the response body to verify the output (e.g., "bram aristyo")
	fmt.Println(string(body))
}