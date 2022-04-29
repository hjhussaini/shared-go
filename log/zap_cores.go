package log

import (
    "os"
    "github.com/TheZeroSlave/zapsentry"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "gopkg.in/graylog2/go-gelf.v2/gelf"
)

// NewStdoutCore returns a zapcore.Core for sending logs with level higher than DEBUG level to
// standard output.
func NewStdoutCore() zapcore.Core {
    jsonEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
    syncWriter := zapcore.AddSync(os.Stdout)
    levelEnabler := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
        return level >= zapcore.DebugLevel
    })

    return zapcore.NewCore(jsonEncoder, syncWriter, levelEnabler)
}

// NewGraylogCore returns a zapcore.Core for sending logs with level higher than desired level
// to graylog. UDP connection is used because blocking TCP connection is not desired.
func NewGraylogCore(
    dataSourceName string,
    facility string,
    minLevel Level,
    maxLevel Level,
) (zapcore.Core, error) {
    udpWriter, err := gelf.NewUDPWriter(dataSourceName)
    if err != nil {
        return nil, err
    }

    udpWriter.Facility = facility

    jsonEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
    syncWriter := zapcore.AddSync(udpWriter)
    levelEnabler := zap.LevelEnablerFunc(func(level zapcore.Level) bool {
        return level >= minLevel && level <= maxLevel
    })
    core := zapcore.NewCore(jsonEncoder, syncWriter, levelEnabler)
    err = core.Sync()
    if err != nil {
        return nil, err
    }

    return core, nil
}

// NewSentryCore returns a zapcore.Core for sending logs with level higher than
// WARNING level to sentry.
func NewSentryCore(dataSourceName string, tags map[string]string) (zapcore.Core, error) {
    configuration := zapsentry.Configuration{
        Level:  zapcore.ErrorLevel,
        Tags:   tags,
    }
    client := zapsentry.NewSentryClientFromDSN(dataSourceName)
    core, err := zapsentry.NewCore(configuration, client)

    return core, err
}
