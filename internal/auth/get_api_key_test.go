package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		authorization string
		wantKey       string
		wantErr       string
	}{
		{
			name:          "valid API key",
			authorization: "ApiKey abc123",
			wantKey:       "abc123",
		},
		{
			name:    "missing authorization header",
			wantErr: "no authorization header included",
		},
		{
			name:          "wrong authorization scheme",
			authorization: "Bearer abc123",
			wantErr:       "malformed authorization header",
		},
		{
			name:          "missing API key",
			authorization: "ApiKey",
			wantErr:       "malformed authorization header",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			headers := make(http.Header)

			if tt.authorization != "" {
				headers.Set("Authorization", tt.authorization)
			}

			gotKey, err := GetAPIKey(headers)

			if tt.wantErr != "" {
				if err == nil {
					t.Fatalf("expected error %q, got nil", tt.wantErr)
				}

				if err.Error() != tt.wantErr {
					t.Errorf("expected error %q, got %q", tt.wantErr, err.Error())
				}

				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if gotKey != tt.wantKey {
				t.Errorf("expected API key %q, got %q", tt.wantKey, gotKey)
			}
		})
	}
}
