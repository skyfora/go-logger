package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewConsoleEncoder(cfg zapcore.EncoderConfig) zapcore.Core {
	consoleEncoder := zapcore.NewConsoleEncoder(cfg)
	stdout := zapcore.AddSync(os.Stdout)
	return zapcore.NewCore(consoleEncoder, stdout, zap.NewAtomicLevelAt(zap.InfoLevel))
}

func NewFileEncoder(cfg zapcore.EncoderConfig, level zap.AtomicLevel, fileConfig fileConfig) zapcore.Core {
	fileEncoder := zapcore.NewJSONEncoder(cfg)
	file := zapcore.AddSync(&lumberjack.Logger{
		Filename:   fileConfig.Filename,
		MaxSize:    fileConfig.MaxSize,
		MaxBackups: fileConfig.MaxBackups,
		MaxAge:     fileConfig.MaxAge,
	})
	return zapcore.NewCore(fileEncoder, file, level)
}

func NewFileEncoderSeparateLevels(cfg zapcore.EncoderConfig, level zap.AtomicLevel, fileConfig fileConfig) []zapcore.Core {
	cores := []zapcore.Core{}
	levels := []zapcore.Level{zap.DebugLevel, zap.InfoLevel, zap.WarnLevel, zap.ErrorLevel, zap.DPanicLevel, zap.PanicLevel, zap.FatalLevel}

	for _, level := range levels {
		filePathWithLevel := fileConfig.Filename + "." + level.String()
		file := zapcore.AddSync(&lumberjack.Logger{
			Filename:   filePathWithLevel,
			MaxSize:    fileConfig.MaxSize,
			MaxBackups: fileConfig.MaxBackups,
			MaxAge:     fileConfig.MaxAge,
		})
		cores = append(cores, zapcore.NewCore(zapcore.NewJSONEncoder(cfg), file, NewSingleLevelEnabler(level)))
	}

	return cores
}

func NewSingleLevelEnabler(level zapcore.Level) zapcore.LevelEnabler {
	return zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == level
	})
}
