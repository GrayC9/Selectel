package main

import (
	"log/slog"
	"os"
)

var (
	password = "hunter2"
	apiKey   = "sk-1234567890"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// --- Rule 1: lowercase ---
	slog.Info("Starting server on port 8080") // BAD: uppercase first letter
	slog.Info("starting server on port 8080") // OK

	// --- Rule 2: english only ---
	slog.Error("ошибка подключения к базе данных") // BAD: not english
	slog.Error("failed to connect to database")    // OK

	// --- Rule 3: special chars / emoji ---
	slog.Info("server started! 🚀")        // BAD: emoji and exclamation mark
	slog.Error("connection failed!!!")      // BAD: repeated punctuation
	slog.Warn("warning: something wrong...") // BAD: repeated punctuation
	slog.Info("server started successfully") // OK

	// --- Rule 4: sensitive data ---
	slog.Info("user password: " + password)         // BAD: contains password
	slog.Debug("api_key=" + apiKey)                  // BAD: contains api_key
	slog.Info("user login successful") // OK

	// --- All good ---
	logger.Info("request processed", "status", 200)
	logger.Debug("cache hit", "key", "users:list")
	logger.Warn("slow query detected", "duration_ms", 1500)
}