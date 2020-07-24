package logger

import (
	"os"
	"pqredis/config"

	"github.com/sirupsen/logrus"
)

// Fields represents key-value pairs and can be used to
// provide additional context in logs
type Fields map[string]interface{}

// Logger represents a generic logging component
type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
	Error(err error, args ...interface{})
	ErrorWithTag(err error, fields Fields)
	InfoWithTag(msg string, fields Fields)
	DebugWithTag(msg string, fields Fields)
}

type logger struct {
	*logrus.Logger
}

type Error struct {
	Error error
}

// Setup the logger with appropriate log-level and format.
// Valid log-levels are debug, info, warn, error, fatal, panic.
// Valid log-formats are plain or json (default: json)

func New(cfg config.Config) Logger {
	level, err := logrus.ParseLevel(cfg.LogLevel())
	if err != nil {
		level = logrus.WarnLevel
	}

	formatter := &logrus.JSONFormatter{
		// put context in the behind by prepending x
		DataKey: "xcontext",
	}

	log := &logger{
		Logger: &logrus.Logger{
			Out:       os.Stderr,
			Hooks:     make(logrus.LevelHooks),
			Level:     level,
			Formatter: formatter,
		},
	}

	if cfg.LogFormat() != "json" {
		log.Formatter = &logrus.TextFormatter{}
	}

	return log
}

func (log *logger) Error(err error, args ...interface{}) {
	log.Logger.Error(err, args)
}

func (log *logger) ErrorWithTag(err error, fields Fields) {
	if err != nil {
		log.WithFields(logrus.Fields(fields)).Error(err.Error())
	}
}

func (log *logger) InfoWithTag(msg string, fields Fields) {
	log.WithFields(logrus.Fields(fields)).Info(msg)
}

func (log *logger) DebugWithTag(msg string, fields Fields) {
	log.WithFields(logrus.Fields(fields)).Debug(msg)
}

func (log *logger) Debugf(format string, args ...interface{}) {
	log.Logger.Debugf(format, args...)
}

func (log *logger) Infof(format string, args ...interface{}) {
	log.Logger.Infof(format, args...)
}

func (log *logger) Warnf(format string, args ...interface{}) {
	log.Logger.Warnf(format, args...)
}

func (log *logger) Errorf(format string, args ...interface{}) {
	log.Logger.Errorf(format, args...)
}

func (log *logger) Fatalf(format string, args ...interface{}) {
	log.Logger.Fatalf(format, args...)
}

func (log *logger) Panicf(format string, args ...interface{}) {
	log.Logger.Panicf(format, args...)
}

