package logger

import (
	"os"
	"time"
	"wecat/common/setting"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var log *zap.SugaredLogger

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpainc": zapcore.DPanicLevel,
	"painc":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func Setup(s *setting.LogSettingS) {
	var syncWriters []zapcore.WriteSyncer
	level := getLoggerLevel(s.Level)

	fileConfig := &lumberjack.Logger{
		Filename:   s.Path,
		MaxSize:    s.MaxSize,
		MaxAge:     s.MaxAge,
		MaxBackups: s.MaxBackups,
		LocalTime:  s.LocalTime,
		Compress:   s.Compress,
	}
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
		pae.AppendString(t.Format("2016-01-02 15:04:05.000000"))
	}

	if s.ConsoleStdout {
		syncWriters = append(syncWriters, zapcore.AddSync(os.Stdout))
	}

	if s.FileStdout {
		syncWriters = append(syncWriters, zapcore.AddSync(fileConfig))
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoder),
		zapcore.NewMultiWriteSyncer(syncWriters...),
		zap.NewAtomicLevelAt(level))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	log = logger.Sugar()
}

func getLoggerLevel(level string) zapcore.Level {
	if le, ok := levelMap[level]; ok {
		return le
	}
	return zapcore.InfoLevel
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func DPanic(args ...interface{}) {
	log.DPanic(args...)
}

func DPanicf(format string, args ...interface{}) {
	log.DPanicf(format, args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}
