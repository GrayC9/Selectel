package special_chars

import "log/slog"

func examples() {
	slog.Info("started")        // valid
	slog.Info("started 🚀")      // want `loglint: log message must contain only English characters` `loglint: log message must not contain emoji`
	slog.Info("what???")        // want `loglint: log message must not contain repeated punctuation`
	slog.Info("oh no!")         // want `loglint: log message must not contain exclamation marks`
	slog.Info("server running") // valid
}
