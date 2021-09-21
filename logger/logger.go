package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

const (
	envLogLevel  = "LOG_LEVEL"
	envLogOutput = "LOG_OUTPUT"
)

var appLog *LogClass

type LogClass struct {
	logger      *zap.Logger
	sugarLogger *zap.SugaredLogger
}

type bookstoreLogger interface {
	Debug(format string, v ...interface{})
	Info(format string, v ...interface{})
	Warning(format string, v ...interface{})
	Error(format string, v ...interface{})
}

func GetLogger() bookstoreLogger {
	if appLog == nil {
		appLog = new(LogClass)
	}
	return appLog
}

func (l *LogClass) Debug(format string, v ...interface{}) {
	if len(v) == 0 {
		appLog.sugarLogger.Debug(format)
	} else {
		appLog.sugarLogger.Debugf(format, v...)
	}
	appLog.sugarLogger.Sync()
}

func (l *LogClass) Info(format string, v ...interface{}) {
	if len(v) == 0 {
		appLog.sugarLogger.Info(format)
	} else {
		appLog.sugarLogger.Infof(format, v...)
	}
	appLog.sugarLogger.Sync()
}

func (l *LogClass) Warning(format string, v ...interface{}) {
	if len(v) == 0 {
		appLog.sugarLogger.Warn(format)
	} else {
		appLog.sugarLogger.Warnf(format, v...)
	}
	appLog.sugarLogger.Sync()
}

func (l *LogClass) Error(format string, v ...interface{}) {
	if len(v) == 0 {
		appLog.sugarLogger.Error(format)
	} else {
		appLog.sugarLogger.Errorf(format, v...)
	}
	appLog.sugarLogger.Sync()
}

func init() {
	var err error
	// Construct encoderconfig
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "severity",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     "\n",
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}

	// Construct config
	config := zap.Config{
		OutputPaths:      []string{getOutput()},
		ErrorOutputPaths: []string{getOutput()},
		Level:            zap.NewAtomicLevelAt(getLevel()),
		Development:      true,
		Encoding:         "json",
		EncoderConfig:    encoderConfig,
	}

	//Can construct a logger
	appLog.logger, err = config.Build()
	if err != nil {
		panic(err)
	}

	// The then Sugarlogger
	appLog.sugarLogger = appLog.logger.Sugar()
}

func getLevel() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(envLogLevel))) {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.DebugLevel
	}
}

func getOutput() string {
	//set use log output file, if not set env variable use terminal instead
	output := strings.TrimSpace(os.Getenv(envLogOutput))
	if output == "" {
		return "stdout"
	}
	return output
}
