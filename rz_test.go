package logbench

import (
	"io"
	"time"

	"github.com/bloom42/rz-go"
)

func init() {
	rz.DurationFieldUnit = time.Millisecond
	rz.DurationFieldInteger = true

	tests["Rz"] = rzTester{}
}

func (o obj) MarshalRzObject(e *rz.Event) {
	e.Append(rz.String("name", o.Name),
		rz.Int("count", o.Count),
		rz.Bool("enabled", o.Enabled))
}

type rzTester struct {
	l rz.Logger
}

var (
	_ logTesterArray = (*rzTester)(nil)
)

func (rzTester) newLogger(out io.Writer, disabled bool) logTester {
	lvl := rz.DebugLevel
	if disabled {
		lvl = rz.Disabled
	}
	return rzTester{
		rz.New(rz.Writer(out), rz.Level(lvl), rz.TimeFieldFormat("")).
			With(rz.Fields(rz.Timestamp(false))),
	}
}

func (t rzTester) logMsg(msg string) {
	t.l.Info(msg)
}

func (t rzTester) logFormat(format string, v ...interface{}) bool {
	return false
}

func (t rzTester) withContext(context map[string]interface{}) (logTester, bool) {
	return rzTester{t.l.With(rz.Fields(rz.Map(context)))}, true
}

func (t rzTester) logBool(msg, key string, value bool) bool {
	t.l.Info(msg, rz.Bool(key, value))
	return true
}

func (t rzTester) logInt(msg, key string, value int) bool {
	t.l.Info(msg, rz.Int(key, value))
	return true
}

func (t rzTester) logFloat32(msg, key string, value float32) bool {
	t.l.Info(msg, rz.Float32(key, value))
	return true
}

func (t rzTester) logFloat64(msg, key string, value float64) bool {
	t.l.Info(msg, rz.Float64(key, value))
	return true
}

func (t rzTester) logTime(msg, key string, value time.Time) bool {
	t.l.Info(msg, rz.Time(key, value))
	return true
}

func (t rzTester) logDuration(msg, key string, value time.Duration) bool {
	t.l.Info(msg, rz.Duration(key, value))
	return true
}

func (t rzTester) logError(msg, key string, value error) bool {
	t.l.Info(msg, rz.Error(key, value))
	return true
}

func (t rzTester) logString(msg, key string, value string) bool {
	t.l.Info(msg, rz.String(key, value))
	return true
}

func (t rzTester) logObject(msg, key string, value *obj) bool {
	t.l.Info(msg, rz.Object(key, value))
	return true
}

func (t rzTester) logBools(msg, key string, value []bool) bool {
	t.l.Info(msg, rz.Bools(key, value))
	return true
}

func (t rzTester) logInts(msg, key string, value []int) bool {
	t.l.Info(msg, rz.Ints(key, value))
	return true
}

func (t rzTester) logFloats32(msg, key string, value []float32) bool {
	t.l.Info(msg, rz.Floats32(key, value))
	return true
}

func (t rzTester) logFloats64(msg, key string, value []float64) bool {
	t.l.Info(msg, rz.Floats64(key, value))
	return true
}

func (t rzTester) logTimes(msg, key string, value []time.Time) bool {
	t.l.Info(msg, rz.Times(key, value))
	return true
}

func (t rzTester) logDurations(msg, key string, value []time.Duration) bool {
	t.l.Info(msg, rz.Durations(key, value))
	return true
}

func (t rzTester) logErrors(msg, key string, value []error) bool {
	t.l.Info(msg, rz.Errors(key, value))
	return true
}

func (t rzTester) logStrings(msg, key string, value []string) bool {
	t.l.Info(msg, rz.Strings(key, value))
	return true
}
