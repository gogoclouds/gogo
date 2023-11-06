package logger

import (
	"context"
	"fmt"
	"github.com/natefinch/lumberjack"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

type SlogLogger struct {
	l *slog.Logger
}

func (sl *SlogLogger) Panic(msg string, keysAndValues ...any) {
	sl.l.Error(msg, keysAndValues...)
}

func (sl *SlogLogger) Panicf(format string, args ...any) {
	sl.l.Error(fmt.Sprintf(format, args...))
}

func (sl *SlogLogger) Panicc(ctx context.Context, msg string, keysAndValues ...any) {
	sl.l.ErrorContext(ctx, msg, keysAndValues...)
}

func (sl *SlogLogger) Paniccf(ctx context.Context, format string, args ...any) {
	sl.l.ErrorContext(ctx, fmt.Sprintf(format, args...))
}

func (sl *SlogLogger) Error(msg string, keysAndValues ...any) {
	sl.l.Error(msg, keysAndValues...)
}

func (sl *SlogLogger) Errorf(format string, args ...any) {
	sl.l.Error(fmt.Sprintf(format, args...))
}

func (sl *SlogLogger) Errorc(ctx context.Context, msg string, keysAndValues ...any) {
	sl.l.ErrorContext(ctx, msg, keysAndValues...)
}

func (sl *SlogLogger) Errorcf(ctx context.Context, format string, args ...any) {
	sl.l.ErrorContext(ctx, fmt.Sprintf(format, args...))
}

func (sl *SlogLogger) Info(msg string, keysAndValues ...any) {
	sl.l.Info(msg, keysAndValues...)
}

func (sl *SlogLogger) Infof(format string, args ...any) {
	sl.l.Info(fmt.Sprintf(format, args...))
}

func (sl *SlogLogger) Infoc(ctx context.Context, msg string, keysAndValues ...any) {
	sl.l.InfoContext(ctx, msg, keysAndValues...)
}

func (sl *SlogLogger) Infocf(ctx context.Context, format string, args ...any) {
	sl.l.InfoContext(ctx, fmt.Sprintf(format, args...))
}

func (sl *SlogLogger) Debug(msg string, keysAndValues ...any) {
	sl.l.Debug(msg, keysAndValues...)
}

func (sl *SlogLogger) Debugf(format string, args ...any) {
	sl.l.Debug(fmt.Sprintf(format, args...))
}

func (sl *SlogLogger) Debugc(ctx context.Context, msg string, keysAndValues ...any) {
	sl.l.DebugContext(ctx, msg, keysAndValues...)
}

func (sl *SlogLogger) Debugcf(ctx context.Context, format string, args ...any) {
	sl.l.DebugContext(ctx, fmt.Sprintf(format, args...))
}

func InitSlogLogger(conf Config) {
	lvl := &slog.LevelVar{}
	lvl.Set(slog.LevelDebug)

	writerFile := readFile(conf)
	logger := slog.New(slog.NewJSONHandler(io.MultiWriter(os.Stderr, writerFile), &slog.HandlerOptions{
		AddSource: true,
		Level:     lvl,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// set time format.
			if a.Key == slog.TimeKey && len(groups) == 0 {
				format := a.Value.Time().Format(time.DateTime + ".000")
				return slog.String(slog.TimeKey, format)
			}
			// Remove the directory from the source's filename.
			if a.Key == slog.SourceKey {
				source := a.Value.Any().(*slog.Source)
				source.File = filepath.Base(source.File)
			}
			return a
		},
	}))
	SetLogger(&SlogLogger{l: logger})
}

func readFile(conf Config) io.Writer {
	filename := conf.makeFilename()
	r := &lumberjack.Logger{
		Filename:   filename,
		LocalTime:  true,
		MaxSize:    int(conf.FileMaxSize),
		MaxAge:     int(conf.FileMaxAge),
		MaxBackups: 1,
		Compress:   conf.FileCompress,
	}
	return r
}
