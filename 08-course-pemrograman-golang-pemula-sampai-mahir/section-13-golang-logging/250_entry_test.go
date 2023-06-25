package belajar_golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestEntry(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.WithField("username", "novalwardhana").Info("Hello logging")

	entry := logrus.NewEntry(logger)
	entry.WithField("email", "noval@gmail.com")
	entry.Info("Hello entry")
}
