package analyzer

import "golang.org/x/tools/go/analysis"

var Analyzer = &analysis.Analyzer{
	Name: "loglinter",
	Doc:  "analyze logs in the code and check their compliance with rules",
	Run:  run,
}

const (
	errIsEmpty          = "лог должен быть непустым"
	errHasEmoji         = "лог не должен содержать спецсимволы или эмодзи"
	errHasNotLower      = "лог должен начинаться со строчной буквы"
	errHasNotEnglish    = "лог должен быть только на английском языке"
	errHasSensitiveData = "лог содержит потенциально чувствительные данные"
)
