package logger_test

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"pqredis/config"
	"pqredis/monitoring/logger"
	"testing"
)

func TestNew(t *testing.T) {
	l := logger.New(config.New())
	assert.NotNil(t, l)
}

func TestNoPanic(t *testing.T) {
	assert.NotPanics(t, func() {
		l := logger.New(config.New())
		l.ErrorWithTag(errors.New("foo"), logger.Fields{"bar": "baz"})
		l.ErrorWithTag(errors.New("foo"), nil)
		l.ErrorWithTag(nil, logger.Fields{"bar": "baz"})

		l.InfoWithTag("msg", logger.Fields{"bar": "baz"})
		l.InfoWithTag("msg", nil)

		l.DebugWithTag("msg", logger.Fields{"bar": "baz"})
		l.DebugWithTag("msg", nil)

		l.Error(errors.New("foo"))
		l.Infof("foo %s", "some info")
		l.Debugf("foo %s", "some debug info")
		l.Warnf("foo %s", "some warn info")
		l.Errorf("foo %s", "some error")
	})
}

func TestWhenPanicfShouldPanic(t *testing.T) {
	assert.Panics(t, func() {
		l := logger.New(config.New())
		l.Panicf("foo %s", "some fatal error")
	})
}

