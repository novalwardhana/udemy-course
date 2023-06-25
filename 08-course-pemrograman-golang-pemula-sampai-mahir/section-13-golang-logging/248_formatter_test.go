package belajar_golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestFormatter(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.SetLevel(logrus.TraceLevel)
	logger.Trace("Ini logger trace")
	logger.Debug("Ini logger debug")
	logger.Info("Ini logger info")
}
