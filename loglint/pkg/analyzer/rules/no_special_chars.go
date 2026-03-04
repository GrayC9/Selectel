package rules

import (
	"go/ast"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

var emojiRanges = &unicode.RangeTable{
	R16: []unicode.Range16{
		{Lo: 0x2600, Hi: 0x26FF, Stride: 1},
		{Lo: 0x2700, Hi: 0x27BF, Stride: 1},
		{Lo: 0xFE00, Hi: 0xFE0F, Stride: 1},
	},
	R32: []unicode.Range32{
		{Lo: 0x1F300, Hi: 0x1F5FF, Stride: 1},
		{Lo: 0x1F600, Hi: 0x1F64F, Stride: 1},
		{Lo: 0x1F680, Hi: 0x1F6FF, Stride: 1},
		{Lo: 0x1F900, Hi: 0x1F9FF, Stride: 1},
		{Lo: 0x1FA70, Hi: 0x1FAFF, Stride: 1},
	},
}

func CheckNoSpecialChars(pass *analysis.Pass, msg string, node ast.Expr) {
	if msg == "" {
		return
	}

	for _, r := range msg {
		if unicode.Is(emojiRanges, r) {
			pass.Report(analysis.Diagnostic{
				Pos:     node.Pos(),
				Message: "loglint: log message must not contain emoji",
			})
			return
		}
	}

	if strings.Contains(msg, "!") {
		pass.Report(analysis.Diagnostic{
			Pos:     node.Pos(),
			Message: "loglint: log message must not contain exclamation marks",
		})
		return
	}

	if hasRepeatedPunctuation(msg) {
		pass.Report(analysis.Diagnostic{
			Pos:     node.Pos(),
			Message: "loglint: log message must not contain repeated punctuation",
		})
		return
	}
}

func hasRepeatedPunctuation(s string) bool {
	prevPunct := false
	prevRune := rune(0)
	for _, r := range s {
		if unicode.IsPunct(r) {
			if prevPunct && r == prevRune {
				return true
			}
			prevPunct = true
		} else {
			prevPunct = false
		}
		prevRune = r
	}
	return false
}
