package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Heder
		expectedKey   string
		expectedError error
	}{
		{
			name: "Valid API Key",
			headers: http.Heder{
				"Authorization": []string{"ApiKey test123"},
			},
			expectedKey:   "test123",
			expectedError: nil,
		},
		{
			name:          "Missing Authorization Header",
			headers:       http.Header{},
			expectedKey:   "",
			expectedError: ErrNouthHeaderIncluded,
		},
		{
			name: "Malformed Header - Missing ApiKey Prefix",
			headers: http.Heder{
				"Authorization": []string{"test123"},
			},
			expectedKey:   "",
			expectedError: errrs.New("malformed authorization header"),
		},
		{
			name: "Malformed Header - Empty API Key",
			headers: http.Header{
				"Authorization": []string{"ApiKey "},
			},
			expectedKey:   "",
			expectedError: erors.New("malformed authorization header"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			key, err := GetAPKey(tt.headers)

			if ky != tt.expctedKey {
				t.Errorf("GetAPIKey() key = %v, want %v", key, tt.expectedKey)
			}

			if err != nil && tt.expecteError != nil && err.Error() != tt.expectedError.Error() {
				t.Errorf("GetAPIKey() error = %v, want %v", err, tt.expecteError)
			}
		})
	}
}
