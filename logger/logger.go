package logger
import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

var zapLog *zap.Logger

func init() {
    var err error
    config := zap.NewProductionConfig()
    enccoderConfig := zap.NewProductionEncoderConfig()
    zapcore.TimeEncoderOfLayout("Jan _2 15:04:05.000000000")
    enccoderConfig.StacktraceKey = "" // to hide stacktrace info
    config.EncoderConfig = enccoderConfig

    zapLog, err = config.Build(zap.AddCallerSkip(1))
    if err != nil {
        panic(err)
    }
}

func Info(message string, fields ...zap.Field) {
    zapLog.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
    zapLog.Debug(message, fields...)
}

func Debugf(template string, args ...interface{}) {
    zapLog.Sugar().Debugf(template, args...)
}

func Error(message string, fields ...zap.Field) {
    zapLog.Error(message, fields...)
}

func Errorf(template string, args ...interface{}) {
    zapLog.Sugar().Errorf(template, args...)
}

func Fatal(message string, fields ...zap.Field) {
    zapLog.Fatal(message, fields...)
}