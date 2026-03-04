package analyzer

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"

	"github.com/GrayC9/Selectel/pkg/analyzer/logcall"
	"github.com/GrayC9/Selectel/pkg/analyzer/rules"
)

var Analyzer = &analysis.Analyzer{
	Name: "loglint",
	Doc:  "checks log messages for common issues: lowercase start, english only, no special chars, no sensitive data",
	Run:  run,
}

var (
	flagLowercase       bool
	flagEnglish         bool
	flagNoSpecialChars  bool
	flagNoSensitiveData bool
)

func init() {
	Analyzer.Flags.BoolVar(&flagLowercase, "lowercase", true, "check that log messages start with a lowercase letter")
	Analyzer.Flags.BoolVar(&flagEnglish, "english", true, "check that log messages contain only English characters")
	Analyzer.Flags.BoolVar(&flagNoSpecialChars, "no_special_chars", true, "check that log messages do not contain special characters")
	Analyzer.Flags.BoolVar(&flagNoSensitiveData, "no_sensitive_data", true, "check that log messages do not contain sensitive data")
}

func run(pass *analysis.Pass) (interface{}, error) {
	calls := logcall.Detect(pass)

	for _, lc := range calls {
		if isNolint(pass, lc.CallExpr) {
			continue
		}

		if flagLowercase {
			rules.CheckLowercase(pass, lc.Msg, lc.MsgNode)
		}
		if flagEnglish {
			rules.CheckEnglish(pass, lc.Msg, lc.MsgNode)
		}
		if flagNoSpecialChars {
			rules.CheckNoSpecialChars(pass, lc.Msg, lc.MsgNode)
		}
		if flagNoSensitiveData {
			rules.CheckNoSensitiveData(pass, lc, nil)
		}
	}

	return nil, nil
}

func isNolint(pass *analysis.Pass, node ast.Node) bool {
	pos := pass.Fset.Position(node.Pos())
	for _, f := range pass.Files {
		for _, cg := range f.Comments {
			for _, c := range cg.List {
				cp := pass.Fset.Position(c.Pos())
				if cp.Line == pos.Line &&
					strings.Contains(c.Text, "nolint:loglint") {
					return true
				}
			}
		}
	}
	return false
}
