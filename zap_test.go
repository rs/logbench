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
	tests["Zap"] = zapTester{}
}

func (o *obj) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddString("name", o.Name)
	enc.AddInt("count", o.Count)
	enc.AddBool("enabled", o.Enabled)
	return nil
}

type zapTester struct {
	l *zap.Logger
}

var (
	_ logTesterArray = (*zapTester)(nil)
)

func zapMillisecondDurationEncoder(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendFloat64(float64(d) / float64(time.Millisecond))
}

func zapEncoder() zapcore.Encoder {
	ec := zap.NewProductionEncoderConfig()
	ec.EncodeDuration = zapMillisecondDurationEncoder
	ec.EncodeTime = zapcore.EpochTimeEncoder
	ec.TimeKey = ""
	ec.MessageKey = "message"
	return zapcore.NewJSONEncoder(ec)
}

func (zapTester) newLogger(out io.Writer, disabled bool) logTester {
	lvl := zap.DebugLevel
	if disabled {
		lvl = zap.FatalLevel
	}
	var w zapcore.WriteSyncer = &zaptest.Discarder{}
	if out != ioutil.Discard {
		w = zapcore.AddSync(out)
	}
	return zapTester{zap.New(zapcore.NewCore(zapEncoder(), w, lvl))}
}

func (t zapTester) logMsg(msg string) {
	t.l.Info(msg)
}

func (t zapTester) logFormat(format string, v ...interface{}) bool {
	return false
}

func (t zapTester) withContext(context map[string]interface{}) (logTester, bool) {
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
	return zapTester{l}, true
}

func (t zapTester) logBool(msg, key string, value bool) bool {
	t.l.Info(msg, zap.Bool(key, value))
	return true
}

func (t zapTester) logInt(msg, key string, value int) bool {
	t.l.Info(msg, zap.Int(key, value))
	return true
}

func (t zapTester) logFloat32(msg, key string, value float32) bool {
	t.l.Info(msg, zap.Float32(key, value))
	return true
}

func (t zapTester) logFloat64(msg, key string, value float64) bool {
	t.l.Info(msg, zap.Float64(key, value))
	return true
}

func (t zapTester) logTime(msg, key string, value time.Time) bool {
	t.l.Info(msg, zap.Time(key, value))
	return true
}

func (t zapTester) logDuration(msg, key string, value time.Duration) bool {
	t.l.Info(msg, zap.Duration(key, value))
	return true
}

func (t zapTester) logError(msg, key string, value error) bool {
	t.l.Info(msg, zap.NamedError(key, value))
	return true
}

func (t zapTester) logString(msg, key string, value string) bool {
	t.l.Info(msg, zap.String(key, value))
	return true
}

func (t zapTester) logObject(msg, key string, value *obj) bool {
	t.l.Info(msg, zap.Object(key, value))
	return true
}

func (t zapTester) logBools(msg, key string, value []bool) bool {
	t.l.Info(msg, zap.Bools(key, value))
	return true
}

func (t zapTester) logInts(msg, key string, value []int) bool {
	t.l.Info(msg, zap.Ints(key, value))
	return true
}

func (t zapTester) logFloats32(msg, key string, value []float32) bool {
	t.l.Info(msg, zap.Float32s(key, value))
	return true
}

func (t zapTester) logFloats64(msg, key string, value []float64) bool {
	t.l.Info(msg, zap.Float64s(key, value))
	return true
}

func (t zapTester) logTimes(msg, key string, value []time.Time) bool {
	t.l.Info(msg, zap.Times(key, value))
	return true
}

func (t zapTester) logDurations(msg, key string, value []time.Duration) bool {
	t.l.Info(msg, zap.Durations(key, value))
	return true
}

func (t zapTester) logErrors(msg, key string, value []error) bool {
	return false
}

func (t zapTester) logStrings(msg, key string, value []string) bool {
	t.l.Info(msg, zap.Strings(key, value))
	return true
}
