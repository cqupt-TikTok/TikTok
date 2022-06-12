package log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	retalog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/logger"
	"math"
	"os"
	"time"
)

// ApiLogger 路由日志记录
func ApiLogger() gin.HandlerFunc {
	filePath := "log/apiLog/log."
	linkName := "apiLatest_log.log"
	src, err := os.OpenFile(filePath, os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("logger err:", err)
	}
	mylogger := logrus.New()
	mylogger.Out = src

	mylogger.SetLevel(logrus.DebugLevel)

	logWriter, _ := retalog.New(
		filePath+"%Y%m%d.log",
		retalog.WithMaxAge(7*24*time.Hour),
		retalog.WithRotationTime(24*time.Hour),
		retalog.WithLinkName(linkName),
	)

	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	mylogger.AddHook(Hook)
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds()/1000000.0))))
		hostName, err := os.Hostname()
		if err != nil {
			hostName = "unknown"
		}
		statusCode := c.Writer.Status()
		clientIp := c.ClientIP()
		userAgent := c.Request.UserAgent()
		dataSize := c.Writer.Size()
		if dataSize < 0 {
			dataSize = 0
		}
		method := c.Request.Method
		path := c.Request.RequestURI

		entry := mylogger.WithFields(logrus.Fields{
			"HostName":  hostName,
			"status":    statusCode,
			"SpendTime": spendTime,
			"Ip":        clientIp,
			"Method":    method,
			"Path":      path,
			"DataSize":  dataSize,
			"Agent":     userAgent,
		})
		if len(c.Errors) > 0 {
			entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}

type MyWriter struct {
	mlog *logrus.Logger
}

//实现gorm/logger.Writer接口
func (m *MyWriter) Printf(format string, v ...interface{}) {
	//利用loggus记录日志
	hostName, err := os.Hostname()
	if err != nil {
		hostName = "unknown"
	}
	m.mlog.Info("HostName: ", hostName, "  sql: ", v)
}

func NewMyWriter() *MyWriter {
	var err error
	var src *os.File
	log := logrus.New()
	filePath := "log/mysqlLog/log."
	linkName := "mysqlLatest_log.log"
	src, err = os.OpenFile(filePath, os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("logger err:", err)
	}
	log.Out = src
	log.SetLevel(logrus.InfoLevel)
	logWriter, _ := retalog.New(
		filePath+"%Y%m%d.log",
		retalog.WithMaxAge(7*24*time.Hour),
		retalog.WithRotationTime(24*time.Hour),
		retalog.WithLinkName(linkName),
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	log.AddHook(Hook)
	return &MyWriter{mlog: log}
}

// NewMysqlLogger mysql日志记录
func NewMysqlLogger() logger.Interface {
	newLogger := logger.New(
		NewMyWriter(),
		logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			Colorful:                  false,
			IgnoreRecordNotFoundError: false,
			LogLevel:                  logger.Info,
		})
	return newLogger
}
