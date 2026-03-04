package logcall

import (
	"go/ast"
	"go/token"
	"strconv"
	"strings"

	"golang.org/x/tools/go/analysis"
)

type LogCall struct {
	Msg      string
	MsgNode  ast.Expr
	Args     []ast.Expr
	CallExpr *ast.CallExpr
}

var logMethods = map[string]bool{
	"Info":   true,
	"Error":  true,
	"Warn":   true,
	"Debug":  true,
	"Fatal":  true,
	"Panic":  true,
	"Infof":  true,
	"Errorf": true,
	"Warnf":  true,
	"Debugf": true,
	"Fatalf": true,
	"Panicf": true,
	"Infow":  true,
	"Errorw": true,
	"Warnw":  true,
	"Debugw": true,
	"Fatalw": true,
	"Panicw": true,
}

var knownLogPackages = map[string]bool{
	"slog":   true,
	"log":    true,
	"logger": true,
	"zap":    true,
}

func Detect(pass *analysis.Pass) []LogCall {
	var calls []LogCall
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}
			sel, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}
			if !logMethods[sel.Sel.Name] {
				return true
			}
			if !isLogReceiver(sel.X) {
				return true
			}
			if len(call.Args) == 0 {
				return true
			}
			lc := LogCall{
				CallExpr: call,
				MsgNode:  call.Args[0],
			}
			if len(call.Args) > 1 {
				lc.Args = call.Args[1:]
			}
			lc.Msg = extractStringValue(call.Args[0])
			calls = append(calls, lc)
			return true
		})
	}
	return calls
}

func isLogReceiver(expr ast.Expr) bool {
	switch x := expr.(type) {
	case *ast.Ident:
		return knownLogPackages[x.Name]
	case *ast.SelectorExpr:
		if ident, ok := x.X.(*ast.Ident); ok {
			if knownLogPackages[ident.Name] {
				return true
			}
		}
		return isLogReceiver(x.X)
	case *ast.CallExpr:
		if sel, ok := x.Fun.(*ast.SelectorExpr); ok {
			return isLogReceiver(sel.X)
		}
		if ident, ok := x.Fun.(*ast.Ident); ok {
			return knownLogPackages[ident.Name]
		}
		return false
	default:
		return true
	}
}

func extractStringValue(expr ast.Expr) string {
	switch e := expr.(type) {
	case *ast.BasicLit:
		if e.Kind == token.STRING {
			s, err := strconv.Unquote(e.Value)
			if err != nil {
				return ""
			}
			return s
		}
	case *ast.BinaryExpr:
		if e.Op == token.ADD {
			left := extractStringValue(e.X)
			if left != "" {
				return left
			}
			return extractStringValue(e.Y)
		}
	case *ast.CallExpr:
		if sel, ok := e.Fun.(*ast.SelectorExpr); ok {
			if ident, ok := sel.X.(*ast.Ident); ok {
				if ident.Name == "fmt" && strings.HasPrefix(sel.Sel.Name, "Sprintf") {
					if len(e.Args) > 0 {
						return extractStringValue(e.Args[0])
					}
				}
			}
		}
	}
	return ""
}

func CollectIdents(expr ast.Expr) []string {
	var idents []string
	ast.Inspect(expr, func(n ast.Node) bool {
		if id, ok := n.(*ast.Ident); ok {
			idents = append(idents, id.Name)
		}
		return true
	})
	return idents
}
