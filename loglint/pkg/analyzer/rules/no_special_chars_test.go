package rules

import (
	"testing"
	"unicode"
)

func TestHasRepeatedPunctuation(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want bool
	}{
		{"no punctuation", "hello world", false},
		{"single dot", "hello.", false},
		{"triple dot", "hello...", true},
		{"double question", "really??", true},
		{"double exclamation", "wow!!", true},
		{"different punct", "hello!?", false},
		{"normal sentence", "hello, world.", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := hasRepeatedPunctuation(tt.s)
			if got != tt.want {
				t.Errorf("hasRepeatedPunctuation(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

func TestEmojiDetection(t *testing.T) {
	tests := []struct {
		name string
		r    rune
		want bool
	}{
		{"rocket", '🚀', true},
		{"smile", '😀', true},
		{"sun", '☀', true},
		{"letter", 'a', false},
		{"digit", '1', false},
		{"checkmark", '✓', true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := unicode.Is(emojiRanges, tt.r)
			if got != tt.want {
				t.Errorf("emoji(%q) = %v, want %v", tt.r, got, tt.want)
			}
		})
	}
}
