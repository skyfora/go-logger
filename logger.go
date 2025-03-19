package logger

import (
	"errors"
	"syscall"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *Logger

type Logger struct {
	// Debug will enable debug level logs.
	Debug bool

	// FilePath is the path to write logs to.  It will not write to file if empty.
	FilePath string

	// WithStdout will write logs to stdout.
	WithStdout bool

	// MaxSize is the maximum size in megabytes of the log file before it gets
	// rotated. It defaults to 100 megabytes.
	MaxSize int

	// MaxAge is the maximum number of days to retain old log files based on the
	// timestamp encoded in their filename.
	MaxAge int

	// MaxBackups is the maximum number of old log files to retain.
	MaxBackups int

	// SeparateLevel will write logs to separate files based on the log level.
	SeparateLevel bool

	zapLog *zap.Logger
}

func Init(l Logger) {
	l = withDefaults(l)

	var level zap.AtomicLevel
	if l.Debug {
		level = zap.NewAtomicLevelAt(zap.DebugLevel)
	} else {
		level = zap.NewAtomicLevelAt(zap.InfoLevel)
	}

	cores := []zapcore.Core{}

	if l.WithStdout {
		cores = append(cores, NewConsoleEncoder(encoderConfig))
	}

	if l.FilePath != "" {
		if l.SeparateLevel {
			cores = append(cores, NewFileEncoderSeparateLevels(encoderConfig, level, fileConfig{
				Filename:   l.FilePath,
				MaxSize:    l.MaxSize,
				MaxBackups: l.MaxBackups,
				MaxAge:     l.MaxAge,
			})...)
		} else {
			cores = append(cores, NewFileEncoder(encoderConfig, level, fileConfig{
				Filename:   l.FilePath,
				MaxSize:    l.MaxSize,
				MaxBackups: l.MaxBackups,
				MaxAge:     l.MaxAge,
			}))
		}
	}

	core := zapcore.NewTee(cores...)
	logger = &Logger{
		Debug:      l.Debug,
		FilePath:   l.FilePath,
		WithStdout: l.WithStdout,
		MaxSize:    l.MaxSize,
		MaxAge:     l.MaxAge,
		MaxBackups: l.MaxBackups,
		zapLog: zap.New(
			core,
			zap.AddCaller(),
			zap.AddCallerSkip(1),
			zap.AddStacktrace(zap.ErrorLevel),
		),
	}
}

func InitEmpty() {
	core := zapcore.NewNopCore()
	logger = &Logger{
		Debug:      true,
		FilePath:   "",
		WithStdout: true,
		MaxSize:    1,
		MaxAge:     1,
		MaxBackups: 1,
		zapLog:     zap.New(core),
	}
}

func Info(message string, fields ...zap.Field) {
	logger.zapLog.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	logger.zapLog.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	logger.zapLog.Error(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	logger.zapLog.Fatal(message, fields...)
}

func Warn(message string, fields ...zap.Field) {
	logger.zapLog.Warn(message, fields...)
}

func DPanic(message string, fields ...zap.Field) {
	logger.zapLog.DPanic(message, fields...)
}

func Panic(message string, fields ...zap.Field) {
	logger.zapLog.Panic(message, fields...)
}

func With(fields ...zap.Field) *zap.Logger {
	return logger.zapLog.With(fields...)
}

func Get() *zap.Logger {
	return logger.zapLog
}

func Sync() {
	err := logger.zapLog.Sync()
	if err != nil && (!errors.Is(err, syscall.EINVAL)) {
		panic(err)
	}
}
