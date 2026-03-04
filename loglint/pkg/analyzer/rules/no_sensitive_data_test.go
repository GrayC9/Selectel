package rules

import (
	"testing"
)

func TestNormalise(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"api_key", "apikey"},
		{"API_KEY", "apikey"},
		{"Password", "password"},
		{"private_key", "privatekey"},
		{"hello", "hello"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := normalise(tt.input)
			if got != tt.want {
				t.Errorf("normalise(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestSensitiveKeywordDetection(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantHit bool
	}{
		{"password in string", "user password: hunter2", true},
		{"api_key in string", "using api_key to connect", true},
		{"clean message", "user login successful", false},
		{"token in string", "invalid token provided", true},
		{"auth in string", "auth failed", true},
		{"partial match", "author info", true}, // "auth" is substring of "author"
		{"secret in string", "secret value found", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			normalized := normalise(tt.input)
			hit := false
			for _, kw := range defaultSensitiveKeywords {
				if contains(normalized, kw) {
					hit = true
					break
				}
			}
			if hit != tt.wantHit {
				t.Errorf("sensitive check %q: got=%v, want=%v", tt.input, hit, tt.wantHit)
			}
		})
	}
}

func contains(s, substr string) bool {
	return len(substr) > 0 && len(s) >= len(substr) && containsStr(s, substr)
}

func containsStr(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
