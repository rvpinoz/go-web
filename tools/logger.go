package tools

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

const (
	serviceLog = "log/go-web.log"
)

var (
	Log *zap.Logger
)

func Logger() {
	serviceWriteSyncer := getLogWriter(serviceLog)
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(encoder, serviceWriteSyncer, zapcore.DebugLevel)
	Log = zap.New(core)
}

func getLogWriter(logFile string) zapcore.WriteSyncer {
	file, _ := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	return zapcore.AddSync(file)
}

func Info(msg string, desc string, t time.Time) {
	Log.Info(msg, zap.String("description", desc), zap.String("timestamp", DateNow("2006-01-02 15:04:05")), zap.Duration("latency", time.Since(t)))
}

func Warn(msg string, desc string, t time.Time) {
	Log.Warn(msg, zap.String("description", desc), zap.String("timestamp", DateNow("2006-01-02 15:04:05")), zap.Duration("latency", time.Since(t)))
}

func Error(msg string, desc string, t time.Time) {
	Log.Error(msg, zap.String("description", desc), zap.String("timestamp", DateNow("2006-01-02 15:04:05")), zap.Duration("latency", time.Since(t)))
}

func Fatal(msg string, desc string, t time.Time) {
	Log.Fatal(msg, zap.String("description", desc), zap.String("timestamp", DateNow("2006-01-02 15:04:05")), zap.Duration("latency", time.Since(t)))
}

func Panic(msg string, desc string, t time.Time) {
	Log.Panic(msg, zap.String("description", desc), zap.String("timestamp", DateNow("2006-01-02 15:04:05")), zap.Duration("latency", time.Since(t)))
}
