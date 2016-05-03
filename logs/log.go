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
	if len(v) > 0 {
		Log.Debug(format, v)
	} else {
		Log.Debug(format)
	}
}

// Warn Log WARN level message.
// compatibility alias for Warning()
func Warn(format string, v ...interface{}) {
	if len(v) > 0 {
		Log.Warn(format, v)
	} else {
		Log.Warn(format)
	}
}

// Info Log INFO level message.
// compatibility alias for Informational()
func Info(format string, v ...interface{}) {
	if len(v) > 0 {
		Log.Info(format, v)
	} else {
		Log.Info(format)
	}
}

// Trace Log TRACE level message.
// compatibility alias for Debug()
func Trace(format string, v ...interface{}) {
	if len(v) > 0 {
		Log.Trace(format, v)
	} else {
		Log.Trace(format)
	}
}

// Error Log ERROR level message.
func Error(format string, v ...interface{}) {
	if len(v) > 0 {
		Log.Error(format, v)
	} else {
		Log.Error(format)
	}
}
