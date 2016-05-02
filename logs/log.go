package logs

import (
	"github.com/astaxie/beego/logs"
)

var Log *logs.BeeLogger

func InitLogger(output string, arg string) {
	Log = logs.NewLogger(1000)
	Log.SetLogger(output, arg)
}

func GetLogger() *logs.BeeLogger {
	return Log
}

// Debug Log DEBUG level message.
func Debug(format string, v ...interface{}) {
	Log.Debug(format, v)
}

// Warn Log WARN level message.
// compatibility alias for Warning()
func Warn(format string, v ...interface{}) {
	Log.Warn(format, v)
}

// Info Log INFO level message.
// compatibility alias for Informational()
func Info(format string, v ...interface{}) {
	Log.Info(format, v)
}

// Trace Log TRACE level message.
// compatibility alias for Debug()
func Trace(format string, v ...interface{}) {
	Log.Trace(format, v)
}

// Error Log ERROR level message.
func Error(format string, v ...interface{}) {
	Log.Error(format, v)
}
