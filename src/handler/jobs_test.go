package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandleJobs(t *testing.T) {
	tests := []struct {
		name           string
		requestMethod  string
		requestBody    string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid POST request",
			requestMethod:  http.MethodPost,
			requestBody:    `{"tasks":[{"command":"echo task1","name":"task1","requires":[]},{"command":"echo task2","name":"task2","requires":["task1"]}]}`,
			expectedStatus: http.StatusOK,
			expectedBody:   "#!/usr/bin/env bash\necho task1\necho task2\n",
		},
		{
			name:           "Invalid request method",
			requestMethod:  http.MethodGet,
			requestBody:    "",
			expectedStatus: http.StatusMethodNotAllowed,
			expectedBody:   "Only POST requests are allowed\n",
		},
		{
			name:           "Invalid JSON in request body",
			requestMethod:  http.MethodPost,
			requestBody:    `{"tasks":[{"command":"echo task1","name":"task1","requires":[]},{"command":"echo task2","name":"task2","requires":["task1"]`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Failed to parse request body: unexpected EOF\n",
		},
		{
			name:           "Cycle detected",
			requestMethod:  http.MethodPost,
			requestBody:    `{"tasks":[{"command":"echo task1","name":"task1","requires":["task2"]},{"command":"echo task2","name":"task2","requires":["task1"]}]}`,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "An error occured while sorting: cycle detected\n",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(test.requestMethod, "/jobs", strings.NewReader(test.requestBody))
			w := httptest.NewRecorder()

			HandleJobs(w, req)

			if w.Code != test.expectedStatus {
				t.Errorf("Expected status code %d but got %d", test.expectedStatus, w.Code)
			}

			if w.Body.String() != test.expectedBody {
				t.Errorf("Expected body %q but got %q", test.expectedBody, w.Body.String())
			}
		})
	}
}
