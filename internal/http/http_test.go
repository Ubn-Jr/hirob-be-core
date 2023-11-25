package http

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSetupHTTPServer(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		url            string
		requestBody    string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Test POST /api/movement",
			method:         "POST",
			url:            "/api/movement",
			requestBody:    "test data",
			expectedStatus: http.StatusOK,
			expectedBody:   "Movement request sent to MQTT",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a new HTTP request with the given method, URL, and request body
			req, err := http.NewRequest(tt.method, tt.url, strings.NewReader(tt.requestBody))
			if err != nil {
				t.Fatal(err)
			}

			// Create a new response recorder to record the HTTP response
			rec := httptest.NewRecorder()

			// Create a new Gin engine and set it as the HTTP handler
			router := SetupHTTPServer()
			router.ServeHTTP(rec, req)

			// Check if the actual status code matches the expected status code
			if rec.Code != tt.expectedStatus {
				t.Errorf("unexpected status code: got %v, want %v", rec.Code, tt.expectedStatus)
			}

			// Check if the actual response body matches the expected response body
			if rec.Body.String() != tt.expectedBody {
				t.Errorf("unexpected response body: got %v, want %v", rec.Body.String(), tt.expectedBody)
			}
		})
	}
}
