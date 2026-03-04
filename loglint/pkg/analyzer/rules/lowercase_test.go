package rules

import (
	"testing"
	"unicode"
	"unicode/utf8"
)

func TestLowercaseLogic(t *testing.T) {
	tests := []struct {
		name    string
		msg     string
		wantErr bool
	}{
		{"lowercase start", "starting server", false},
		{"uppercase start", "Starting server", true},
		{"empty string", "", false},
		{"number start", "123 items", false},
		{"single char upper", "A", true},
		{"single char lower", "a", false},
		{"unicode lower", "über cool", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.msg == "" {
				if tt.wantErr {
					t.Error("empty string should not trigger error")
				}
				return
			}
			r, _ := utf8.DecodeRuneInString(tt.msg)
			hasUpper := unicode.IsUpper(r)
			if hasUpper != tt.wantErr {
				t.Errorf("msg=%q: isUpper=%v, wantErr=%v", tt.msg, hasUpper, tt.wantErr)
			}
		})
	}
}
