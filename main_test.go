package logbench

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"testing"
	"time"

	"github.com/juju/testing/checkers"
)

type tester interface {
	newLogger(out io.Writer, disabled bool) logTester
}

type logTester interface {
	logMsg(msg string)
	logFormat(format string, v ...interface{}) bool

	withContext(context map[string]interface{}) (logTester, bool)

	logBool(msg, key string, value bool) bool
	logInt(msg, key string, value int) bool
	logFloat32(msg, key string, value float32) bool
	logFloat64(msg, key string, value float64) bool
	logTime(msg, key string, value time.Time) bool
	logDuration(msg, key string, value time.Duration) bool
	logError(msg, key string, value error) bool
	logString(msg, key string, value string) bool
	logBytes(msg, key string, value []byte) bool
	logHex(msg, key string, value []byte) bool
	logObject(msg, key string, value *obj) bool
}

type logTesterArray interface {
	logBools(msg, key string, value []bool) bool
	logInts(msg, key string, value []int) bool
	logFloats32(msg, key string, value []float32) bool
	logFloats64(msg, key string, value []float64) bool
	logTimes(msg, key string, value []time.Time) bool
	logDurations(msg, key string, value []time.Duration) bool
	logErrors(msg, key string, value []error) bool
	logStrings(msg, key string, value []string) bool
}

var tests = map[string]tester{}

type obj struct {
	Name    string `json:"name"`
	Count   int    `json:"count"`
	Enabled bool   `json:"enabled"`
}

func (o obj) jsonMap() map[string]interface{} {
	return map[string]interface{}{
		"name":    o.Name,
		"count":   o.Count,
		"enabled": o.Enabled,
	}
}

type objs []*obj

var (
	sampleBools    = []bool{true, false, true, false, true, false, true, false, true, false}
	sampleInts     = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	sampleFloats32 = []float32{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	sampleFloats64 = []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.0, 0.1}
	sampleStrings  = fakeStrings(10)
	sampleErrors   = fakeErrors(10)
	sampleTimes    = []time.Time{
		time.Unix(0, 0),
		time.Unix(1, 0),
		time.Unix(2, 0),
		time.Unix(3, 0),
		time.Unix(4, 0),
		time.Unix(5, 0),
		time.Unix(6, 0),
		time.Unix(7, 0),
		time.Unix(8, 0),
		time.Unix(9, 0),
	}
	sampleDurations = []time.Duration{
		0 * time.Millisecond,
		1 * time.Millisecond,
		2 * time.Millisecond,
		3 * time.Millisecond,
		4 * time.Millisecond,
		5 * time.Millisecond,
		6 * time.Millisecond,
		7 * time.Millisecond,
		8 * time.Millisecond,
		9 * time.Millisecond,
	}
	sampleObjects = objs{
		&obj{"a", 1, true},
		&obj{"b", 2, false},
		&obj{"c", 3, true},
		&obj{"d", 4, false},
		&obj{"e", 5, true},
		&obj{"f", 6, false},
		&obj{"g", 7, true},
		&obj{"h", 8, false},
		&obj{"i", 9, true},
		&obj{"j", 0, false},
	}

	sampleMsg        = "Test logging, but use a somewhat realistic message length."
	sampleFormat     = "Test format %d %f %s"
	sampleFormatArgs = []interface{}{1, 1.0, "some string"}

	sampleContext = map[string]interface{}{
		"ctx_int":    sampleInts[0],
		"ctx_string": sampleStrings[0],
		"ctx_error":  sampleErrors[0],
		"ctx_object": sampleObjects[0],
	}
)

func fakeStrings(n int) []string {
	strs := make([]string, n)
	for i := range strs {
		strs[i] = fmt.Sprintf("sample string %d", i)
	}
	return strs
}

func fakeErrors(n int) []error {
	errs := make([]error, n)
	for i := range errs {
		errs[i] = fmt.Errorf("sample error number %d", i)
	}
	return errs
}

