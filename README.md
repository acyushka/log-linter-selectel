# log-linter

Линтер для Go, совместимый с golangci-lint, 
который анализирует лог-записи в коде и проверяет их 
соответствие установленным правилам.

## Правила проверки

### 1. Первая буква
- Сообщение должно начинаться со строчной буквы

### 2. Английский язык
- Сообщение должно быть на английском


### 3. Эмодзи и спецсимволы
- Запрещено использование эмодзи и спецсиволов в логах

### 4. Чувствительные данные
- Запрещено логирование паролей, токенов и ключей


## Поддерживаемые логгеры
- `log/slog`
- `go.uber.org/zap`

# Подключение к golangci-lint

1. **Склонировать репозитории, если их нет**
   ```bash
   git clone https://github.com/golangci/golangci-lint
   git clone https://github.com/acyushka/log-linter-selectel
   ```

2. **Перенести файл плагина в golangci-lint**
   ```bash
   cp loglinter/pkg/plugin/loglinter.go golangci-lint/pkg/golinters/loglinter/
   ```

3. **Зарегистрировать линтер**
   
   Открыть `golangci-lint/pkg/lint/lintersdb/builder_linter.go` и добавить:

   ```go
   import "github.com/golangci/golangci-lint/v2/pkg/golinters/loglinter"
   // ....
   // ....
   // В методе Build() в return []*linter.Config{}:
    linter.NewConfig(loglinter.New(nil)).WithLoadForGoAnalysis(),
   ```

4. **Собрать golangci-lint**
   ```bash
   make build
   ```

5. **Линтер готов к использованию**
    В проекте, где он будет использоваться можно скопировать конфиг из pkg/config_example

   ```bash
   ./golangci-lint run --enable=loglinter ./...
   ```
