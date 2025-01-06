// By Saif Hamdan, Team Lead
// Date: 2025/1/6
package logger

import (
	"log"
	"os"

	"github.com/saifhamdan/go-apigateway-blueprint/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	// DPanic, Panic and Fatal level can not be set by user
	DebugLevelStr   string = "debug"
	InfoLevelStr    string = "info"
	WarningLevelStr string = "warning"
	ErrorLevelStr   string = "error"
)

type Logger struct {
	*zap.SugaredLogger
}

func NewLogger(cfg *config.Config) *Logger {
	log.Printf("Initializing logger %s", cfg.LogFile)

	var level zapcore.Level
	switch cfg.LogLevel {
	case DebugLevelStr:
		level = zap.DebugLevel
	case InfoLevelStr:
		level = zap.InfoLevel
	case WarningLevelStr:
		level = zap.WarnLevel
	case ErrorLevelStr:
		level = zap.ErrorLevel
	default:
		log.Fatalf("unknown log level %s", cfg.LogLevel)
	}

	ws := zapcore.AddSync(&lumberjack.Logger{
		Filename:   cfg.LogFile,
		MaxSize:    1024, //MB
		MaxBackups: 30,
		MaxAge:     30, //days
		Compress:   false,
	})

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // Human-readable time format

	core := zapcore.NewCore(
		// use NewConsoleEncoder for human readable output
		zapcore.NewJSONEncoder(encoderConfig),
		// write to stdout as well as log files
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), ws),
		zap.NewAtomicLevelAt(level),
	)

	var _globalLogger *zap.Logger
	if cfg.AppEnv == "development" {
		_globalLogger = zap.New(core, zap.AddCaller(), zap.Development())
	} else {
		_globalLogger = zap.New(core)
	}

	zap.ReplaceGlobals(_globalLogger)

	sugarredLogger := _globalLogger.Sugar()

	return &Logger{
		sugarredLogger,
	}
}
