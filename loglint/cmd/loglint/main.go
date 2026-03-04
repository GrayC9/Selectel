package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/GrayC9/Selectel/pkg/analyzer"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}
