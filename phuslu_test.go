package logbench

import (
	"io"
	"time"

	"github.com/phuslu/log"
)

func init() {
	tests["PhusluLog"] = phusluLogTester{}
}

// MarshalObject implements log.ObjectMarshaler of phuslu/log.
func (o obj) MarshalObject(e *log.Entry) {
	e.Str("name", o.Name).Int("count", o.Count).Bool("enabled", o.Enabled)
}

type phusluLogTester struct {
	l log.Logger
}

var (
	_ logTesterArray = (*phusluLogTester)(nil)
)

func (t phusluLogTester) newLogger(out io.Writer, disabled bool) logTester {
	lvl := log.DebugLevel
	if disabled {
		lvl = log.PanicLevel
	}

	return phusluLogTester{l: log.Logger{
		Level:      lvl,
		Caller:     0,
		TimeField:  "ts",
		TimeFormat: log.TimeFormatUnix,
		Writer:     &log.IOWriter{Writer: out},
	}}
}

func (t phusluLogTester) logMsg(msg string) {
	t.l.Info().Msg(msg)
}

func (t phusluLogTester) logFormat(format string, v ...interface{}) bool {
	t.l.Info().Msgf(format, v...)
	return true
}

func (t phusluLogTester) withContext(context map[string]interface{}) (logTester, bool) {
	return nil, false
}

func (t phusluLogTester) logBool(msg, key string, value bool) bool {
	t.l.Info().Bool(key, value).Msg(msg)
	return true
}

func (t phusluLogTester) logInt(msg, key string, value int) bool {
	t.l.Info().Int(key, value).Msg(msg)
	return true
}

func (t phusluLogTester) logFloat32(msg, key string, value float32) bool {
	t.l.Info().Float32(key, value).Msg(msg)
	return true
}

func (t phusluLogTester) logFloat64(msg, key string, value float64) bool {
	t.l.Info().Float64(key, value).Msg(msg)
	return true
}

func (t phusluLogTester) logTime(msg, key string, value time.Time) bool {
	t.l.Info().TimeFormat(key, log.TimeFormatUnix, value).Msg(msg)
	return true
}

func (t phusluLogTester) logDuration(msg, key string, value time.Duration) bool {
	t.l.Info().Dur(key, value).Msg(msg)
	return true
}

func (t phusluLogTester) logError(msg, key string, value error) bool {
	t.l.Info().AnErr(key, value).Msg(msg)
	return true
}

func (t phusluLogTester) logString(msg, key string, value string) bool {
	t.l.Info().Str(key, value).Msg(msg)
	return true
}

func (t phusluLogTester) logObject(msg, key string, value *obj) bool {
	t.l.Info().Object(key, value).Msg(msg)
	return true
}

func (t phusluLogTester) logBools(msg, key string, value []bool) bool {
	t.l.Info().Bools(key, value).Msg(msg)
	return true
}

func (t phusluLogTester) logInts(msg, key string, value []int) bool {
	t.l.Info().Ints(key, value).Msg(msg)
	return true
}

func (t phusluLogTester) logFloats32(msg, key string, value []float32) bool {
	t.l.Info().Floats32(key, value).Msg(msg)
	return true
}

func (t phusluLogTester) logFloats64(msg, key string, value []float64) bool {
	t.l.Info().Floats64(key, value).Msg(msg)
	return true
}

func (t phusluLogTester) logTimes(msg, key string, value []time.Time) bool {
	t.l.Info().TimesFormat(key, log.TimeFormatUnix, value).Msg(msg)
	return true
}

func (t phusluLogTester) logDurations(msg, key string, value []time.Duration) bool {
	t.l.Info().Durs(key, value).Msg(msg)
	return true
}

func (t phusluLogTester) logErrors(msg, key string, value []error) bool {
	t.l.Info().Errs(key, value).Msg(msg)
	return true
}

func (t phusluLogTester) logStrings(msg, key string, value []string) bool {
	t.l.Info().Strs(key, value).Msg(msg)
	return true
}
