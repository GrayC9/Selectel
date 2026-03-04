package logcall

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestExtractStringValue(t *testing.T) {
	tests := []struct {
		name string
		code string
		want string
	}{
		{"basic string", `"hello world"`, "hello world"},
		{"empty string", `""`, ""},
		{"integer", `42`, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			expr, err := parser.ParseExpr(tt.code)
			if err != nil {
				t.Fatalf("parse error: %v", err)
			}
			got := extractStringValue(expr)
			if got != tt.want {
				t.Errorf("extractStringValue(%q) = %q, want %q", tt.code, got, tt.want)
			}
		})
	}
}

func TestCollectIdents(t *testing.T) {
	expr, err := parser.ParseExpr(`"hello" + password`)
	if err != nil {
		t.Fatal(err)
	}
	idents := CollectIdents(expr)
	found := false
	for _, id := range idents {
		if id == "password" {
			found = true
		}
	}
	if !found {
		t.Error("expected to find ident 'password'")
	}
}

func TestIsLogReceiver(t *testing.T) {
	tests := []struct {
		name string
		expr ast.Expr
		want bool
	}{
		{
			"known ident slog",
			&ast.Ident{Name: "slog"},
			true,
		},
		{
			"known ident logger",
			&ast.Ident{Name: "logger"},
			true,
		},
		{
			"unknown ident",
			&ast.Ident{Name: "foo"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isLogReceiver(tt.expr)
			if got != tt.want {
				t.Errorf("isLogReceiver() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractStringValueConcat(t *testing.T) {
	fset := token.NewFileSet()
	src := `package p; var x = "hello " + "world"`
	f, err := parser.ParseFile(fset, "test.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	var binExpr *ast.BinaryExpr
	ast.Inspect(f, func(n ast.Node) bool {
		if b, ok := n.(*ast.BinaryExpr); ok {
			binExpr = b
			return false
		}
		return true
	})
	if binExpr == nil {
		t.Fatal("no binary expr found")
	}
	got := extractStringValue(binExpr)
	if got != "hello " {
		t.Errorf("got %q, want %q", got, "hello ")
	}
}