func Benchmark(b *testing.B) {
	for name, t := range tests {
		b.Run(name, func(b *testing.B) {
			b.Run("Enabled", func(b *testing.B) {
				l := t.newLogger(ioutil.Discard, false)
				benchmarkContext(b, l)
			})
			b.Run("Disabled", func(b *testing.B) {
				l := t.newLogger(ioutil.Discard, true)
				benchmarkContext(b, l)
			})
		})
	}
}

func benchmarkContext(b *testing.B, l logTester) {
	b.Run("NoContext", func(b *testing.B) {
		benchmark(b, l)
	})
	if l, supported := l.withContext(sampleContext); supported {
		b.Run("WithContext", func(b *testing.B) {
			benchmark(b, l)
		})
	}
}

func benchmark(b *testing.B, l logTester) {
	b.Run("Msg", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				l.logMsg(sampleMsg)
			}
		})
	})
	if l.logFormat(sampleFormat, sampleFormatArgs...) {
		b.Run("Formatting", func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					l.logFormat(sampleFormat, sampleFormatArgs...)
				}
			})
		})
	}
	b.Run("Fields", func(b *testing.B) {
		benchmarkFields(b, l)
	})
}

func benchmarkFields(b *testing.B, l logTester) {
	bs := map[string]func(){}
	if l.logBool(sampleMsg, "bool", sampleBools[0]) {
		bs["Bool"] = func() {
			l.logBool(sampleMsg, "bool", sampleBools[0])
		}
	}
	if l.logInt(sampleMsg, "int", sampleInts[0]) {
		bs["Int"] = func() {
			l.logInt(sampleMsg, "int", sampleInts[0])
		}
	}
	if l.logFloat32(sampleMsg, "float32", sampleFloats32[0]) {
		bs["Float32"] = func() {
			l.logFloat32(sampleMsg, "float32", sampleFloats32[0])
		}
	}
	if l.logFloat64(sampleMsg, "float64", sampleFloats64[0]) {
		bs["Float64"] = func() {
			l.logFloat64(sampleMsg, "float64", sampleFloats64[0])
		}
	}
	if l.logTime(sampleMsg, "time", sampleTimes[0]) {
		bs["Time"] = func() {
			l.logTime(sampleMsg, "time", sampleTimes[0])
		}
	}
	if l.logDuration(sampleMsg, "duration", sampleDurations[0]) {
		bs["Time"] = func() {
			l.logDuration(sampleMsg, "duration", sampleDurations[0])
		}
	}
	if l.logString(sampleMsg, "string", sampleStrings[0]) {
		bs["String"] = func() {
			l.logString(sampleMsg, "string", sampleStrings[0])
		}
	}
	if l.logError(sampleMsg, "error", sampleErrors[0]) {
		bs["Error"] = func() {
			l.logError(sampleMsg, "error", sampleErrors[0])
		}
	}
	if l.logObject(sampleMsg, "object", sampleObjects[0]) {
		bs["Object"] = func() {
			l.logObject(sampleMsg, "object", sampleObjects[0])
		}
	}

	if l, ok := l.(logTesterArray); ok {
		if l.logBools(sampleMsg, "bools", sampleBools) {
			bs["Bools"] = func() {
				l.logBools(sampleMsg, "bools", sampleBools)
			}
		}
		if l.logInts(sampleMsg, "ints", sampleInts) {
			bs["Ints"] = func() {
				l.logInts(sampleMsg, "ints", sampleInts)
			}
		}
		if l.logFloats32(sampleMsg, "floats32", sampleFloats32) {
			bs["Floats32"] = func() {
				l.logFloats32(sampleMsg, "floats32", sampleFloats32)
			}
		}
		if l.logFloats64(sampleMsg, "floats64", sampleFloats64) {
			bs["Floats64"] = func() {
				l.logFloats64(sampleMsg, "floats64", sampleFloats64)
			}
		}
		if l.logTimes(sampleMsg, "time", sampleTimes) {
			bs["Time"] = func() {
				l.logTimes(sampleMsg, "time", sampleTimes)
			}
		}
		if l.logDurations(sampleMsg, "duration", sampleDurations) {
			bs["Time"] = func() {
				l.logDurations(sampleMsg, "duration", sampleDurations)
			}
		}
		if l.logStrings(sampleMsg, "strings", sampleStrings) {
			bs["Strings"] = func() {
				l.logStrings(sampleMsg, "strings", sampleStrings)
			}
		}
		if l.logErrors(sampleMsg, "errors", sampleErrors) {
			bs["Errors"] = func() {
				l.logErrors(sampleMsg, "errors", sampleErrors)
			}
		}
	}

	b.ResetTimer()
	for name, f := range bs {
		b.Run(name, func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					f()
				}
			})
		})
	}
}

