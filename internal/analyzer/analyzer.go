package analyzer

import "golang.org/x/tools/go/analysis"

var Analyzer = &analysis.Analyzer{
	Name: "log-linter",
	Doc:  "analyze logs in the code and check their compliance with rules.",
	Run:  run,
}
