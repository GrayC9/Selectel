package rules

import (
	"go/ast"
	"go/token"
	"strings"
	"unicode"
	"unicode/utf8"

	"golang.org/x/tools/go/analysis"
)

func CheckLowercase(pass *analysis.Pass, msg string, node ast.Expr) {
	if msg == "" {
		return
	}
	r, _ := utf8.DecodeRuneInString(msg)
	if r == utf8.RuneError {
		return
	}
	if !unicode.IsUpper(r) {
		return
	}

	diag := analysis.Diagnostic{
		Pos:     node.Pos(),
		Message: "loglint: log message must start with a lowercase letter",
	}

	if lit, ok := node.(*ast.BasicLit); ok && lit.Kind == token.STRING {
		lower := strings.ToLower(string(r))
		start := lit.Pos() + 1
		end := start + token.Pos(utf8.RuneLen(r))
		diag.SuggestedFixes = []analysis.SuggestedFix{
			{
				Message: "lowercase first letter",
				TextEdits: []analysis.TextEdit{
					{Pos: start, End: end, NewText: []byte(lower)},
				},
			},
		}
	}

	pass.Report(diag)
}
