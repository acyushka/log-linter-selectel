package english

import (
	"log/slog"
	"testing"
)

func TestEnglish(t *testing.T) {

	slog.Info("user logged in") // want

	slog.Info("user logged in 123") // want

	slog.Info("user вошел") // want "лог должен быть только на английском языке"

	slog.Info("user вошел in") // want "лог должен быть только на английском языке"

	slog.Info("user 你好") // want "лог должен быть только на английском языке"
}
