package learn_pzn_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// HTTP TESTING NOTES:
// This code demonstrates how to test HTTP handlers in Go using the httptest package.
// The httptest package allows you to simulate HTTP requests and capture responses without
// starting a real server, making it ideal for unit testing.
// Key components:
// 1. httptest.NewRequest(method, url, body) - Creates a mock HTTP request with specified method, URL, and optional body.
// 2. httptest.NewRecorder() - Creates a ResponseRecorder to capture the handler's response (status, headers, body).
// 3. Handler execution - Call the handler with the mock request and recorder to simulate request processing.
// 4. Response verification - Extract and verify the response (status code, headers, or body) from the recorder.
// Note: This approach is non-blocking and suitable for unit tests, unlike starting a real server with ListenAndServe.

// HelloHandler is a simple HTTP handler that writes "Hello World" to the response.
// It uses fmt.Fprintln, which automatically adds a newline character (\n) to the output.
// This handler responds to all requests with "Hello World\n" and sets the default HTTP status code (200 OK).
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// Write "Hello World" to the response with a newline
	fmt.Fprintln(w, "Hello World")
}

// TestHelloHandler tests the HelloHandler function using the httptest package.
// It creates a mock GET request to "/hello", captures the response, and verifies that
// the response body matches the expected output ("Hello World\n").
// The test uses t.Errorf to report failures if the response does not match expectations.
// Note: Error handling for io.ReadAll is ignored (_); in production tests, consider checking for errors.
func TestHelloHandler(t *testing.T) {
	// Step 1: Create a mock HTTP GET request to the "/hello" endpoint
	// The nil body indicates no request body (typical for GET requests)
	request := httptest.NewRequest("GET", "http://localhost:8080/hello", nil)
	// Alternative: request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)

	// Step 2: Create a ResponseRecorder to capture the handler's output (status, headers, body)
	recorder := httptest.NewRecorder()

	// Step 3: Call the HelloHandler with the mock request and recorder
	HelloHandler(recorder, request)

	// Step 4: Get the response result from the recorder
	result := recorder.Result()

	// Step 5: Read the response body as a byte slice
	body, _ := io.ReadAll(result.Body)

	// Step 6: Convert the body to a string and verify it matches the expected output
	// Note: fmt.Fprintln adds a newline, so the expected output is "Hello World\n"
	if string(body) != "Hello World\n" {
		t.Errorf("Expected 'Hello World\\n' but got %s", string(body))
	}
}

// REMEMBER:
// - httptest.NewRequest(method, url, body) creates mock requests with the specified HTTP method, URL, and optional body (use nil for no body).
// - httptest.NewRecorder() captures the handler's response without starting a real server, making it ideal for unit tests.
// - Always convert the response body ([]byte) to a string for comparison in tests.
// - fmt.Fprintln adds a newline character (\n) to the output, while fmt.Fprint does not.
// - Check the HTTP status code (result.StatusCode) if specific status codes are expected.
// - Consider adding error handling for io.ReadAll to make tests more robust.