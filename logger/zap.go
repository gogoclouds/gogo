package logger

import (
	"context"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

type ZapLogger struct {
	logger *zap.Logger
}

func (zl *ZapLogger) Panic(args ...any) {
	zl.logger.Sugar().Panic(args...)
}

func (zl *ZapLogger) Panicf(format string, args ...any) {
	zl.logger.Sugar().Panicf(format, args...)
}

func (zl *ZapLogger) Panicw(msg string, keysAndValues ...any) {
	zl.logger.Sugar().Panicw(msg, keysAndValues...)
}

func (zl *ZapLogger) Panicc(ctx context.Context, msg string) {
	zl.logger.Sugar().Panic(msg)

}

func (zl *ZapLogger) Paniccf(ctx context.Context, format string, args ...any) {
	zl.logger.Sugar().Panicf(format, args...)
}

func (zl *ZapLogger) Paniccw(ctx context.Context, msg string, keysAndValues ...any) {
	zl.logger.Sugar().Panicw(msg, keysAndValues...)
}

func (zl *ZapLogger) Error(args ...any) {
	zl.logger.Sugar().Error(args)
}

func (zl *ZapLogger) Errorf(format string, args ...any) {
	zl.logger.Sugar().Errorf(format, args...)
}

func (zl *ZapLogger) Errorw(msg string, keysAndValues ...any) {
	zl.logger.Sugar().Errorw(msg, keysAndValues...)
}

func (zl *ZapLogger) Errorc(ctx context.Context, msg string) {
	zl.logger.Sugar().Error(msg)
}

func (zl *ZapLogger) Errorcf(ctx context.Context, format string, args ...any) {
	zl.logger.Sugar().Errorf(format, args...)
}

func (zl *ZapLogger) Errorcw(ctx context.Context, msg string, keysAndValues ...any) {
	zl.logger.Sugar().Errorw(msg, keysAndValues...)
}

func (zl *ZapLogger) Info(args ...any) {
	zl.logger.Sugar().Info(args)
}

func (zl *ZapLogger) Infof(format string, args ...any) {
	zl.logger.Sugar().Infof(format, args...)
}

func (zl *ZapLogger) Infow(msg string, keysAndValues ...any) {
	zl.logger.Sugar().Infow(msg, keysAndValues...)
}

func (zl *ZapLogger) Infoc(ctx context.Context, msg string) {
	zl.Info(msg)
}

func (zl *ZapLogger) Infocf(ctx context.Context, format string, args ...any) {
	zl.Infof(format, args...)
}

func (zl *ZapLogger) Infocw(ctx context.Context, msg string, keysAndValues ...any) {
	zl.Infow(msg, keysAndValues...)
}

func (zl *ZapLogger) Debug(args ...any) {
	zl.logger.Sugar().Debug(args)
}

func (zl *ZapLogger) Debugf(format string, args ...any) {
	zl.logger.Sugar().Debugf(format, args...)
}

func (zl *ZapLogger) Debugw(msg string, keysAndValues ...any) {
	zl.logger.Sugar().Debugw(msg, keysAndValues...)
}

func (zl *ZapLogger) Debugc(ctx context.Context, msg string) {
	zl.logger.Sugar().Debug(msg)
}

func (zl *ZapLogger) Debugcf(ctx context.Context, format string, args ...any) {
	zl.logger.Sugar().Debugf(format, args...)
}

func (zl *ZapLogger) Debugcw(ctx context.Context, msg string, keysAndValues ...any) {
	zl.logger.Sugar().Debugw(msg, keysAndValues...)
}

// atomicLevel 动态更新限制日志打印级别
var atomicLevel zap.AtomicLevel

func InitZapLogger(conf Config) {
	atomicLevel = zap.NewAtomicLevel()
	go func() {
		for level := range conf.LevelCh {
			_ = atomicLevel.UnmarshalText([]byte(level))
		}
	}()
	fileCore := zapcore.NewCore( // 输出到日志文件
		setJSONEncoder(conf.TimeFormat, conf.FileJsonEncoder),
		setLoggerWriter(conf),
		atomicLevel,
	)
	consoleCore := zapcore.NewCore( // 输出到控制台
		setConsoleEncoder(conf.TimeFormat),
		zapcore.Lock(os.Stdout),
		atomicLevel,
	)
	l := zap.New(zapcore.NewTee(fileCore, consoleCore), zap.AddCaller(), zap.AddCallerSkip(2))
	SetLogger(&ZapLogger{logger: l})
}

func setConsoleEncoder(timeFormat string) zapcore.Encoder {
	ec := setEncoderConf(timeFormat)
	ec.EncodeLevel = zapcore.CapitalColorLevelEncoder // 终端输出 日志级别有颜色
	return zapcore.NewConsoleEncoder(ec)
}

func setJSONEncoder(timeFormat string, isJsonEncoder bool) zapcore.Encoder {
	ec := setEncoderConf(timeFormat)
	ec.EncodeLevel = zapcore.CapitalLevelEncoder // eg: info -> INFO
	if isJsonEncoder {
		return zapcore.NewJSONEncoder(ec)
	}
	return zapcore.NewConsoleEncoder(ec)
}

func setEncoderConf(timeFmt string) zapcore.EncoderConfig {
	ec := zap.NewProductionEncoderConfig()
	ec.EncodeTime = zapcore.TimeEncoderOfLayout(timeFmt) // 转换编码的时间戳
	return ec
}

func setLoggerWriter(conf Config) zapcore.WriteSyncer {
	fName := conf.makeFilename()
	return zapcore.AddSync(
		&lumberjack.Logger{
			Filename:   fName,                 // 要写入的日志文件
			MaxSize:    int(conf.FileMaxSize), // 日志文件的大小（M）
			MaxAge:     int(conf.FileMaxAge),  // 存留天数
			MaxBackups: 1,                     // 备份数量
			Compress:   conf.FileCompress,     // 压缩
			LocalTime:  true,                  // 默认 UTC 时间
		})
}
