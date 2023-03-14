package logs

import (
	timeutil "go-skeleton/pkg/utils/time"
	"io"
	"os"
	"reflect"

	"github.com/go-resty/resty/v2"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

// Logrus : implement Logger
type Logrus struct {
	*logrus.Logger
}

var (
	Log           *Logrus
	serverlessURL string
	client        *resty.Client
)

const (
	SetRoutePath = "ROUTE-PATH"
)

type (
	Fields = logrus.Fields
)

// Init init log wrapper
func Init(url string) {
	if url != "" {
		client = resty.New()
		serverlessURL = url
	}
	log := logrus.New()
	log.SetOutput(os.Stdout)
	Log = &Logrus{log}
}

// Level returns logger level
func (l Logrus) Level() log.Lvl {
	switch l.Logger.Level {
	case logrus.DebugLevel:
		return log.DEBUG
	case logrus.WarnLevel:
		return log.WARN
	case logrus.ErrorLevel:
		return log.ERROR
	case logrus.InfoLevel:
		return log.INFO
	default:
		l.Panic("Invalid level")
	}

	return log.OFF
}

// SetHeader is a stub to satisfy interface
// It's controlled by Logger
func (l Logrus) SetHeader(_ string) {}

// SetPrefix It's controlled by Logger
func (l Logrus) SetPrefix(s string) {}

// Prefix It's controlled by Logger
func (l Logrus) Prefix() string {
	return ""
}

// SetLevel set level to logger from given log.Lvl
func (l Logrus) SetLevel(lvl log.Lvl) {
	switch lvl {
	case log.DEBUG:
		l.Logger.SetLevel(logrus.DebugLevel)
	case log.WARN:
		l.Logger.SetLevel(logrus.WarnLevel)
	case log.ERROR:
		l.Logger.SetLevel(logrus.ErrorLevel)
	case log.INFO:
		l.Logger.SetLevel(logrus.InfoLevel)
	default:
		l.Panic("Invalid level")
	}
}

// Output logger output func
func (l Logrus) Output() io.Writer {
	return l.Out
}

// Printj print json log
func (l Logrus) Printj(j log.JSON) {
	l.WithFields(logrus.Fields(j)).Print()
}

// Debugj debug json log
func (l Logrus) Debugj(j log.JSON) {
	l.WithFields(logrus.Fields(j)).Debug()
}

// Infoj info json log
func (l Logrus) Infoj(j log.JSON) {
	l.WithFields(logrus.Fields(j)).Info()
}

// Warnj warning json log
func (l Logrus) Warnj(j log.JSON) {
	l.WithFields(logrus.Fields(j)).Warn()
}

// Errorj error json log
func (l Logrus) Errorj(j log.JSON) {
	l.WithFields(logrus.Fields(j)).Error()
}

// Fatalj fatal json log
func (l Logrus) Fatalj(j log.JSON) {
	l.WithFields(logrus.Fields(j)).Fatal()
}

// Panicj panic json log
func (l Logrus) Panicj(j log.JSON) {
	l.WithFields(logrus.Fields(j)).Panic()
}

func PushPanicLog(cl *logrus.Entry) {
	send("xxx_error", cl)
}

func PushErrorLog(err error) {
	cl := Log.WithFields(Fields{
		"type_str":     "ERR-ERROR",
		"error_type":   reflect.ValueOf(err).Type().String(),
		"error_string": err.Error(),
	})
	PushPanicLog(cl)
}

func PushLog(collectionName string, cl *logrus.Entry) {
	send(collectionName, cl)
}

func ActivityLog(cl *logrus.Entry) {
	send("xxx_activity", cl)
}

func PushDebugLog(cl *logrus.Entry) {
	send("xxx_debug", cl)
}

func send(collectionName string, cl *logrus.Entry) {
	if serverlessURL == "" {
		return
	}

	payload := (map[string]interface{})(cl.Data)
	payload["created_at"] = timeutil.NowStr()

	client.R().
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
		}).
		SetBody(
			map[string]interface{}{
				"collection_name": collectionName,
				"payload":         payload,
			},
		).
		Post(serverlessURL + "/v1/log")
}
