package logger

import "go.uber.org/zap/zapcore"

var encoderConfig = zapcore.EncoderConfig{
	TimeKey:        "timestamp",
	LevelKey:       "level",
	NameKey:        "logger",
	CallerKey:      "caller",
	FunctionKey:    "func",
	MessageKey:     "msg",
	StacktraceKey:  "stacktrace",
	LineEnding:     zapcore.DefaultLineEnding,
	EncodeLevel:    zapcore.LowercaseLevelEncoder,
	EncodeTime:     zapcore.ISO8601TimeEncoder,
	EncodeDuration: zapcore.SecondsDurationEncoder,
	EncodeCaller:   zapcore.ShortCallerEncoder,
}

func withDefaults(l Logger) Logger {
	if l.MaxAge == 0 {
		l.MaxAge = 30
	}

	if l.MaxBackups == 0 {
		l.MaxBackups = 10
	}

	if l.MaxSize == 0 {
		l.MaxSize = 100
	}
	return l
}

type fileConfig struct {
	Filename   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
}
