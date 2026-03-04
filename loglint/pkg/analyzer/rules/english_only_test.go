package rules

import (
	"testing"
)

func TestIsAllowedEnglishRune(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want bool
	}{
		{"latin a", 'a', true},
		{"latin Z", 'Z', true},
		{"digit", '5', true},
		{"space", ' ', true},
		{"comma", ',', true},
		{"dot", '.', true},
		{"colon", ':', true},
		{"slash", '/', true},
		{"dash", '-', true},
		{"cyrillic", 'д', false},
		{"chinese", '中', false},
		{"arabic", 'ع', false},
		{"percent", '%', true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isAllowedEnglishRune(tt.r)
			if got != tt.want {
				t.Errorf("isAllowedEnglishRune(%q) = %v, want %v", tt.r, got, tt.want)
			}
		})
	}
}

func TestEnglishOnlyLogic(t *testing.T) {
	tests := []struct {
		name    string
		msg     string
		wantErr bool
	}{
		{"english only", "starting server on port 8080", false},
		{"cyrillic", "запуск сервера", true},
		{"mixed", "starting сервер", true},
		{"empty", "", false},
		{"punctuation", "hello, world: test-123", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hasNonEnglish := false
			for _, r := range tt.msg {
				if !isAllowedEnglishRune(r) {
					hasNonEnglish = true
					break
				}
			}
			if hasNonEnglish != tt.wantErr {
				t.Errorf("msg=%q: hasNonEnglish=%v, wantErr=%v", tt.msg, hasNonEnglish, tt.wantErr)
			}
		})
	}
}
