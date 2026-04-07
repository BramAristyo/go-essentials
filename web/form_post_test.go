package web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// FormPost is an HTTP handler that processes POST requests with form data.
// It parses the form data from the request body, extracts the "first_name" and "last_name"
// fields, and writes them to the response, separated by a space, with a newline.
// If form parsing fails, it panics, which is not ideal for production use.
func FormPost(w http.ResponseWriter, r *http.Request) {
	// Parse the form data from the request body (expects Content-Type: application/x-www-form-urlencoded)
	// This populates r.PostForm with key-value pairs from the form
	err := r.ParseForm()
	if err != nil {
		// Panic if form parsing fails (e.g., malformed form data)
		// Note: In production, consider returning an HTTP error (e.g., 400 Bad Request) instead of panicking
		panic(err)
	}

	// Extract "first_name" and "last_name" from the parsed form data
	firstName := r.PostForm.Get("first_name")
	lastName := r.PostForm.Get("last_name")

	// Write the first and last names to the response, separated by a space, with a newline
	// Note: fmt.Fprintln adds a newline character (\n) to the output
	fmt.Fprintln(w, firstName, lastName)
}

// TestFormPost is a unit test for the FormPost handler function.
// It simulates a POST request with URL-encoded form data containing "first_name" and "last_name".
// The test uses the httptest package to create a mock request and capture the response without
// starting a real server, making it suitable for unit testing.
// The response body is printed to the console for manual verification.
// Note: The test does not include assertions to verify the response; consider adding them for robustness.
func TestFormPost(t *testing.T) {
	// Create a request body with URL-encoded form data: "first_name=Bram&last_name=Aristyo"
	requestBody := strings.NewReader("first_name=Bram&last_name=Aristyo")

	// Create a mock HTTP POST request to the "/" endpoint with the form data
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080/", requestBody)

	// Set the Content-Type header to indicate URL-encoded form data
	// This is required for r.ParseForm() to correctly parse the form data
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Create a ResponseRecorder to capture the handler's output (status, headers, body)
	recorder := httptest.NewRecorder()

	// Call the FormPost handler with the mock request and recorder
	FormPost(recorder, request)

	// Get the response from the recorder
	response := recorder.Result()

	// Read the response body as a byte slice
	// Note: Error handling for io.ReadAll is ignored (_); in production tests, check for errors
	body, _ := io.ReadAll(response.Body)

	// Print the response body to verify the output (e.g., "Bram Aristyo\n")
	fmt.Println(string(body))
}