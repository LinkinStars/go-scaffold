package zap

import (
	"time"

	"github.com/LinkinStars/go-scaffold/logger"
	"go.uber.org/zap"
)

var _ logger.Logger = (*Logger)(nil)

// Logger zap logger config
type Logger struct {
	conf LoggerConfig
	log  *zap.Logger
	slog *zap.SugaredLogger
}

type LoggerConfig struct {
	// log level
	level logger.Level
	// log file name
	name string
	// log file path, if it is empty no output file
	path string
	// if stdout is true will output stdout
	stdout bool
	// file max save duration, default is 7 days
	maxAge time.Duration
	// file rotation time, default is 1 file per day
	rotationTime time.Duration
	// if callerFullPath is true will output caller fullpath
	callerFullPath bool
}

// NewLogger new zap logger
func NewLogger(level logger.Level, options ...LogOption) *Logger {
	l := &Logger{
		conf: LoggerConfig{
			level:        level,
			name:         "log",
			stdout:       true,
			maxAge:       7 * 24 * time.Hour,
			rotationTime: 24 * time.Hour,
		},
	}
	for _, option := range options {
		option(l)
	}
	l.log = InitZap(l.conf)
	l.slog = l.log.Sugar()
	return l
}

// Debug log
func (z *Logger) Debug(v ...interface{}) {
	if z.conf.level <= logger.LevelDebug {
		z.slog.Debug(v...)
	}
}

// Debugf log
func (z *Logger) Debugf(format string, v ...interface{}) {
	if z.conf.level <= logger.LevelDebug {
		z.slog.Debugf(format, v...)
	}
}

// Info log
func (z *Logger) Info(v ...interface{}) {
	if z.conf.level <= logger.LevelInfo {
		z.slog.Info(v...)
	}
}

// Infof log
func (z *Logger) Infof(format string, v ...interface{}) {
	if z.conf.level <= logger.LevelInfo {
		z.slog.Infof(format, v...)
	}
}

// Warn log
func (z *Logger) Warn(v ...interface{}) {
	if z.conf.level <= logger.LevelWarn {
		z.slog.Warn(v...)
	}
}

// Warnf log
func (z *Logger) Warnf(format string, v ...interface{}) {
	if z.conf.level <= logger.LevelWarn {
		z.slog.Warnf(format, v...)
	}
}

// Error log
func (z *Logger) Error(v ...interface{}) {
	if z.conf.level <= logger.LevelError {
		z.slog.Error(v...)
	}
}

// Errorf log
func (z *Logger) Errorf(format string, v ...interface{}) {
	if z.conf.level <= logger.LevelError {
		z.slog.Errorf(format, v...)
	}
}
