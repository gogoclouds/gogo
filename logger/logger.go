package logger

import (
	"context"
)

type ILogger interface {
	Panic(args ...any)
	Panicf(format string, args ...any)
	Panicw(msg string, keysAndValues ...any)
	Panicc(ctx context.Context, msg string)
	Paniccf(ctx context.Context, format string, args ...any)
	Paniccw(ctx context.Context, msg string, keysAndValues ...any)

	Error(args ...any)
	Errorf(format string, args ...any)
	Errorw(msg string, keysAndValues ...any)
	Errorc(ctx context.Context, msg string)
	Errorcf(ctx context.Context, format string, args ...any)
	Errorcw(ctx context.Context, msg string, keysAndValues ...any)

	Info(args ...any)
	Infof(format string, args ...any)
	Infow(msg string, keysAndValues ...any)
	Infoc(ctx context.Context, msg string)
	Infocf(ctx context.Context, format string, args ...any)
	Infocw(ctx context.Context, msg string, keysAndValues ...any)

	Debug(args ...any)
	Debugf(format string, args ...any)
	Debugw(msg string, keysAndValues ...any)
	Debugc(ctx context.Context, msg string)
	Debugcf(ctx context.Context, format string, args ...any)
	Debugcw(ctx context.Context, msg string, keysAndValues ...any)
}

var log ILogger

func SetLogger(l ILogger) {
	log = l
}

func Panic(args ...any) {
	log.Panic(args...)
}

func Panicf(template string, args ...any) {
	log.Panicf(template, args...)
}

func Panicw(msg string, keysAndValues ...any) {
	log.Panicw(msg, keysAndValues...)
}

func Error(args ...any) {
	log.Error(args...)
}

func Errorf(template string, args ...any) {
	log.Errorf(template, args...)
}

func Errorw(msg string, keysAndValues ...any) {
	log.Errorw(msg, keysAndValues...)
}

func Info(args ...any) {
	log.Info(args...)
}

func Infof(template string, args ...any) {
	log.Infof(template, args...)
}

func Infow(msg string, keysAndValues ...any) {
	log.Infow(msg, keysAndValues...)
}

func Debug(args ...any) {
	log.Debug(args...)
}

func Debugf(template string, args ...any) {
	log.Debugf(template, args...)
}

func Debugw(msg string, keysAndValues ...any) {
	log.Debugw(msg, keysAndValues...)
}
