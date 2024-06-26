package auth

import (
	"errors"
	"net/http"
	"testing"
)

// Assume ErrNoAuthHeaderIncluded is defined somewhere in yourpackage
var errNoAuthHeaderIncluded = errors.New("no authorization header included")

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		expectedKey   string
		expectedError error
	}{
		{
			name:          "Valid header",
			headers:       http.Header{"Authorization": []string{"ApiKey myapikey"}},
			expectedKey:   "myapikey",
			expectedError: errNoAuthHeaderIncluded,
		},
		{
			name:          "No header",
			headers:       http.Header{},
			expectedKey:   "",
			expectedError: errNoAuthHeaderIncluded,
		},
		{
			name:          "Malformed header - missing ApiKey",
			headers:       http.Header{"Authorization": []string{"Bearer myapikey"}},
			expectedKey:   "",
			expectedError: errors.New("malformed authorization header"),
		},
		{
			name:          "Malformed header - missing key",
			headers:       http.Header{"Authorization": []string{"ApiKey"}},
			expectedKey:   "",
			expectedError: errors.New("malformed authorization header"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPIKey(tt.headers)
			if key != tt.expectedKey {
				t.Errorf("expected key %q, got %q", tt.expectedKey, key)
			}
			if (err != nil && tt.expectedError == nil) || (err == nil && tt.expectedError != nil) || (err != nil && err.Error() != tt.expectedError.Error()) {
				t.Errorf("expected error %v, got %v", tt.expectedError, err)
			}
		})
	}
}
