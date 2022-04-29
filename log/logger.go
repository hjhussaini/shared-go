package log

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

// Field for embedding key-value pairs to a log entry.
type Field struct {
    Key     string
    Value   interface{}
}

type Logger interface {
    Debug(message string, fields ...Field)
    Info(message string, fields ...Field)
    Warning(message string, fields ...Field)
    Error(message string, fields ...Field)
    Fatal(message string, fields ...Field)
    Panic(message string, fields ...Field)
}

// NewLogger returns a zapLogger implementation of Logger interface.
func NewLogger(logLevel Level, cores ...zapcore.Core) Logger {
    core := zapcore.NewTee(cores...)
    levelEnabler := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
        return level >= logLevel
    })
    logger := zap.New(core, zap.AddStacktrace(levelEnabler))

    return &zapLogger{logger: logger}
}
