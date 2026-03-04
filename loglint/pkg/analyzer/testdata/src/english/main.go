package english

import "log/slog"

func examples() {
	slog.Info("starting server")   // valid
	slog.Info("запуск сервера")    // want `loglint: log message must contain only English characters`
	slog.Info("server on port 80") // valid
}
