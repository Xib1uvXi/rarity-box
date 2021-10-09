package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"runtime"
	"strconv"
	"time"
)

var Logger *zap.Logger

func init() {
	if Logger == nil {
		newLog()
	}
}

func newLog() {
	eConfig := zap.NewProductionEncoderConfig()
	eConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	eConfig.EncodeTime = timeEncoder
	eConfig.EncodeCaller = callerEncoder

	consoleEncoder := zapcore.NewConsoleEncoder(eConfig)

	core := zapcore.NewCore(consoleEncoder, os.Stdout, zap.DebugLevel)

	opts := newLogOptions()
	opts = append(opts, zap.AddCaller())

	Logger = zap.New(core, opts...)
}

func newLogOptions() []zap.Option {
	return []zap.Option{
		zap.AddStacktrace(zapcore.ErrorLevel),
	}
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func callerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(printCaller())
}

func printCaller() string {
	if pc, _, lineNo, ok := runtime.Caller(6); ok {
		return runtime.FuncForPC(pc).Name() + ":" + strconv.FormatInt(int64(lineNo), 10)
	}

	return ""
}
