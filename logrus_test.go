package logbench

import (
	"io"
	"time"

	"github.com/sirupsen/logrus"
)

func init() {
	tests["Logrus"] = logrusTester{}
}

type logrusTester struct {
	l logrus.FieldLogger
}

var (
	_ logTesterArray = (*logrusTester)(nil)
)

func (logrusTester) newLogger(out io.Writer, disabled bool) logTester {
	lvl := logrus.DebugLevel
	if disabled {
		lvl = logrus.PanicLevel
	}
	return logrusTester{&logrus.Logger{
		Out: out,
		Formatter: &logrus.JSONFormatter{
			DisableTimestamp: true,
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime: "",
				logrus.FieldKeyMsg:  "message",
			},
		},
		Hooks: make(logrus.LevelHooks),
		Level: lvl,
	}}
}

func (t logrusTester) logMsg(msg string) {
	t.l.Info(msg)
}

func (t logrusTester) logFormat(format string, v ...interface{}) bool {
	t.l.Infof(format, v...)
	return true
}

func (t logrusTester) withContext(context map[string]interface{}) (logTester, bool) {
	return logrusTester{t.l.WithFields(context)}, true
}

func (t logrusTester) logBool(msg, key string, value bool) bool {
	t.l.WithFields(logrus.Fields{key: value}).Info(msg)
	return true
}

func (t logrusTester) logInt(msg, key string, value int) bool {
	t.l.WithFields(logrus.Fields{key: value}).Info(msg)
	return true
}

func (t logrusTester) logFloat32(msg, key string, value float32) bool {
	t.l.WithFields(logrus.Fields{key: value}).Info(msg)
	return true
}

func (t logrusTester) logFloat64(msg, key string, value float64) bool {
	t.l.WithFields(logrus.Fields{key: value}).Info(msg)
	return true
}

func (t logrusTester) logTime(msg, key string, value time.Time) bool {
	return false
}

func (t logrusTester) logDuration(msg, key string, value time.Duration) bool {
	return false
}

func (t logrusTester) logError(msg, key string, value error) bool {
	return false
}

func (t logrusTester) logString(msg, key string, value string) bool {
	t.l.WithFields(logrus.Fields{key: value}).Info(msg)
	return true
}

func (t logrusTester) logObject(msg, key string, value *obj) bool {
	t.l.WithFields(logrus.Fields{key: value}).Info(msg)
	return true
}

func (t logrusTester) logBools(msg, key string, value []bool) bool {
	t.l.WithFields(logrus.Fields{key: value}).Info(msg)
	return true
}

func (t logrusTester) logInts(msg, key string, value []int) bool {
	t.l.WithFields(logrus.Fields{key: value}).Info(msg)
	return true
}

func (t logrusTester) logFloats32(msg, key string, value []float32) bool {
	t.l.WithFields(logrus.Fields{key: value}).Info(msg)
	return true
}

func (t logrusTester) logFloats64(msg, key string, value []float64) bool {
	t.l.WithFields(logrus.Fields{key: value}).Info(msg)
	return true
}

func (t logrusTester) logTimes(msg, key string, value []time.Time) bool {
	return false
}

func (t logrusTester) logDurations(msg, key string, value []time.Duration) bool {
	return false
}

func (t logrusTester) logErrors(msg, key string, value []error) bool {
	return false
}

func (t logrusTester) logStrings(msg, key string, value []string) bool {
	t.l.WithFields(logrus.Fields{key: value}).Info(msg)
	return true
}
