package logger_test

import (
	"github.com/gogoclouds/gogo/enum"
	"github.com/gogoclouds/gogo/logger"
	"runtime"
	"testing"
)

var logConf logger.Config

func init() {
	logConf = logger.NewConfig(
		logger.WithLevel(enum.LoggerLevel_Info),
		logger.WithTimeFormat("2006-01-02 15:04:05.000"),
		logger.WithFilename("server"),
		logger.WithFileMaxSize(10),
		logger.WithFileMaxAge(6*30),
		logger.WithFileCompress(true),
	)
	logger.InitZapLogger(logConf)
}

func TestLogger(t *testing.T) {
	logger.Debug("The is ", "Debug")
	logger.Info("The is ", "Info")
	logger.Error("The is ", "Error")

	go logConf.SetLevel(enum.LoggerLevel_Debug)
	runtime.Gosched()
	logger.Debugf("The is %s", "Debugf")
	logger.Infof("The is %s", "Info")
	logger.Errorf("The is %s", "Errorf")

	logger.Debugw("The is Debugw", "username", "lanjin.wei")
	logger.Infow("The is Infow", "age", 28)
	logger.Errorw("The is Errorw", "city", "shenzhen")

	// output:
	// 2023-06-12 21:32:35.931	DEBUG	logger/logger_test.go:21	The is Debug
	// 2023-06-12 21:32:35.932	INFO	logger/logger_test.go:22	The is Info
	// 2023-06-12 21:32:35.932	ERROR	logger/logger_test.go:23	The is Error
	// 2023-06-12 21:32:35.932	DEBUG	logger/logger_test.go:25	The is Debugf
	// 2023-06-12 21:32:35.932	INFO	logger/logger_test.go:26	The is Info
	// 2023-06-12 21:32:35.932	ERROR	logger/logger_test.go:27	The is Errorf
	// 2023-06-12 21:32:35.932	DEBUG	logger/logger_test.go:29	The is Debug	{"username": "lanjin.wei"}
	// 2023-06-12 21:32:35.932	INFO	logger/logger_test.go:30	The is Info	{"age": 28}
	// 2023-06-12 21:32:35.933	ERROR	logger/logger_test.go:31	The is Error	{"city": "shenzhen"}
}
