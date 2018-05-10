package logbench

import (
	"io"
	"time"

	"github.com/rs/zerolog"
)

func init() {
	zerolog.TimeFieldFormat = ""
	zerolog.DurationFieldInteger = true
	zerolog.MessageFieldName = "message"

	tests["Zerolog"] = zerologTester{}
}

func (o obj) MarshalZerologObject(e *zerolog.Event) {
	e.Str("name", o.Name).
		Int("count", o.Count).
		Bool("enabled", o.Enabled)
}

type zerologTester struct {
	l zerolog.Logger
}

var (
	_ logTesterArray = (*zerologTester)(nil)
)

func (zerologTester) newLogger(out io.Writer, disabled bool) logTester {
	lvl := zerolog.DebugLevel
	if disabled {
		lvl = zerolog.Disabled
	}
	return zerologTester{zerolog.New(out).Level(lvl)}
}

func (t zerologTester) logMsg(msg string) {
	t.l.Info().Msg(msg)
}

func (t zerologTester) logFormat(format string, v ...interface{}) bool {
	t.l.Info().Msgf(format, v...)
	return true
}

func (t zerologTester) withContext(context map[string]interface{}) (logTester, bool) {
	return zerologTester{t.l.With().Fields(context).Logger()}, true
}

func (t zerologTester) logBool(msg, key string, value bool) bool {
	t.l.Info().Bool(key, value).Msg(msg)
	return true
}

func (t zerologTester) logInt(msg, key string, value int) bool {
	t.l.Info().Int(key, value).Msg(msg)
	return true
}

func (t zerologTester) logFloat32(msg, key string, value float32) bool {
	t.l.Info().Float32(key, value).Msg(msg)
	return true
}

func (t zerologTester) logFloat64(msg, key string, value float64) bool {
	t.l.Info().Float64(key, value).Msg(msg)
	return true
}

func (t zerologTester) logTime(msg, key string, value time.Time) bool {
	t.l.Info().Time(key, value).Msg(msg)
	return true
}

func (t zerologTester) logDuration(msg, key string, value time.Duration) bool {
	t.l.Info().Dur(key, value).Msg(msg)
	return true
}

func (t zerologTester) logError(msg, key string, value error) bool {
	t.l.Info().AnErr(key, value).Msg(msg)
	return true
}

func (t zerologTester) logString(msg, key string, value string) bool {
	t.l.Info().Str(key, value).Msg(msg)
	return true
}

func (t zerologTester) logObject(msg, key string, value *obj) bool {
	t.l.Info().Object(key, value).Msg(msg)
	return true
}

func (t zerologTester) logBools(msg, key string, value []bool) bool {
	t.l.Info().Bools(key, value).Msg(msg)
	return true
}

func (t zerologTester) logInts(msg, key string, value []int) bool {
	t.l.Info().Ints(key, value).Msg(msg)
	return true
}

func (t zerologTester) logFloats32(msg, key string, value []float32) bool {
	t.l.Info().Floats32(key, value).Msg(msg)
	return true
}

func (t zerologTester) logFloats64(msg, key string, value []float64) bool {
	t.l.Info().Floats64(key, value).Msg(msg)
	return true
}

func (t zerologTester) logTimes(msg, key string, value []time.Time) bool {
	t.l.Info().Times(key, value).Msg(msg)
	return true
}

func (t zerologTester) logDurations(msg, key string, value []time.Duration) bool {
	t.l.Info().Durs(key, value).Msg(msg)
	return true
}

func (t zerologTester) logErrors(msg, key string, value []error) bool {
	t.l.Info().Errs(key, value).Msg(msg)
	return true
}

func (t zerologTester) logStrings(msg, key string, value []string) bool {
	t.l.Info().Strs(key, value).Msg(msg)
	return true
}
