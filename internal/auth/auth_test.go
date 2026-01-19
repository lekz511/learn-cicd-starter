package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name    string
		headers http.Header
		want    string
		wantErr bool
	}{
		{
			name:    "No Auth Header",
			headers: http.Header{},
			want:    "",
			wantErr: false,
		},
		{
			name:    "Malformed Header",
			headers: http.Header{"Authorization": {"Bearer"}},
			want:    "",
			wantErr: true,
		},
		{
			name:    "Valid API Key",
			headers: http.Header{"Authorization": {"ApiKey 12345"}},
			want:    "12345",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
