package lowercase

import "log/slog"

func examples() {
	slog.Info("starting server")   // valid
	slog.Info("Starting server")   // want `loglint: log message must start with a lowercase letter`
	slog.Info("another valid msg") // valid
	slog.Info("Bad message here")  // want `loglint: log message must start with a lowercase letter`
	slog.Info("")                  // valid - empty string
	slog.Info("123 items")         // valid - starts with digit
}
