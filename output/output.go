package output

import (
	"context"
	"fmt"
	"io"
	"os"

	"gitoa.ru/go-4devs/console/output/verbosity"
)

func writeError(_ int, err error) {
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
}

type Output func(ctx context.Context, verb verbosity.Verbosity, msg string, args ...KeyValue) (int, error)

func (o Output) Print(ctx context.Context, args ...interface{}) {
	writeError(o(ctx, verbosity.Norm, fmt.Sprint(args...)))
}

func (o Output) PrintKV(ctx context.Context, msg string, kv ...KeyValue) {
	writeError(o(ctx, verbosity.Norm, msg, kv...))
}

func (o Output) Printf(ctx context.Context, format string, args ...interface{}) {
	writeError(o(ctx, verbosity.Norm, fmt.Sprintf(format, args...)))
}

func (o Output) Println(ctx context.Context, args ...interface{}) {
	writeError(o(ctx, verbosity.Norm, fmt.Sprintln(args...)))
}

func (o Output) Info(ctx context.Context, args ...interface{}) {
	writeError(o(ctx, verbosity.Info, fmt.Sprint(args...)))
}

func (o Output) InfoKV(ctx context.Context, msg string, kv ...KeyValue) {
	writeError(o(ctx, verbosity.Info, msg, kv...))
}

func (o Output) Debug(ctx context.Context, args ...interface{}) {
	writeError(o(ctx, verbosity.Debug, fmt.Sprint(args...)))
}

func (o Output) DebugKV(ctx context.Context, msg string, kv ...KeyValue) {
	writeError(o(ctx, verbosity.Debug, msg, kv...))
}

func (o Output) Trace(ctx context.Context, args ...interface{}) {
	writeError(o(ctx, verbosity.Trace, fmt.Sprint(args...)))
}

func (o Output) TraceKV(ctx context.Context, msg string, kv ...KeyValue) {
	writeError(o(ctx, verbosity.Trace, msg, kv...))
}

func (o Output) Write(b []byte) (int, error) {
	return o(context.Background(), verbosity.Norm, string(b))
}

func (o Output) Writer(ctx context.Context, verb verbosity.Verbosity) io.Writer {
	return verbosityWriter{ctx, o, verb}
}

type verbosityWriter struct {
	ctx  context.Context
	out  Output
	verb verbosity.Verbosity
}

func (w verbosityWriter) Write(b []byte) (int, error) {
	return w.out(w.ctx, w.verb, string(b))
}
