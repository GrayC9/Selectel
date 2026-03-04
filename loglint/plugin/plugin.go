//go:build plugin

package main

import (
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"

	"github.com/GrayC9/Selectel/pkg/analyzer"
)

func init() {
	register.Plugin("loglint", func(conf any) (register.LinterPlugin, error) {
		return &Plugin{}, nil
	})
}

type Plugin struct{}

func (p *Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{analyzer.Analyzer}, nil
}

func (p *Plugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}
