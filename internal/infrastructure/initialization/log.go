package initialization

import (
	"os"

	"github.com/InstaySystem/is_v2-be/internal/infrastructure/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Log struct {
	zap *zap.Logger
}

func InitLog(cfg config.LogConfig) (*Log, error) {
	level := zapcore.InfoLevel
	if err := level.UnmarshalText([]byte(cfg.Level)); err != nil {
		return nil, err
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   cfg.OutputPath,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
	})

	consoleWriter := zapcore.AddSync(os.Stdout)

	var encoder zapcore.Encoder
	if cfg.Encoding == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	core := zapcore.NewTee(
		zapcore.NewCore(encoder, fileWriter, level),
		zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), consoleWriter, level),
	)

	zap := zap.New(core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)

	return &Log{
		zap,
	}, nil
}

func (l *Log) Close() {
	_ = l.zap.Sync()
}

func (l *Log) Logger() *zap.Logger {
	return l.zap
}