package logbench

import (
	"io"
	"time"

	"github.com/axpira/gop/log"
	"github.com/axpira/goplogjson"
)

func init() {
	// zerolog.TimeFieldFormat = ""
	// zerolog.DurationFieldInteger = true
	// zerolog.MessageFieldName = "message"
	// goplog.TimestampEnable = false
	// goplog.MessageFieldName = "message"
	// goplog.TimeFormat = "UNIX"
	goplogjson.New()

	tests["GoplogJson"] = goplogjsonTester{}
}

func (o obj) MarshalLog(e log.FieldBuilder) {
	e.Str("name", o.Name).
		Int("count", o.Count).
		Bool("enabled", o.Enabled)
}

type goplogjsonTester struct {
	l log.Logger
}

// var (
// 	_ logTesterArray = (*goplogTester)(nil)
// )

func (goplogjsonTester) newLogger(out io.Writer, disabled bool) logTester {
	lvl := log.DebugLevel
	if disabled {
		lvl = log.DisabledLevel
	}
	goplogjson.TimestampEnabled = false
	goplogjson.MessageFieldName = "message"
	return goplogjsonTester{goplogjson.New(goplogjson.WithOutput(out), goplogjson.WithLevel(lvl))}
}

func (t goplogjsonTester) logMsg(msg string) {
	t.l.Info(msg)
}

func (t goplogjsonTester) logFormat(format string, v ...interface{}) bool {
	t.l.Infof(format, v...)
	return true
}

func (t goplogjsonTester) withContext(context map[string]interface{}) (logTester, bool) {
	return goplogjsonTester{t.l.With(log.Fields(context))}, true
}

func (t goplogjsonTester) logBool(msg, key string, value bool) bool {
	t.l.Inf(log.Bool(key, value).Msg(msg))
	return true
}

func (t goplogjsonTester) logInt(msg, key string, value int) bool {
	t.l.Inf(log.Int(key, value).Msg(msg))
	return true
}

func (t goplogjsonTester) logFloat32(msg, key string, value float32) bool {
	t.l.Inf(log.Float32(key, value).Msg(msg))
	return true
}

func (t goplogjsonTester) logFloat64(msg, key string, value float64) bool {
	t.l.Inf(log.Float64(key, value).Msg(msg))
	return true
}

func (t goplogjsonTester) logTime(msg, key string, value time.Time) bool {
	// t.l.Inf(log.Time(key, value).Msg(msg))
	return false
}

func (t goplogjsonTester) logDuration(msg, key string, value time.Duration) bool {
	t.l.Inf(log.Dur(key, value).Msg(msg))
	return true
}

func (t goplogjsonTester) logError(msg, key string, value error) bool {
	t.l.Inf(log.Error(key, value).Msg(msg))
	return true
}

func (t goplogjsonTester) logString(msg, key string, value string) bool {
	t.l.Inf(log.Str(key, value).Msg(msg))
	return true
}

func (t goplogjsonTester) logObject(msg, key string, value *obj) bool {
	t.l.Inf(t.l.NewFieldBuilder().Marshal(key, value).Msg(msg))
	return true
}

// func (t goplogTester) logBools(msg, key string, value []bool) bool {
// 	t.l.Info().Bools(key, value).Msg(msg)
// 	return true
// }

// func (t goplogTester) logInts(msg, key string, value []int) bool {
// 	t.l.Info().Ints(key, value).Msg(msg)
// 	return true
// }

// func (t goplogTester) logFloats32(msg, key string, value []float32) bool {
// 	t.l.Info().Floats32(key, value).Msg(msg)
// 	return true
// }

// func (t goplogTester) logFloats64(msg, key string, value []float64) bool {
// 	t.l.Info().Floats64(key, value).Msg(msg)
// 	return true
// }

// func (t goplogTester) logTimes(msg, key string, value []time.Time) bool {
// 	t.l.Info().Times(key, value).Msg(msg)
// 	return true
// }

// func (t goplogTester) logDurations(msg, key string, value []time.Duration) bool {
// 	t.l.Info().Durs(key, value).Msg(msg)
// 	return true
// }

// func (t goplogTester) logErrors(msg, key string, value []error) bool {
// 	t.l.Info().Errs(key, value).Msg(msg)
// 	return true
// }

// func (t goplogTester) logStrings(msg, key string, value []string) bool {
// 	t.l.Info().Strs(key, value).Msg(msg)
// 	return true
// }