type testCtx struct {
	tester  logTester
	buf     *bytes.Buffer
	context map[string]interface{}
	enabled bool
}

func Test(t *testing.T) {
	for name, lt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Run("Enabled", func(t *testing.T) {
				buf := &bytes.Buffer{}
				ctx := &testCtx{
					tester:  lt.newLogger(buf, false),
					buf:     buf,
					enabled: true,
				}
				testContext(t, ctx)
			})
			t.Run("Disabled", func(t *testing.T) {
				buf := &bytes.Buffer{}
				ctx := &testCtx{
					tester:  lt.newLogger(buf, true),
					buf:     buf,
					enabled: false,
				}
				testContext(t, ctx)
			})
		})
	}
}

func testContext(t *testing.T, ctx *testCtx) {
	t.Run("NoContext", func(t *testing.T) {
		test(t, ctx)
	})
	if l, supported := ctx.tester.withContext(sampleContext); supported {
		t.Run("WithContext", func(t *testing.T) {
			ctx.context = sampleContext
			ctx.tester = l
			test(t, ctx)
		})
	}
}

func test(t *testing.T, ctx *testCtx) {
	ctx.tester.logMsg(sampleMsg)
	validate(t, ctx, "Msg", true, map[string]interface{}{"message": sampleMsg})
	validate(t, ctx, "Formatting",
		ctx.tester.logFormat(sampleFormat, sampleFormatArgs...),
		map[string]interface{}{"message": fmt.Sprintf(sampleFormat, sampleFormatArgs...)})
	t.Run("Fields", func(t *testing.T) {
		testFields(t, ctx)
	})
}

