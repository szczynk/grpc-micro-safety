package logger

import (
	"os"
	"time"
	"user/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Logger methods interface
type Logger interface {
	InitLogger()
	ZapLogger() *zap.Logger
	Sync() error
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	WarnErrMsg(msg string, err error)
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	Err(msg string, err error)
	DPanic(args ...interface{})
	DPanicf(template string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
	Named(name string)
	GrpcMiddlewareAccessLogger(method string, time time.Duration, metadata map[string][]string, err error)
	GrpcClientInterceptorLogger(method string, req interface{}, reply interface{}, time time.Duration, metaData map[string][]string, err error)
	// MessageProducer(ctx context.Context, msg string, level zapcore.Level, code codes.Code, err error, duration zapcore.Field)
}

// Logger
type apiLogger struct {
	cfg         *config.Config
	logger      *zap.Logger
	sugarLogger *zap.SugaredLogger
}

// App Logger constructor
func NewAPILogger(cfg *config.Config) Logger {
	return &apiLogger{cfg: cfg}
}

// For mapping config logger to app logger levels
var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func (l *apiLogger) getLoggerLevel(cfg *config.Config) zapcore.Level {
	level, exist := loggerLevelMap[cfg.Logger.Level]
	if !exist {
		return zapcore.DebugLevel
	}

	return level
}

// Init logger
func (l *apiLogger) InitLogger() {
	logLevel := l.getLoggerLevel(l.cfg)

	logWriter := zapcore.AddSync(os.Stderr)

	var encoderCfg zapcore.EncoderConfig
	if l.cfg.Server.Mode == "development" {
		encoderCfg = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderCfg = zap.NewProductionEncoderConfig()
	}

	var encoder zapcore.Encoder
	encoderCfg.LevelKey = "LEVEL"
	encoderCfg.CallerKey = "CALLER"
	encoderCfg.TimeKey = "TIME"
	encoderCfg.NameKey = "NAME"
	encoderCfg.MessageKey = "MESSAGE"

	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderCfg.EncodeCaller = zapcore.ShortCallerEncoder
	encoderCfg.EncodeDuration = zapcore.StringDurationEncoder

	if l.cfg.Logger.Encoding == "console" {
		encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
		// encoderCfg.EncodeCaller = zapcore.FullCallerEncoder
		encoderCfg.ConsoleSeparator = " | "
		encoder = zapcore.NewConsoleEncoder(encoderCfg)
	} else {
		encoderCfg.FunctionKey = "FUNC"
		encoderCfg.EncodeName = zapcore.FullNameEncoder
		encoder = zapcore.NewJSONEncoder(encoderCfg)
	}

	core := zapcore.NewCore(encoder, logWriter, zap.NewAtomicLevelAt(logLevel))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	l.logger = logger
	l.sugarLogger = logger.Sugar()
	if err := l.sugarLogger.Sync(); err != nil {
		l.sugarLogger.Error(err)
	}
}

// Logger methods

func (l *apiLogger) ZapLogger() *zap.Logger {
	return l.logger
}

// Sync flushes any buffered log entries
func (l *apiLogger) Sync() error {
	go l.logger.Sync() // nolint: errcheck
	return l.sugarLogger.Sync()
}

func (l *apiLogger) Debug(args ...interface{}) {
	l.sugarLogger.Debug(args...)
}

func (l *apiLogger) Debugf(template string, args ...interface{}) {
	l.sugarLogger.Debugf(template, args...)
}

func (l *apiLogger) Info(args ...interface{}) {
	l.sugarLogger.Info(args...)
}

func (l *apiLogger) Infof(template string, args ...interface{}) {
	l.sugarLogger.Infof(template, args...)
}

func (l *apiLogger) Warn(args ...interface{}) {
	l.sugarLogger.Warn(args...)
}

func (l *apiLogger) Warnf(template string, args ...interface{}) {
	l.sugarLogger.Warnf(template, args...)
}

func (l *apiLogger) WarnErrMsg(msg string, err error) {
	l.logger.Warn(msg, zap.String("error", err.Error()))
}

func (l *apiLogger) Error(args ...interface{}) {
	l.sugarLogger.Error(args...)
}

func (l *apiLogger) Errorf(template string, args ...interface{}) {
	l.sugarLogger.Errorf(template, args...)
}

func (l *apiLogger) Err(msg string, err error) {
	l.logger.Error(msg, zap.Error(err))
}

func (l *apiLogger) DPanic(args ...interface{}) {
	l.sugarLogger.DPanic(args...)
}

func (l *apiLogger) DPanicf(template string, args ...interface{}) {
	l.sugarLogger.DPanicf(template, args...)
}

func (l *apiLogger) Panic(args ...interface{}) {
	l.sugarLogger.Panic(args...)
}

func (l *apiLogger) Panicf(template string, args ...interface{}) {
	l.sugarLogger.Panicf(template, args...)
}

func (l *apiLogger) Fatal(args ...interface{}) {
	l.sugarLogger.Fatal(args...)
}

func (l *apiLogger) Fatalf(template string, args ...interface{}) {
	l.sugarLogger.Fatalf(template, args...)
}

// Named add logger microservice name
func (l *apiLogger) Named(name string) {
	l.logger = l.logger.Named(name)
	l.sugarLogger = l.sugarLogger.Named(name)
}

func (l *apiLogger) GrpcMiddlewareAccessLogger(method string, time time.Duration, metadata map[string][]string, err error) {
	l.logger.Info(
		"GRPC",
		zap.String("method", method),
		zap.Duration("time", time),
		zap.Any("metadata", metadata),
		zap.String("error", err.Error()),
	)
}

func (l *apiLogger) GrpcClientInterceptorLogger(method string, req, reply interface{}, time time.Duration, metadata map[string][]string, err error) {
	l.logger.Info(
		"GRPC",
		zap.String("method", method),
		zap.Any("req", req),
		zap.Any("reply", reply),
		zap.Duration("time", time),
		zap.Any("metadata", metadata),
		zap.String("error", err.Error()),
	)
}

// func (l *apiLogger) MessageProducer(ctx context.Context, msg string, level zapcore.Level, code codes.Code, err error, duration zapcore.Field) {
// 	// re-extract logger from newCtx, as it may have extra fields that changed in the holder.
// 	ctxzap.Extract(ctx).Check(level, msg).Write(
// 		zap.String("error", err.Error()),
// 		zap.String("grpc.code", code.String()),
// 		duration,
// 	)
// }
