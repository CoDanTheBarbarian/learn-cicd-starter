package auth

import (
	"net/http"
	"testing"
)

func TestAuth(t *testing.T) {
	tests := []struct {
		name    string
		header  http.Header
		wantKey string
		wantErr bool
	}{
		{
			name:    "valid API key",
			header:  http.Header{"Authorization": []string{"ApiKey 1234567890"}},
			wantKey: "1234567890",
			wantErr: false,
		},
		{
			name:    "missing auth header",
			header:  http.Header{},
			wantKey: "",
			wantErr: true,
		},
		{
			name:    "malformed auth header",
			header:  http.Header{"Authorization": []string{"1234567890"}},
			wantKey: "",
			wantErr: true,
		},
		{
			name:    "empty auth header value",
			header:  http.Header{"Authorization": []string{""}},
			wantKey: "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotKey, err := GetAPIKey(tt.header)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotKey != tt.wantKey {
				t.Errorf("GetAPIKey() gotKey = %v, want %v", gotKey, tt.wantKey)
			}
		})
	}
}
