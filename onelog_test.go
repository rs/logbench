package logbench

import (
	"fmt"
	"io"
	"time"

	"github.com/francoispqt/gojay"
	"github.com/francoispqt/onelog"
)

func init() {
	tests["Onelog"] = onelogTester{}
}

func (o obj) MarshalObject(enc *gojay.Encoder) {
	enc.AddStringKey("name", o.Name)
	enc.AddIntKey("count", o.Count)
	enc.AddIntKey("created", int(o.Created.Unix()))
	enc.AddBoolKey("enabled", o.Enabled)
}

func (o obj) IsNil() bool {
	return false
}

type onelogTester struct {
	l *onelog.Logger
}

func (onelogTester) newLogger(out io.Writer, disabled bool) logTester {
	lvl := onelog.ALL
	if disabled {
		lvl = 0
	}
	return onelogTester{onelog.New(out, lvl)}
}

func (t onelogTester) logMsg(msg string) {
	t.l.Info(msg)
}

func (t onelogTester) logFormat(format string, v ...interface{}) bool {
	return false
}

func (t onelogTester) withContext(context map[string]interface{}) (logTester, bool) {
	l := t.l.With(func(e onelog.Entry) {
		for k, v := range context {
			switch v := v.(type) {
			case int:
				e.Int(k, v)
			case string:
				e.String(k, v)
			case error:
				e.String(k, v.Error())
			case time.Time:
				e.Int64(k, v.Unix())
			case *obj:
				e.Object(k, v)
			default:
				panic(fmt.Sprintf("onelog: unsupported context field type: %v", v))
			}
		}
	})
	return onelogTester{l}, true
}

func (t onelogTester) logBool(msg, key string, value bool) bool {
	t.l.InfoWithFields(msg, func(e onelog.Entry) {
		e.Bool(key, value)
	})
	return true
}

func (t onelogTester) logInt(msg, key string, value int) bool {
	t.l.InfoWithFields(msg, func(e onelog.Entry) {
		e.Int(key, value)
	})
	return true
}

func (t onelogTester) logFloat32(msg, key string, value float32) bool {
	return false
}

func (t onelogTester) logFloat64(msg, key string, value float64) bool {
	t.l.InfoWithFields(msg, func(e onelog.Entry) {
		e.Float(key, value)
	})
	return true
}

func (t onelogTester) logTime(msg, key string, value time.Time) bool {
	return false
}

func (t onelogTester) logDuration(msg, key string, value time.Duration) bool {
	return false
}

func (t onelogTester) logError(msg, key string, value error) bool {
	t.l.InfoWithFields(msg, func(e onelog.Entry) {
		e.Error(key, value)
	})
	return true
}

func (t onelogTester) logString(msg, key string, value string) bool {
	t.l.InfoWithFields(msg, func(e onelog.Entry) {
		e.String(key, value)
	})
	return true
}

func (t onelogTester) logBytes(msg, key string, value []byte) bool {
	return false
}

func (t onelogTester) logHex(msg, key string, value []byte) bool {
	return false
}

func (t onelogTester) logObject(msg, key string, value *obj) bool {
	t.l.InfoWithFields(msg, func(e onelog.Entry) {
		e.Object(key, value)
	})
	return true
}
