package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger zap.Logger

func NewLogger(level, encoding string) *Logger {
	var lv zapcore.Level

	switch level {
	case "debug":
		lv = zapcore.DebugLevel
	case "info":
		lv = zapcore.InfoLevel
	case "warn":
		lv = zapcore.WarnLevel
	case "error":
		lv = zapcore.ErrorLevel
	default:
		lv = zapcore.InfoLevel
	}

	logConfig := zap.Config{
		Level:             zap.NewAtomicLevelAt(lv), // 日志级别
		Development:       false,                    // 开发模式，堆栈跟踪
		DisableStacktrace: false,                    // 关闭自动堆栈捕获
		Encoding:          encoding,                 // 输出格式 console 或 json
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			CallerKey:      "caller",
			LineEnding:     "\n",
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.EpochTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		InitialFields:    nil,
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, err := logConfig.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	return (*Logger)(logger)
}
