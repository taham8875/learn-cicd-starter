package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headers       map[string]string
		expectedKey   string
		expectedError error
	}{
		{
			name: "valid API key header",
			headers: map[string]string{
				"Authorization": "ApiKey test-api-key-123",
			},
			expectedKey:   "test-api-key-123",
			expectedError: nil,
		},
		{
			name: "valid API key with spaces in key",
			headers: map[string]string{
				"Authorization": "ApiKey test-api-key-with-spaces",
			},
			expectedKey:   "test-api-key-with-spaces",
			expectedError: nil,
		},
		{
			name: "no authorization header",
			headers: map[string]string{
				"Content-Type": "application/json",
			},
			expectedKey:   "",
			expectedError: ErrNoAuthHeaderIncluded,
		},
		{
			name: "empty authorization header",
			headers: map[string]string{
				"Authorization": "",
			},
			expectedKey:   "",
			expectedError: ErrNoAuthHeaderIncluded,
		},
		{
			name: "malformed authorization header - wrong prefix",
			headers: map[string]string{
				"Authorization": "Bearer test-token",
			},
			expectedKey:   "",
			expectedError: errors.New("malformed authorization header"),
		},
		{
			name: "malformed authorization header - no space",
			headers: map[string]string{
				"Authorization": "ApiKey",
			},
			expectedKey:   "",
			expectedError: errors.New("malformed authorization header"),
		},
		{
			name: "malformed authorization header - only one part",
			headers: map[string]string{
				"Authorization": "ApiKey",
			},
			expectedKey:   "",
			expectedError: errors.New("malformed authorization header"),
		},
		{
			name: "malformed authorization header - multiple spaces",
			headers: map[string]string{
				"Authorization": "ApiKey  test-key  extra",
			},
			expectedKey:   "",
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create http.Header from map
			headers := make(http.Header)
			for key, value := range tt.headers {
				headers.Set(key, value)
			}

			// Call the function
			key, err := GetAPIKey(headers)

			// Check the API key
			if key != tt.expectedKey {
				t.Errorf("GetAPIKey() key = %v, want %v", key, tt.expectedKey)
			}

			// Check the error
			if tt.expectedError == nil {
				if err != nil {
					t.Errorf("GetAPIKey() error = %v, want nil", err)
				}
			} else {
				if err == nil {
					t.Errorf("GetAPIKey() error = nil, want %v", tt.expectedError)
				} else if err.Error() != tt.expectedError.Error() {
					t.Errorf("GetAPIKey() error = %v, want %v", err, tt.expectedError)
				}
			}
		})
	}
}
