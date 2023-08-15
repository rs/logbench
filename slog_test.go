package logbench

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"strings"
	"time"
)

func init() {
	tests["slog"] = slogTester{}
}

type slogTester struct {
	l slog.Logger
}

var (
	_ logTesterArray = (*slogTester)(nil)
)

func (slogTester) newLogger(out io.Writer, disabled bool) logTester {
	lvl := slog.LevelInfo
	testFixup := func(groups []string, a slog.Attr) slog.Attr {
		switch a.Key {
		case "time":
			return slog.Attr{} // discard timestamps
		case "replaceTime-time":
			return slog.Attr{Key: "time", Value: slog.Int64Value(a.Value.Time().Unix())}
		case "replaceTimes-time":
			anyTimes := a.Value.Any().([]time.Time)
			anyUnixTimes := make([]int64, len(anyTimes))
			for i, anyTime := range anyTimes {
				anyUnixTimes[i] = anyTime.Unix()
			}
			return slog.Attr{Key: "time", Value: slog.AnyValue(anyUnixTimes)}
		case "duration":
			anyValue := a.Value.Any()
			switch anyValue := anyValue.(type) {
			case time.Duration:
				return slog.Attr{Key: "duration", Value: slog.IntValue(int(anyValue.Milliseconds()))}
			case []time.Duration:
				anyDurations := anyValue
				anyIntDurations := make([]int, len(anyDurations))
				for i, anyTime := range anyDurations {
					anyIntDurations[i] = int(anyTime.Milliseconds())
				}
				return slog.Attr{Key: "duration", Value: slog.AnyValue(anyIntDurations)}
			default:
				return a
			}
		case "level":
			lowerValue := slog.StringValue(strings.ToLower(a.Value.String()))
			return slog.Attr{Key: a.Key, Value: lowerValue}
		case "msg":
			return slog.Attr{Key: "message", Value: a.Value}
		case "errors":
			valueErrors := a.Value.Any().([]error)
			errorStrings := make([]string, len(valueErrors))
			for i, err := range valueErrors {
				errorStrings[i] = err.Error()
			}
			return slog.Attr{Key: "errors", Value: slog.AnyValue(errorStrings)}
		default:
			return a
		}
	}
	var handler slog.Handler = slog.NewJSONHandler(out, &slog.HandlerOptions{Level: lvl, ReplaceAttr: testFixup})
	if disabled {
		handler = disabledHandler{}
	}
	return slogTester{*slog.New(handler)}
}

func (t slogTester) logMsg(msg string) {
	t.l.Info(msg)
}

func (t slogTester) logFormat(format string, v ...interface{}) bool {
	t.l.Info(fmt.Sprintf(format, v...))
	return true
}

func (t slogTester) withContext(context map[string]interface{}) (logTester, bool) {
	var args = make([]any, len(context)*2)
	index := 0
	for s, i := range context {
		args[index] = s
		args[index+1] = i
		index += 2
	}
	return slogTester{*t.l.With(args...)}, true
}

func (t slogTester) logBool(msg, key string, value bool) bool {
	t.l.Info(msg, key, value)
	return true
}

func (t slogTester) logInt(msg, key string, value int) bool {
	t.l.Info(msg, key, value)
	return true
}

func (t slogTester) logFloat32(msg, key string, value float32) bool {
	t.l.Info(msg, key, value)
	return true
}

func (t slogTester) logFloat64(msg, key string, value float64) bool {
	t.l.Info(msg, key, value)
	return true
}

func (t slogTester) logTime(msg, key string, value time.Time) bool {
	t.l.Info(msg, "replaceTime-"+key, value)
	return true
}

func (t slogTester) logDuration(msg, key string, value time.Duration) bool {
	t.l.Info(msg, key, value)
	return true
}

func (t slogTester) logError(msg, key string, value error) bool {
	t.l.Info(msg, key, value)
	return true
}

func (t slogTester) logString(msg, key string, value string) bool {
	t.l.Info(msg, key, value)
	return true
}

func (t slogTester) logObject(msg, key string, value *obj) bool {
	t.l.Info(msg, key, value)
	return true
}

func (t slogTester) logBools(msg, key string, value []bool) bool {
	t.l.Info(msg, key, value)
	return true
}

func (t slogTester) logInts(msg, key string, value []int) bool {
	t.l.Info(msg, key, value)
	return true
}

func (t slogTester) logFloats32(msg, key string, value []float32) bool {
	t.l.Info(msg, key, value)
	return true
}

func (t slogTester) logFloats64(msg, key string, value []float64) bool {
	t.l.Info(msg, key, value)
	return true
}

func (t slogTester) logTimes(msg, key string, value []time.Time) bool {
	t.l.Info(msg, "replaceTimes-"+key, value)
	return true
}

func (t slogTester) logDurations(msg, key string, value []time.Duration) bool {
	t.l.Info(msg, key, value)
	return true
}

func (t slogTester) logErrors(msg, key string, value []error) bool {
	t.l.Info(msg, key, value)
	return true
}

func (t slogTester) logStrings(msg, key string, value []string) bool {
	t.l.Info(msg, key, value)
	return true
}

type disabledHandler struct {
}

func (d disabledHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return false
}

func (d disabledHandler) Handle(ctx context.Context, record slog.Record) error {
	return nil
}

func (d disabledHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return d
}

func (d disabledHandler) WithGroup(name string) slog.Handler {
	return d
}
