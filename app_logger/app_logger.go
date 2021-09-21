package app_logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"strings"
)

const (
	envLogLevel  = "LOG_LEVEL"
	envLogOutput = "LOG_OUTPUT"
)

var (
	appLog *LogClass
)

type LogClass struct {
	logger      *zap.Logger
	sugarLogger *zap.SugaredLogger
}

func GetLogger() *LogClass {
	return appLog
}

func (l *LogClass) Debug(format string, v ...interface{}) {
	if len(v) == 0 {
		appLog.sugarLogger.Debug(format)
	} else {
		appLog.sugarLogger.Debugf(format, v...)
	}
	err := appLog.sugarLogger.Sync()
	if err != nil {
		log.Fatalln("sugarLogger.Sync failed for Debug function")
	}
}

func (l *LogClass) Info(format string, v ...interface{}) {
	if len(v) == 0 {
		appLog.sugarLogger.Info(format)
	} else {
		appLog.sugarLogger.Infof(format, v...)
	}
	err := appLog.sugarLogger.Sync()
	if err != nil {
		log.Fatalln("sugarLogger.Sync failed for Info function")
	}
}

func (l *LogClass) Warning(format string, v ...interface{}) {
	if len(v) == 0 {
		appLog.sugarLogger.Warn(format)
	} else {
		appLog.sugarLogger.Warnf(format, v...)
	}
	err := appLog.sugarLogger.Sync()
	if err != nil {
		log.Fatalln("sugarLogger.Sync failed for Warning function")
	}
}

func (l *LogClass) Error(format string, v ...interface{}) {
	if len(v) == 0 {
		appLog.sugarLogger.Error(format)
	} else {
		appLog.sugarLogger.Errorf(format, v...)
	}
	err := appLog.sugarLogger.Sync()
	if err != nil {
		log.Fatalln("sugarLogger.Sync failed for Error function")
	}
}

// init() functions can be used within a package block and regardless of how many times that package is imported, the init() function will only be called once.
// https://tutorialedge.net/golang/the-go-init-function/
// how to create singlton
// https://golangbyexample.com/singleton-design-pattern-go/
func init() {
	//the init function is only called once per file in a package, appLog will init only once
	appLog = &LogClass{}
	var err error
	// Construct encoderconfig
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "severity",
		NameKey:        "app_logger",
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

	//Can construct a app_logger
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
