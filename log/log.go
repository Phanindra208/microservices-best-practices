package log

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Local initializes new logger for local environment.
func Local() {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	config.DisableStacktrace = true
	l, err := config.Build()
	if err != nil {
		panic(fmt.Sprintf("failed to build development logger: %v", err))
	}

	replaceWith(l)
}

// Development initializes new logger for development environment.
func Development() {
	config := zap.NewDevelopmentConfig()
	config.EncoderConfig = zap.NewProductionEncoderConfig()
	config.Encoding = "json"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.DisableStacktrace = true
	l, err := config.Build()
	if err != nil {
		panic(fmt.Sprintf("failed to build development logger: %v", err))
	}

	replaceWith(l)
}

// Production initializes new logger for production environment.
func Production() {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.DisableStacktrace = true

	l, err := config.Build()
	if err != nil {
		panic(fmt.Sprintf("failed to build production logger: %v", err))
	}

	replaceWith(l)
}

func replaceWith(l *zap.Logger) {
	zap.ReplaceGlobals(l)
	zap.RedirectStdLog(l)
}

// Sugar returns a new Sugaredlogger.
func Sugar(fields ...zapcore.Field) *zap.SugaredLogger {
	return zap.L().With(fields...).Sugar()
}
