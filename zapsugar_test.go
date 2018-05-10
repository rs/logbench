package logbench

import (
	"fmt"
	"io"
	"io/ioutil"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"
)

func init() {
	tests["ZapSugar"] = zapSugarTester{}
}

type zapSugarTester struct {
	l *zap.SugaredLogger
}

var (
	_ logTesterArray = (*zapSugarTester)(nil)
)

func (zapSugarTester) newLogger(out io.Writer, disabled bool) logTester {
	// TOFIX: use out
	lvl := zap.DebugLevel
	if disabled {
		lvl = zap.FatalLevel
	}
	var w zapcore.WriteSyncer = &zaptest.Discarder{}
	if out != ioutil.Discard {
		w = zapcore.AddSync(out)
	}
	return zapSugarTester{zap.New(zapcore.NewCore(zapEncoder(), w, lvl)).Sugar()}
}

func (t zapSugarTester) logMsg(msg string) {
	t.l.Info(msg)
}

func (t zapSugarTester) logFormat(format string, v ...interface{}) bool {
	t.l.Infof(format, v...)
	return true
}

func (t zapSugarTester) withContext(context map[string]interface{}) (logTester, bool) {
	l := t.l
	for k, v := range context {
		switch v := v.(type) {
		case int:
			l = l.With(zap.Int(k, v))
		case string:
			l = l.With(zap.String(k, v))
		case error:
			l = l.With(zap.NamedError(k, v))
		case time.Time:
			l = l.With(zap.Time(k, v))
		case *obj:
			l = l.With(zap.Object(k, v))
		default:
			panic(fmt.Sprintf("zap: unsupported context field type: %v", v))
		}
	}
	return zapSugarTester{l}, true
}

func (t zapSugarTester) logBool(msg, key string, value bool) bool {
	t.l.Infow(msg, key, value)
	return true
}

func (t zapSugarTester) logInt(msg, key string, value int) bool {
	t.l.Infow(msg, key, value)
	return true
}

func (t zapSugarTester) logFloat32(msg, key string, value float32) bool {
	t.l.Infow(msg, key, value)
	return true
}

func (t zapSugarTester) logFloat64(msg, key string, value float64) bool {
	t.l.Infow(msg, key, value)
	return true
}

func (t zapSugarTester) logTime(msg, key string, value time.Time) bool {
	t.l.Infow(msg, key, value)
	return true
}

func (t zapSugarTester) logDuration(msg, key string, value time.Duration) bool {
	t.l.Infow(msg, key, value)
	return true
}

func (t zapSugarTester) logError(msg, key string, value error) bool {
	t.l.Infow(msg, key, value)
	return true
}

func (t zapSugarTester) logString(msg, key string, value string) bool {
	t.l.Infow(msg, key, value)
	return true
}

func (t zapSugarTester) logObject(msg, key string, value *obj) bool {
	t.l.Infow(msg, key, value)
	return true
}

func (t zapSugarTester) logBools(msg, key string, value []bool) bool {
	t.l.Infow(msg, key, value)
	return true
}

func (t zapSugarTester) logInts(msg, key string, value []int) bool {
	t.l.Infow(msg, key, value)
	return true
}

func (t zapSugarTester) logFloats32(msg, key string, value []float32) bool {
	t.l.Infow(msg, key, value)
	return true
}

func (t zapSugarTester) logFloats64(msg, key string, value []float64) bool {
	t.l.Infow(msg, key, value)
	return true
}

func (t zapSugarTester) logTimes(msg, key string, value []time.Time) bool {
	t.l.Infow(msg, key, value)
	return true
}

func (t zapSugarTester) logDurations(msg, key string, value []time.Duration) bool {
	t.l.Infow(msg, key, value)
	return true
}

func (t zapSugarTester) logErrors(msg, key string, value []error) bool {
	return false
}

func (t zapSugarTester) logStrings(msg, key string, value []string) bool {
	t.l.Infow(msg, key, value)
	return true
}
