package rules

import (
	"go/ast"
	"go/token"
	"strconv"
	"strings"

	"golang.org/x/tools/go/analysis"

	"github.com/GrayC9/Selectel/pkg/analyzer/logcall"
)

var defaultSensitiveKeywords = []string{
	"password", "passwd", "pwd", "secret", "token",
	"apikey", "apitoken", "auth", "credential", "credentials",
	"privatekey",
}

func CheckNoSensitiveData(pass *analysis.Pass, lc logcall.LogCall, extraPatterns []string) {
	keywords := defaultSensitiveKeywords
	for _, p := range extraPatterns {
		normalised := normalise(p)
		if normalised != "" {
			keywords = append(keywords, normalised)
		}
	}

	if lc.Msg != "" {
		normalized := normalise(lc.Msg)
		for _, kw := range keywords {
			if strings.Contains(normalized, kw) {
				pass.Report(analysis.Diagnostic{
					Pos:     lc.MsgNode.Pos(),
					Message: "loglint: log message may contain sensitive data",
				})
				return
			}
		}
	}

	idents := logcall.CollectIdents(lc.MsgNode)
	for _, id := range idents {
		normalized := normalise(id)
		for _, kw := range keywords {
			if strings.Contains(normalized, kw) {
				pass.Report(analysis.Diagnostic{
					Pos:     lc.MsgNode.Pos(),
					Message: "loglint: log message may contain sensitive data",
				})
				return
			}
		}
	}

	for _, arg := range lc.Args {
		call, ok := arg.(*ast.CallExpr)
		if !ok {
			continue
		}
		sel, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			continue
		}
		ident, ok := sel.X.(*ast.Ident)
		if !ok {
			continue
		}
		if ident.Name != "zap" {
			continue
		}
		if len(call.Args) < 1 {
			continue
		}
		lit, ok := call.Args[0].(*ast.BasicLit)
		if !ok || lit.Kind != token.STRING {
			continue
		}
		key, err := strconv.Unquote(lit.Value)
		if err != nil {
			continue
		}
		normalized := normalise(key)
		for _, kw := range keywords {
			if strings.Contains(normalized, kw) {
				pass.Report(analysis.Diagnostic{
					Pos:     lit.Pos(),
					Message: "loglint: log message may contain sensitive data",
				})
				return
			}
		}
	}
}

func normalise(s string) string {
	return strings.ToLower(strings.ReplaceAll(s, "_", ""))
}
