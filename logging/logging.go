package logging

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zlog *zap.SugaredLogger

func InitLogger() {
	writeSyncer := getLogWriter()
	encoderFile := getEncoder()
	coreFile := zapcore.NewCore(encoderFile, writeSyncer, zapcore.InfoLevel)

	encoderConsole := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	coreConsole := zapcore.NewCore(encoderConsole, zapcore.Lock(os.Stdout), zap.ErrorLevel)

	coreTee := zapcore.NewTee(
		coreFile,
		coreConsole,
	)

	logger := zap.New(coreTee)
	zlog = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename: "./pick.log",
		// set the log file size to 10m
		MaxSize: 10,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func Info(args ...interface{}) {
	zlog.Info(args)
}

func Warn(args ...interface{}) {
	zlog.Warn(args)
}

func Debug(args ...interface{}) {
	zlog.Debug(args)
}

func Error(args ...interface{}) {
	zlog.Error(args)
}

func Panic(args ...interface{}) {
	zlog.Panic(args)
}

func Fatal(args ...interface{}) {
	zlog.Fatal(args)
}
