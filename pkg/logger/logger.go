package logger

import (
	"fmt"
	"github.com/newrelic/go-agent/v3/integrations/nrpkgerrors"
	"github.com/newrelic/go-agent/v3/newrelic"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"os"
)

type Logger interface {
	Debug(args interface{})
	Info(args interface{})
	Warn(args interface{})
	Error(args interface{})
	Fatal(args interface{})
}

type apiLogger struct {
	log *logrus.Logger
	app *newrelic.Application
}

func NewApiLogger(app *newrelic.Application) *apiLogger {
	var l = logrus.New()
	l.Out = os.Stdout
	l.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	return &apiLogger{
		log: l,
		app: app,
	}
}

func (l *apiLogger) Debug(args interface{}) {
	l.log.WithFields(logrus.Fields{"event": "general"}).Debug(args)
}

func (l *apiLogger) Info(args interface{}) {
	l.log.WithFields(logrus.Fields{"event": "general"}).Info(args)
}

func (l *apiLogger) Warn(args interface{}) {
	l.log.WithFields(logrus.Fields{"event": "general"}).Warn(args)
}

func (l *apiLogger) Error(args interface{}) {
	txn := l.app.StartTransaction("log")
	defer txn.End()
	e := errors.New(fmt.Sprint(args))
	txn.NoticeError(nrpkgerrors.Wrap(e))

	l.log.WithFields(logrus.Fields{"event": "general"}).Error(args)
}

func (l *apiLogger) Fatal(args interface{}) {
	txn := l.app.StartTransaction("log")
	defer txn.End()
	e := errors.New(fmt.Sprint(args))
	txn.NoticeError(nrpkgerrors.Wrap(e))

	l.log.WithFields(logrus.Fields{"event": "general"}).Fatal(args)
}
