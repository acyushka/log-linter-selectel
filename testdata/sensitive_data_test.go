package analyzer_test

import (
	"go/ast"
	"go/token"
	"loglinter/internal/analyzer"
	"strconv"
	"testing"
)

type TestSensitive struct {
	name string
	x    string
	y    string
	want bool
}

func TestSensitiveData(t *testing.T) {
	var creditCard = "5555 5555 5555 5555"
	var pwd = "secret"

	tests := []TestSensitive{
		{
			name: "without sensitive words - OK",
			x:    "user logged in",
			want: false,
		},
		{
			name: "слово password в тексте (без данных) - OK",
			x:    "password is required",
			want: false,
		},
		{
			name: "слово token в тексте (без данных) - OK",
			x:    "token expired",
			want: false,
		},
		{
			name: "sensitive",
			x:    "credit card:",
			y:    creditCard,
			want: true,
		},
		{
			name: "sensitive",
			x:    "password: ",
			y:    pwd,
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var expr ast.Expr

			if tt.y == "" {
				expr = &ast.BasicLit{
					Kind:  token.STRING,
					Value: strconv.Quote(tt.x),
				}
			} else {
				expr = &ast.BinaryExpr{
					Op: token.ADD,
					X: &ast.BasicLit{
						Kind:  token.STRING,
						Value: strconv.Quote(tt.x),
					},
					Y: &ast.Ident{
						Name: tt.y,
					},
				}
			}

			got := analyzer.HasSensitiveData(expr)
			if got != tt.want {
				t.Errorf("HasSensitiveData(%q) = %v, want %v", tt.x+tt.y, got, tt.want)
			}
		})
	}
}
