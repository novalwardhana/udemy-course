package belajar_golang_logging

import (
	"testing"

	"github.com/sirupsen/logrus"
)

func TestField(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.WithField("user", "noval").Info("Hello world")
	logger.WithField("username", "novalwardhana").WithField("email", "noval@gmail.com").Info("Hello world")

	logger.Info("Hello")
}

func TestFields(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.WithFields(logrus.Fields{
		"username": "novalwardhana",
		"city":     "Bantul",
	}).Info("Hello world")
}
