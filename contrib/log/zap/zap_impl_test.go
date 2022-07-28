package zap

import (
	"testing"

	"github.com/LinkinStars/go-scaffold/logger"
)

func TestLogger(t *testing.T) {
	logger.SetLogger(NewLogger(WithName("go-scaffold"), WithCallerFullPath()))
	logger.Debug("debug")
	logger.Debugf("debug")
	logger.Info("info")
	logger.Infof("info")
	logger.Warn("warn")
	logger.Warnf("warn")
	logger.Error("error")
	logger.Errorf("error")
}
