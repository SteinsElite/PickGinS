package logging

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	_globalZ *zap.Logger
)

func InitLogger(logDir string) {
	// First, define our level-handling logic.
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	fileDebugging := getLogWriter(logDir + "/pick_info.log")
	fileError := getLogWriter(logDir + "/pick_error.log")

	consoleDebugging := zapcore.Lock(os.Stdout)
	consoleError := zapcore.Lock(os.Stderr)

	encoder := getEncoder()

	coreTee := zapcore.NewTee(
		zapcore.NewCore(encoder, fileError, highPriority),
		zapcore.NewCore(encoder, fileDebugging, lowPriority),
		zapcore.NewCore(encoder, consoleError, highPriority),
		zapcore.NewCore(encoder, consoleDebugging, lowPriority),
	)

	// From a zapcore.Core, it's easy to construct a Logger.
	_globalZ = zap.New(coreTee)
}

// SyncLogger please make sure to sync the logger before exit the program
func SyncLogger() {
	_globalZ.Sync()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()

	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(filePath string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename: filePath,
		// set the log file size to 10 megabytes
		MaxSize: 10,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func Z() *zap.Logger {
	return _globalZ
}
