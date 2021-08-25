package log

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// logger
type LoggerEntry struct {
	IsDevelopment bool   // is development
	SystemName    string // system name
	FilePath      string // log file path
	NameSpace     string // name space
}

type Entry struct {
	*LoggerEntry
	*logrus.Entry
}

// log
func log(level string, format string, args ...interface{}) {
	file, line, function := getCaller(CALLDEPTH)
	fmts := fmt.Sprintf("%s:%d %s level:%s message: %s", file, line, function, level, format)
	message := fmt.Sprintf(fmts, args...)
	// echo log
	Logger.Entry.WithFields(
		logrus.Fields{
			"@timestamp": time.Now().Format("2006-01-02 15:04:05"),
			"@fields": map[string]interface{}{
				"result": message,
			},
		}).Warningln()
}

// Export
//
var Logger *Entry

// RegisterLogger
func RegisterLogger(loggerEntry *LoggerEntry) {
	if loggerEntry == nil {
		panic("new logger err,logger entry empty nil")
	}
	// logrus set log config
	if loggerEntry.IsDevelopment {
		logrus.SetLevel(logrus.DebugLevel)
		logrus.SetOutput(os.Stderr)
		logrus.SetFormatter(&logrus.TextFormatter{})
	} else {
		logFile, err := os.OpenFile(loggerEntry.FilePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModePerm)
		if err != nil {
			logrus.Printf("logger file stream closed.")
			logFile.Close()
		}
		logrus.SetLevel(logrus.WarnLevel)
		logrus.SetOutput(logFile)
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
	// logrus entry
	source := loggerEntry.SystemName
	if source == "" {
		source, _ = os.Hostname()
	}
	Logger = &Entry{
		LoggerEntry: loggerEntry,
		Entry: logrus.WithFields(logrus.Fields{
			"@source":    source,
			"@namespace": loggerEntry.NameSpace,
			"@service":   loggerEntry.SystemName,
			"@hostIp":    "",
			"@podIp":     getLocalIP(),
		}),
	}
	return
}

// LOACCESS access log
func LOGACCESS(format string, args ...interface{}) {
	if Logger == nil || Logger.LoggerEntry == nil {
		panic("v")
		return
	}
	level := fmt.Sprintf("%s_ACCESS", Logger.SystemName)
	log(strings.ToUpper(level), format, args...)
}

// LOGE error log
func LOGERROR(format string, args ...interface{}) {
	if Logger == nil || Logger.LoggerEntry == nil {
		return
	}
	level := fmt.Sprintf("%s_ERROR", Logger.SystemName)
	log(strings.ToUpper(level), format, args...)
}

// LOGFLOW flow log
func LOGFLOW(format string, args ...interface{}) {
	if Logger == nil || Logger.LoggerEntry == nil {
		return
	}
	level := fmt.Sprintf("%s_FLOW", Logger.SystemName)
	log(strings.ToUpper(level), format, args...)
}

// LOGD debug log
func LOGDEBUG(format string, args ...interface{}) {
	if Logger == nil || Logger.LoggerEntry == nil {
		return
	}
	level := fmt.Sprintf("%s_DEBUG", Logger.SystemName)
	log(strings.ToUpper(level), format, args...)
}

func LOGSQL(format string, args ...interface{}) {
	if Logger == nil || Logger.LoggerEntry == nil {
		return
	}
	level := fmt.Sprintf("%s_SQL", Logger.SystemName)
	log(strings.ToUpper(level), format, args...)
}

// PrintFunc
type PrintFunc func(routApi string, t time.Time, err *error, in interface{}, out interface{})

// PrintLog
func PrintLog(routApi string, t time.Time, err *error, in interface{}, out interface{}) {
	LOGACCESS("routApi(%v),err(%v),latency(%v),req(%v),rep(%v)", routApi, *err, time.Now().Sub(t),
		structToString(in), structToString(out))
}

// PrintGrpcLog,grpc log print
func PrintGrpcLog(grpcInterface string, t time.Time, err *error, in interface{}, out interface{}) {
	LOGFLOW("grpc interface(%v),err(%v),latency(%v),req(%v),rep(%v)", grpcInterface, *err, time.Now().Sub(t), structToString(in), structToString(out))
}
