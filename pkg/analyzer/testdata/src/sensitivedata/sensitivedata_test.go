package sensitivedata

import (
	"log/slog"
	"testing"
)

func TestSensitiveData(t *testing.T) {
	var pwd = "secret123"
	var creditCard = "4444 4444 4444 4444"

	slog.Info("user logged in") // want

	slog.Info("password is required") // want

	slog.Info("token expired") // want

	slog.Info("credit card: " + creditCard) // want "лог содержит потенциально чувствительные данные"

	slog.Info("password: " + pwd) // want "лог содержит потенциально чувствительные данные"
}
