package service

import "github.com/sirupsen/logrus"

type Logger interface {
	Warn(string)
	Debug(string)
	Info(string)
	Error(string)
}

// LoggerImpl is a wrapper struct around the logrus logger client
type LoggerImpl struct {
	logger *logrus.Logger
}

func NewLogger() Logger {
	logClient := logrus.New()
	logClient.SetFormatter(&logrus.JSONFormatter{})

	return &LoggerImpl{
		logger: logClient,
	}
}

func (l *LoggerImpl) Warn(msg string) {
	l.logger.Warn(msg)
}

func (l *LoggerImpl) Debug(msg string) {
	l.logger.Debug(msg)
}

func (l *LoggerImpl) Info(msg string) {
	l.logger.Info(msg)
}

func (l *LoggerImpl) Error(msg string) {
	l.logger.Error(msg)
}
