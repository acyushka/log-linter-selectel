package english

import (
	"log"
	"log/slog"
	"testing"
)

func TestEnglish(t *testing.T) {

	slog.Info("user logged in") // want

	log.Println("user logged in 123") // want

	log.Println("user вошел") // want "лог должен быть только на английском языке"

	log.Println("user вошел in") // want "лог должен быть только на английском языке"

	log.Println("user 你好") // want "лог должен быть только на английском языке"
}
