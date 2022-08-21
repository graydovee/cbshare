package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var (
	logger *zap.Logger
	sugar  *zap.SugaredLogger
)

func Logger() *zap.Logger {
	if logger == nil {
		InitLogger()
	}
	return logger
}

func Sugar() *zap.SugaredLogger {
	if sugar == nil {
		InitLogger()
	}
	return sugar
}

func Init(log *zap.Logger) {
	logger = log
	sugar = logger.Sugar()
}

func InitLoggerWithOptions(options Options) {
	log := NewLoggerWithLevel(zap.DebugLevel, options)
	Init(log)
}

func InitLogger() {
	InitLoggerWithOptions(DefaultOptions)
}

func NewLoggerWithLevel(level zapcore.Level, opt Options) *zap.Logger {
	atomLevel := zap.NewAtomicLevelAt(level)

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.NanosDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	var encoder zapcore.Encoder
	if opt.EncoderJson {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}
	infoSyncer, errSyncer := NewSink(opt.LogFile)
	core := zapcore.NewCore(encoder, infoSyncer, atomLevel)

	options := make([]zap.Option, 0)
	options = append(options, zap.AddCaller())
	options = append(options, zap.AddStacktrace(zap.ErrorLevel))
	options = append(options, zap.ErrorOutput(errSyncer))

	if opt.Development {
		options = append(options, zap.Development())
	}

	log := zap.New(core, options...)
	return log
}

func NewSink(logFile string) (zapcore.WriteSyncer, zapcore.WriteSyncer) {
	if len(logFile) == 0 {
		return zapcore.AddSync(os.Stdout), zapcore.AddSync(os.Stderr)
	}
	ll := &lumberjack.Logger{
		Filename:   logFile,
		MaxSize:    24,
		MaxBackups: 3,
		MaxAge:     1,
		Compress:   true,
	}
	infoSink := zapcore.Lock(zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(ll)))
	errSink := zapcore.Lock(zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stderr), zapcore.AddSync(ll)))
	return infoSink, errSink
}
