package output

import (
	"context"
	"fmt"
	"io"
)

type Verbosity int

const (
	VerbosityQuiet Verbosity = iota - 1
	VerbosityNorm
	VerbosityInfo
	VerbosityDebug
	VerbosityTrace
)

type Output func(ctx context.Context, verb Verbosity, msg string, args ...KeyValue) (int, error)

func (o Output) Print(ctx context.Context, args ...interface{}) {
	o(ctx, VerbosityNorm, fmt.Sprint(args...))
}

func (o Output) PrintKV(ctx context.Context, msg string, kv ...KeyValue) {
	o(ctx, VerbosityNorm, msg, kv...)
}

func (o Output) Printf(ctx context.Context, format string, args ...interface{}) {
	o(ctx, VerbosityNorm, fmt.Sprintf(format, args...))
}

func (o Output) Println(ctx context.Context, args ...interface{}) {
	o(ctx, VerbosityNorm, fmt.Sprintln(args...))
}

func (o Output) Info(ctx context.Context, args ...interface{}) {
	o(ctx, VerbosityInfo, fmt.Sprint(args...))
}

func (o Output) InfoKV(ctx context.Context, msg string, kv ...KeyValue) {
	o(ctx, VerbosityInfo, msg, kv...)
}

func (o Output) Debug(ctx context.Context, args ...interface{}) {
	o(ctx, VerbosityDebug, fmt.Sprint(args...))
}

func (o Output) DebugKV(ctx context.Context, msg string, kv ...KeyValue) {
	o(ctx, VerbosityDebug, msg, kv...)
}

func (o Output) Trace(ctx context.Context, args ...interface{}) {
	o(ctx, VerbosityTrace, fmt.Sprint(args...))
}

func (o Output) TraceKV(ctx context.Context, msg string, kv ...KeyValue) {
	o(ctx, VerbosityTrace, msg, kv...)
}

func (o Output) Write(b []byte) (int, error) {
	return o(context.Background(), VerbosityNorm, string(b))
}

func (o Output) Writer(ctx context.Context, verb Verbosity) io.Writer {
	return verbosityWriter{ctx, o, verb}
}

type verbosityWriter struct {
	ctx  context.Context
	out  Output
	verb Verbosity
}

func (w verbosityWriter) Write(b []byte) (int, error) {
	return w.out(w.ctx, w.verb, string(b))
}
