package web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// ResponseCode is an HTTP handler that processes a GET request with a "name" query parameter.
// It checks if the "name" parameter is empty:
// - If empty, it sets the HTTP status code to 400 Bad Request and responds with "Name is empty."
// - If not empty, it sets the HTTP status code to 200 OK and responds with "Hi <name>".
// The handler uses w.WriteHeader to explicitly set the status code before writing the response.
func ResponseCode(w http.ResponseWriter, r *http.Request) {
	// Extract the "name" query parameter from the URL
	name := r.URL.Query().Get("name")

	if name == "" {
		// If the name parameter is empty, return a 400 Bad Request status
		w.WriteHeader(http.StatusBadRequest)
		// Alternative: w.WriteHeader(400) // Same as http.StatusBadRequest
		fmt.Fprint(w, "Name is empty.")
	} else {
		// If the name parameter is provided, return a 200 OK status
		w.WriteHeader(http.StatusOK)
		// Alternative: w.WriteHeader(200) // Same as http.StatusOK
		fmt.Fprint(w, "Hi ", name)
	}
}

// TestResponseCode is a unit test for the ResponseCode handler function.
// It simulates a GET request to "/?name=bram" using httptest and verifies the response.
// The test captures the response status code and body, then prints them for manual verification.
// Note: The commented-out request without the "name" parameter suggests testing the empty name case,
// but it is not currently executed. Consider adding additional test cases for robustness.
func TestResponseCode(t *testing.T) {
	// Create a mock GET request with the query parameter "name=bram"
	request := httptest.NewRequest("GET", "https://localhost:8080?name=bram", nil)
	// Alternative (commented): 
	// request := httptest.NewRequest("GET", "https://localhost:8080", nil)
	// This would test the case where the "name" parameter is missing

	// Create a ResponseRecorder to capture the handler's output (status, headers, body)
	recorder := httptest.NewRecorder()

	// Call the ResponseCode handler with the mock request and recorder
	ResponseCode(recorder, request)

	// Get the response from the recorder
	response := recorder.Result()

	// Read the response body as a byte slice
	// Note: Error handling for io.ReadAll is ignored (_); in production tests, check for errors
	body, _ := io.ReadAll(response.Body)

	// Print the response status code, status string, and body for verification
	fmt.Println("Code :", response.StatusCode) // e.g., 200
	fmt.Println("Status Code :", response.Status) // e.g., "200 OK"
	fmt.Println(string(body)) // e.g., "Hi bram"
}