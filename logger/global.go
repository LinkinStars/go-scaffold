package logger

// global
var global Logger

func init() {
	global = DefaultLogger
}

// SetLogger set logger
func SetLogger(log Logger) {
	global = log
}

// GetLogger get logger
func GetLogger() Logger {
	return global
}

// Debug log
func Debug(v ...interface{}) {
	global.Debug(v...)
}

// Debugf log
func Debugf(format string, v ...interface{}) {
	global.Debugf(format, v...)
}

// Info log
func Info(v ...interface{}) {
	global.Info(v...)
}

// Infof log
func Infof(format string, v ...interface{}) {
	global.Infof(format, v...)
}

// Warn log
func Warn(v ...interface{}) {
	global.Warn(v...)
}

// Warnf log
func Warnf(format string, v ...interface{}) {
	global.Warnf(format, v...)
}

// Error log
func Error(v ...interface{}) {
	global.Error(v...)
}

// Errorf log
func Errorf(format string, v ...interface{}) {
	global.Errorf(format, v...)
}
