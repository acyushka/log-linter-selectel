package analyzer

import (
	"go/ast"
	"go/token"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

var sensitiveData = []string{
	"password", "passwd", "pwd",
	"secret", "token", "api_key", "apikey", "key",
	"auth", "credential", "credit", "ssn",
}

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			selector, ok := call.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}

			if !isLogger(selector.X) {
				return true
			}

			validateLogMessage(pass, call)

			return true
		})
	}

	return nil, nil
}

func isLogger(expr ast.Expr) bool {
	for {
		switch e := expr.(type) {
		case *ast.Ident:
			if e.Name == "slog" || e.Name == "log" || e.Name == "zap" {
				return true
			}
			return false

		case *ast.SelectorExpr:
			expr = e.X

		case *ast.CallExpr:
			if sel, ok := e.Fun.(*ast.SelectorExpr); ok {
				expr = sel.X
			} else {
				return false
			}

		default:
			return false
		}
	}
}

func validateLogMessage(pass *analysis.Pass, call *ast.CallExpr) {
	if len(call.Args) == 0 {
		return
	}

	msg := call.Args[0]

	if HasSensitiveData(msg) {
		pass.Reportf(call.Pos(), errHasSensitiveData)
	}

	lit, ok := msg.(*ast.BasicLit)
	if !ok {
		return
	}

	if lit.Kind == token.STRING {
		msgStr, err := strconv.Unquote(lit.Value)
		if err != nil {
			return
		}

		if errors := ValidateMsg(msgStr); len(errors) > 0 {
			for _, err := range errors {
				pass.Reportf(lit.Pos(), "%s", err)
			}
		}
	}
}

func HasSensitiveData(msg ast.Expr) bool {
	expr, ok := msg.(*ast.BinaryExpr)
	if !ok || expr.Op != token.ADD {
		return false
	}

	if _, ok := expr.Y.(*ast.Ident); ok {
		if lit, ok := expr.X.(*ast.BasicLit); ok && lit.Kind == token.STRING {
			msgLower := strings.ToLower(lit.Value)

			for _, word := range sensitiveData {
				if strings.Contains(msgLower, word) {
					return true
				}
			}
		}
	}

	return false
}

func ValidateMsg(msg string) []string {
	var errors = []string{}

	if msg == "" {
		errors = append(errors, errIsEmpty)

		return errors
	}

	var hasEmoji, hasNotEnglish, hasNotLower bool

	for i, r := range msg {
		if i == 0 {
			if !(unicode.IsLetter(r) && unicode.IsLower(r)) {
				hasNotLower = true
			}
			break
		}
	}

	for _, r := range msg {
		if !isEnglish(r) {
			if isEmoji(r) {
				hasEmoji = true
			} else {
				hasNotEnglish = true
			}
		}

		if hasEmoji && hasNotEnglish {
			break
		}
	}

	if hasNotLower {
		errors = append(errors, errHasNotLower)
	}
	if hasEmoji {
		errors = append(errors, errHasEmoji)
	}
	if hasNotEnglish {
		errors = append(errors, errHasNotEnglish)
	}

	return errors
}

func isEnglish(r rune) bool {
	if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
		return true
	}

	if unicode.IsSpace(r) || r == ':' || unicode.IsDigit(r) {
		return true
	}

	return false
}

func isEmoji(r rune) bool {
	if (r >= 0x1F600 && r <= 0x1F64F) ||
		(r >= 0x1F300 && r <= 0x1F5FF) ||
		(r >= 0x1F680 && r <= 0x1F6FF) ||
		(r >= 0x1F900 && r <= 0x1F9FF) ||
		(r >= 0x1FA70 && r <= 0x1FAFF) ||
		(r >= 0x2600 && r <= 0x26FF) ||
		(r >= 0x2700 && r <= 0x27BF) ||
		(r >= 0x1F1E6 && r <= 0x1F1FF) {
		return true
	}

	special := "!@#$%^&*()_+{}[];<>,.?~\\|/`'\""
	for _, s := range special {
		if r == s {
			return true
		}
	}

	return false
}
