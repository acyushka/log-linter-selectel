package firstsymbol

import (
	"log/slog"
	"testing"
)

func TestFirstSymbol(t *testing.T) {

	slog.Info("user logged in") // want

	slog.Info("User logged in") // want "лог должен начинаться со строчной буквы"

	slog.Info("1user logged in") // want "лог должен начинаться со строчной буквы"

	slog.Info("$user logged in") // want "лог должен начинаться со строчной буквы" "лог не должен содержать спецсимволы или эмодзи"

	slog.Info("") // want "лог должен быть непустым"

	slog.Info(" user logged in") // want "лог должен начинаться со строчной буквы"

	slog.Info(".user logged in") // want "лог должен начинаться со строчной буквы" "лог не должен содержать спецсимволы или эмодзи"

	slog.Info("123 user logged in") // want "лог должен начинаться со строчной буквы"
}
