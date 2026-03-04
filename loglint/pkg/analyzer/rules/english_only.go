package rules

import (
	"go/ast"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

func CheckEnglish(pass *analysis.Pass, msg string, node ast.Expr) {
	if msg == "" {
		return
	}
	for _, r := range msg {
		if isAllowedEnglishRune(r) {
			continue
		}
		pass.Report(analysis.Diagnostic{
			Pos:     node.Pos(),
			Message: "loglint: log message must contain only English characters",
		})
		return
	}
}

func isAllowedEnglishRune(r rune) bool {
	if unicode.IsLetter(r) {
		return unicode.In(r, unicode.Latin)
	}
	if unicode.IsDigit(r) {
		return true
	}
	switch r {
	case ' ', ',', '-', '.', ':', '/', '\'', '"', '(', ')', '[', ']',
		'{', '}', '_', '=', '+', '#', '@', '&', '*', ';', '!', '?',
		'\n', '\t', '\r', '%', '<', '>', '|', '\\', '^', '~', '`', '$':
		return true
	}
	return false
}
