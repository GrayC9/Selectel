package sensitive

import "log/slog"

var password = "hunter2"

func examples() {
	slog.Info("user login successful")  // valid
	slog.Info("user password: hunter2") // want `loglint: log message may contain sensitive data`
	slog.Info("value: " + password)     // want `loglint: log message may contain sensitive data`
}
