package analyzer_test

import (
	"loglinter/internal/analyzer"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type Test struct {
	name string
	msg  string
	want []string
}

const (
	errIsEmpty          = "лог должен быть непустым"
	errHasEmoji         = "лог не должен содержать спецсимволы или эмодзи"
	errHasNotLower      = "лог должен начинаться со строчной буквы"
	errHasNotEnglish    = "лог должен быть только на английском языке"
	errHasSensitiveData = "лог содержит потенциально чувствительные данные"
)

func TestFirstSymbol(t *testing.T) {
	tests := []Test{
		{
			name: "строчная буква - OK",
			msg:  "user logged in",
			want: []string{},
		},
		{
			name: "заглавная буква - ERROR",
			msg:  "User logged in",
			want: []string{errHasNotLower},
		},
		{
			name: "1 символ цифра - ERROR",
			msg:  "1user logged in",
			want: []string{errHasNotLower},
		},
		{
			name: "1 символ спецсимвол - ERROR",
			msg:  "$user logged in",
			want: []string{errHasNotLower, errHasEmoji},
		},
		{
			name: "пустая строка - ERROR",
			msg:  "",
			want: []string{errIsEmpty},
		},
		{
			name: "пробел в начале - ERROR",
			msg:  " user logged in",
			want: []string{errHasNotLower},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := analyzer.ValidateMsg(tt.msg)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("ValidateMsg(%q) = %v, want %v", tt.msg, got, tt.want)
			}
		})
	}
}
