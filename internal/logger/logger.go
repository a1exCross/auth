package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var globalLogger *zap.Logger

// Init - инкапсуляциия логгера для метода Init
func Init(core zapcore.Core, options ...zap.Option) {
	globalLogger = zap.New(core, options...)
}

// Debug - инкапсуляциия логгера для метода Debug
func Debug(msg string, fields ...zap.Field) {
	globalLogger.Debug(msg, fields...)
}

// Info - инкапсуляциия логгера для метода Info
func Info(msg string, fields ...zap.Field) {
	globalLogger.Info(msg, fields...)
}

// Warn - инкапсуляциия логгера для метода Warn
func Warn(msg string, fields ...zap.Field) {
	globalLogger.Warn(msg, fields...)
}

// Error - инкапсуляциия логгера для метода Error
func Error(msg string, fields ...zap.Field) {
	globalLogger.Error(msg, fields...)
}

// Fatal - инкапсуляциия логгера для метода Fatal
func Fatal(msg string, fields ...zap.Field) {
	globalLogger.Fatal(msg, fields...)
}
