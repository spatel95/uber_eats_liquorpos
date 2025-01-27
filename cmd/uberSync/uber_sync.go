package uberSync

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func SetupLogger() *zap.SugaredLogger {
	pe := zap.NewProductionEncoderConfig()
	pe.EncodeLevel = zapcore.CapitalColorLevelEncoder
	pe.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleEncoder := zapcore.NewConsoleEncoder(pe)
	// set level
	core := zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)
	logger := zap.New(core)
	return logger.Sugar()
}
