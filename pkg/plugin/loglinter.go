package loglinter

import (
	"loglinter/pkg/analyzer"
	// заменить после вставки на "github.com/golangci/golangci-lint/v2/pkg/golinters/loglinter/analyzer"

	"github.com/golangci/golangci-lint/v2/pkg/goanalysis"
	"golang.org/x/tools/go/analysis"
)

func New() *goanalysis.Linter {
	return goanalysis.NewLinter(
		"loglinter",
		"analyze logs in the code and check their compliance with rules",
		[]*analysis.Analyzer{analyzer.Analyzer},
		nil,
	).WithLoadMode(goanalysis.LoadModeSyntax)
}
