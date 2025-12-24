package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	l *zap.Logger
}

func NewZapLogger() (*ZapLogger, error) {
	cfg := zap.NewDevelopmentConfig()

	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	cfg.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	cfg.EncoderConfig.TimeKey = "ts"
	cfg.EncoderConfig.MessageKey = "msg"
	cfg.EncoderConfig.CallerKey = "caller"

	logger, err := cfg.Build()
	if err != nil {
		return nil, err
	}

	return &ZapLogger{l: logger}, nil
}

func (z *ZapLogger) Info(msg string, fields ...Field) {
	z.l.Info(msg, toZap(fields)...)
}

func (z *ZapLogger) Error(msg string, fields ...Field) {
	z.l.Error(msg, toZap(fields)...)
}

func (z *ZapLogger) Warn(msg string, fields ...Field) {
	z.l.Warn(msg, toZap(fields)...)
}

func toZap(fields []Field) []zap.Field {
	zf := make([]zap.Field, 0, len(fields))
	for _, f := range fields {
		zf = append(zf, zap.Any(f.Key, f.Value))
	}
	return zf
}
