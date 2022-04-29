package log

import (
    "go.uber.org/zap"
)

type zapLogger struct {
    logger  *zap.Logger
}

// toZapFields converts a slice of Field to a slice of zap.Field.
func (zapLog *zapLogger) toZapFields(fields ...Field)  []zap.Field {
    zapFields := make([]zap.Field, 0, len(fields))
    for _, field := range fields {
        zapFields = append(zapFields, zap.Any(field.Key, field.Value))
    }

    return zapFields
}

func (zapLog *zapLogger) Debug(message string, fields ...Field) {
    zapFields := zapLog.toZapFields(fields...)
    zapLog.logger.Debug(message, zapFields...)
}

func (zapLog *zapLogger) Info(message string, fields ...Field) {
    zapFields := zapLog.toZapFields(fields...)
    zapLog.logger.Info(message, zapFields...)
}

func (zapLog *zapLogger) Warning(message string, fields ...Field) {
    zapFields := zapLog.toZapFields(fields...)
    zapLog.logger.Warn(message, zapFields...)
}

func (zapLog *zapLogger) Error(message string, fields ...Field) {
    zapFields := zapLog.toZapFields(fields...)
    zapLog.logger.Error(message, zapFields...)
}

func (zapLog *zapLogger) Fatal(message string, fields ...Field) {
    zapFields := zapLog.toZapFields(fields...)
    zapLog.logger.Fatal(message, zapFields...)
}

func (zapLog *zapLogger) Panic(message string, fields ...Field) {
    zapFields := zapLog.toZapFields(fields...)
    zapLog.logger.Panic(message, zapFields...)
}
