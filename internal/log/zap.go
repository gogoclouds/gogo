package logger

import (
	"fmt"
	"github.com/gogoclouds/gogo/enum"
	"github.com/gogoclouds/gogo/internal/conf"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"time"
)

func L() *zap.SugaredLogger {
	return zap.S()
}

// 从配置文件映射结构

var (
	defLoggerPath           = "/logs"
	defLoggerFilenamePrefix = "log"
)

var atomicLevel zap.AtomicLevel

// Initialize logger handle
// appName log file prefix
func Initialize(appName string, conf conf.Log) {
	SetLevel(conf.Level)
	fileCore := zapcore.NewCore( // 输出到日志文件
		setJSONEncoder(),
		setLoggerWriter(appName, conf),
		atomicLevel,
	)

	consoleCore := zapcore.NewCore( // 输出到控制台
		setConsoleEncoder(),
		zapcore.Lock(os.Stdout),
		atomicLevel,
	)
	l := zap.New(zapcore.NewTee(fileCore, consoleCore), zap.AddCaller(), zap.AddCallerSkip(1))
	zap.ReplaceGlobals(l)
}

// SetLevel 动态更新限制日志打印级别
// zapcore.Levels ("debug", "info", "warn", "error", "dpanic", "panic", and "fatal").
func SetLevel(level enum.LoggerLevel) {
	if err := atomicLevel.UnmarshalText([]byte(level)); err != nil {
		zap.S().Error(err)
	}
}

func setConsoleEncoder() zapcore.Encoder {
	ec := setEncoderConf()
	ec.EncodeLevel = zapcore.CapitalColorLevelEncoder // 终端输出 日志级别有颜色
	return zapcore.NewConsoleEncoder(ec)
}

func setJSONEncoder() zapcore.Encoder {
	ec := setEncoderConf()
	ec.EncodeLevel = zapcore.CapitalLevelEncoder // eg: info -> INFO
	return zapcore.NewConsoleEncoder(ec)
}

func setEncoderConf() zapcore.EncoderConfig {
	ec := zap.NewProductionEncoderConfig()
	ec.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000") // 转换编码的时间戳
	return ec
}

func setLoggerWriter(appName string, conf conf.Log) zapcore.WriteSyncer {
	fName := makeFilename(conf.DirPath, appName)
	return zapcore.AddSync(
		&lumberjack.Logger{
			Filename:   fName,                 // 要写入的日志文件
			MaxSize:    int(conf.FileSizeMax), // 日志文件的大小（M）
			MaxBackups: 1,                     // 备份数量
			MaxAge:     int(conf.FileAgeMax),  // 存留天数
			Compress:   true,                  // 压缩
			LocalTime:  true,                  // 默认 UTC 时间
		})
}

// xxx/logs/xxx-service-2006-01-01-150405.log
func makeFilename(path, name string) string {
	if path = strings.TrimSpace(path); path == "" {
		path = defLoggerPath
	}
	if path = strings.TrimSpace(path); name == "" {
		name = defLoggerFilenamePrefix
	}
	nowTime := time.Now().Format("2006-01-02-150405")
	return fmt.Sprintf("%s%s-%s.log", path+string(os.PathSeparator), name, nowTime)
}