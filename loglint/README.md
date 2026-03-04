# loglint

Линтер, который проверяет логи в коде на соответствие правилам оформления. Работает с `log/slog` и `go.uber.org/zap`, можно подключить как плагин к golangci-lint.


## Правила

| Правило | ID | Что проверяет |
|---|---|---|
| Строчная буква | `lowercase` | Сообщение начинается с маленькой буквы |
| Английский язык | `english` | Только латиница, без кириллицы и прочего |
| Спецсимволы | `no_special_chars` | Нет эмодзи, `!!!`, `???` и восклицательных знаков |
| Чувствительные данные | `no_sensitive_data` | Нет паролей, токенов, ключей в логах |

Примеры того, что линтер поймает:

```go
slog.Info("Starting server")          // <- заглавная буква
slog.Info("запуск сервера")            // <- не английский
log.Info("server started! 🚀")        // <- эмодзи и !
log.Info("user password: " + password) // <- пароль в логах
```

## Установка

```bash
go install github.com/GrayC9/Selectel/cmd/loglint@latest
```

## Использование

Запуск на проекте:

```bash
loglint ./...
```

Можно отключить отдельные правила через флаги:

```bash
loglint -lowercase=false ./...
```

Вывод выглядит примерно так:

```
main.go:15:2: loglint: log message must start with a lowercase letter
main.go:18:2: loglint: log message must contain only English characters
main.go:24:2: loglint: log message may contain sensitive data
```

Если нужно подавить предупреждение на конкретной строке:

```go
slog.Info("Starting server") //nolint:loglint
```

## Быстрый старт (example/)

В каталоге `example/` есть готовый пример с "плохими" и "хорошими" лог-вызовами. Запуск одной командой:

```bash
bash example/run.sh
```

Вывод:

```
example/main.go:17:12: loglint: log message must start with a lowercase letter
example/main.go:21:13: loglint: log message must contain only English characters
example/main.go:25:12: loglint: log message must not contain emoji
example/main.go:26:13: loglint: log message must not contain exclamation marks
example/main.go:27:12: loglint: log message must not contain repeated punctuation
example/main.go:31:12: loglint: log message may contain sensitive data
example/main.go:32:13: loglint: log message may contain sensitive data
```

## Проверка на реальных проектах

Линтер протестирован на open-source проектах:

| Проект | Находки | Крашей |
|---|---|---|
| [golang/example](https://github.com/golang/example) | 1 (lowercase) | 0 |
| [lmittmann/tint](https://github.com/lmittmann/tint) | 5 (lowercase) | 0 |
| [uber-go/zap](https://github.com/uber-go/zap) | 22 (lowercase, special chars) | 0 |

## Интеграция с golangci-lint

Собираем плагин:

```bash
go build -tags plugin -buildmode=plugin -o loglint.so ./plugin/
```

Добавляем в `.golangci.yml`:

```yaml
linters-settings:
  custom:
    loglint:
      path: loglint.so
      description: "Проверка лог-сообщений"
      original-url: "github.com/GrayC9/Selectel"

linters:
  enable:
    - loglint
```

## Конфигурация

Можно создать `.loglint.yaml` в корне проекта, чтобы настроить правила и добавить свои паттерны для поиска чувствительных данных:

```yaml
rules:
  lowercase: true
  english: true
  no_special_chars: true
  no_sensitive_data: true

sensitive_patterns:
  - "ssn"
  - "credit_card"
```

По умолчанию все правила включены. Встроенный список ключевых слов: `password`, `passwd`, `pwd`, `secret`, `token`, `api_key`, `auth`, `credential`, `private_key` и т.д.

## Сборка и тесты

```bash
make build       # собрать
make test        # тесты с -race
make cover       # покрытие
make vet         # go vet
make plugin-build # собрать плагин для golangci-lint
make all         # vet + test + build
```

## Структура проекта

```
cmd/loglint/         — точка входа (standalone)
pkg/analyzer/        — основной анализатор
  logcall/           — детектор вызовов логгеров
  rules/             — реализация правил
plugin/              — плагин для golangci-lint
config/              — загрузка конфигурации из yaml
example/             — пример использования с run.sh
testdata/            — тестовые файлы для analysistest
```
