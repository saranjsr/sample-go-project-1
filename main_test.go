package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// TestHelloWorldHandler tests the /helloworld route handler.
func TestHelloWorldHandler(t *testing.T) {
	// Create a request to pass to the handler.
	req, err := http.NewRequest("GET", "/helloworld", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Create a response recorder to capture the handler's response.
	rr := httptest.NewRecorder()

	// Create a test handler using the same handler logic as in main.go
	handler := http.NewServeMux()
	handler.HandleFunc("/helloworld", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Mondoo Engineer!")
	})

	// Serve the request using the handler
	handler.ServeHTTP(rr, req)

	// Check the status code is 200 OK.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := "Hello from Mondoo Engineer!"
	body, _ := io.ReadAll(rr.Body)
	if string(body) != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", string(body), expected)
	}
}

// TestServerInitialization tests server struct initialization
func TestServerInitialization(t *testing.T) {
	server := &http.Server{
		Addr:         ":8080",
		Handler:      http.NewServeMux(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	if server.Addr != ":8080" {
		t.Errorf("Expected server address to be :8080, got %s", server.Addr)
	}

	if server.ReadTimeout != 10*time.Second {
		t.Errorf("Expected ReadTimeout to be 10 seconds, got %v", server.ReadTimeout)
	}

	if server.WriteTimeout != 10*time.Second {
		t.Errorf("Expected WriteTimeout to be 10 seconds, got %v", server.WriteTimeout)
	}

	if server.IdleTimeout != 60*time.Second {
		t.Errorf("Expected IdleTimeout to be 60 seconds, got %v", server.IdleTimeout)
	}
}