func testFields(t *testing.T, ctx *testCtx) {
	l := ctx.tester
	validate(t, ctx, "Bool",
		l.logBool(sampleMsg, "bool", sampleBools[0]),
		map[string]interface{}{"message": sampleMsg, "bool": sampleBools[0]})
	validate(t, ctx, "Int",
		l.logInt(sampleMsg, "int", sampleInts[0]),
		map[string]interface{}{"message": sampleMsg, "int": sampleInts[0]})
	validate(t, ctx, "Float32",
		l.logFloat32(sampleMsg, "float32", sampleFloats32[0]),
		map[string]interface{}{"message": sampleMsg, "float32": sampleFloats32[0]})
	validate(t, ctx, "Float64",
		l.logFloat64(sampleMsg, "float64", sampleFloats64[0]),
		map[string]interface{}{"message": sampleMsg, "float64": sampleFloats64[0]})
	validate(t, ctx, "Time",
		l.logTime(sampleMsg, "time", sampleTimes[0]),
		map[string]interface{}{"message": sampleMsg, "time": sampleTimes[0]})
	validate(t, ctx, "Duration",
		l.logDuration(sampleMsg, "duration", sampleDurations[0]),
		map[string]interface{}{"message": sampleMsg, "duration": sampleDurations[0]})
	validate(t, ctx, "String",
		l.logString(sampleMsg, "string", sampleStrings[0]),
		map[string]interface{}{"message": sampleMsg, "string": sampleStrings[0]})
	validate(t, ctx, "Error",
		l.logError(sampleMsg, "error", sampleErrors[0]),
		map[string]interface{}{"message": sampleMsg, "error": sampleErrors[0]})
	validate(t, ctx, "Object",
		l.logObject(sampleMsg, "object", sampleObjects[0]),
		map[string]interface{}{"message": sampleMsg, "object": sampleObjects[0]})

	if l, ok := l.(logTesterArray); ok {
		validate(t, ctx, "Bools",
			l.logBools(sampleMsg, "bools", sampleBools),
			map[string]interface{}{"message": sampleMsg, "bools": sampleBools})
		validate(t, ctx, "Ints",
			l.logInts(sampleMsg, "ints", sampleInts),
			map[string]interface{}{"message": sampleMsg, "ints": sampleInts})
		validate(t, ctx, "Floats32",
			l.logFloats32(sampleMsg, "floats32", sampleFloats32),
			map[string]interface{}{"message": sampleMsg, "floats32": sampleFloats32})
		validate(t, ctx, "Floats64",
			l.logFloats64(sampleMsg, "floats64", sampleFloats64),
			map[string]interface{}{"message": sampleMsg, "floats64": sampleFloats64})
		validate(t, ctx, "Times",
			l.logTimes(sampleMsg, "time", sampleTimes),
			map[string]interface{}{"message": sampleMsg, "time": sampleTimes})
		validate(t, ctx, "Durations",
			l.logDurations(sampleMsg, "duration", sampleDurations),
			map[string]interface{}{"message": sampleMsg, "duration": sampleDurations})
		validate(t, ctx, "Strings",
			l.logStrings(sampleMsg, "strings", sampleStrings),
			map[string]interface{}{"message": sampleMsg, "strings": sampleStrings})
		validate(t, ctx, "Errors",
			l.logErrors(sampleMsg, "errors", sampleErrors),
			map[string]interface{}{"message": sampleMsg, "errors": sampleErrors})
	}
}

func validate(t *testing.T, ctx *testCtx, name string, supported bool, want map[string]interface{}) {
	t.Helper()
	defer ctx.buf.Reset()
	if !supported {
		return
	}
	t.Run(name, func(t *testing.T) {
		if !ctx.enabled {
			if ctx.buf.Len() != 0 {
				t.Errorf("wrote output while disabled: %v", ctx.buf.String())
			}
			return
		}
		want["level"] = "info"
		for k, v := range ctx.context {
			want[k] = v
		}
		fixTypes(want)
		wj, _ := json.Marshal(want)
		want = map[string]interface{}{}
		json.Unmarshal(wj, &want)
		var got map[string]interface{}
		if err := json.Unmarshal(ctx.buf.Bytes(), &got); err != nil {
			t.Fatalf("invalid JSON output: %v: %v", err, ctx.buf.String())
		}
		if eq, err := checkers.DeepEqual(got, want); !eq {
			t.Errorf("invalid output: %v\ngot: %s\nwant: %s", err, ctx.buf.String(), wj)
		}
	})
}

func fixTypes(m map[string]interface{}) {
	for k, v := range m {
		switch v := v.(type) {
		case *obj:
			m[k] = v.jsonMap()
		case map[string]interface{}:
			fixTypes(v)
		case time.Time:
			m[k] = v.Unix()
		case []time.Time:
			v2 := make([]int64, len(v))
			for i, t := range v {
				v2[i] = t.Unix()
			}
			m[k] = v2
		case time.Duration:
			m[k] = v / time.Millisecond
		case []time.Duration:
			v2 := make([]int64, len(v))
			for i, d := range v {
				v2[i] = int64(d / time.Millisecond)
			}
			m[k] = v2
		case error:
			m[k] = v.Error()
		case []error:
			v2 := make([]string, len(v))
			for i, e := range v {
				v2[i] = e.Error()
			}
			m[k] = v2
		}
	}
}
