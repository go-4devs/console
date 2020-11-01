package descriptor

import (
	"bytes"
	"context"
	"fmt"
	"strconv"
	"strings"
	"text/template"
	"time"

	"gitoa.ru/go-4devs/console/input"
	"gitoa.ru/go-4devs/console/input/value"
	"gitoa.ru/go-4devs/console/input/value/flag"
	"gitoa.ru/go-4devs/console/output"
)

const (
	defaultSpace = 2
	infoLen      = 13
)

//nolint:gochecknoglobals
var (
	txtFunc = template.FuncMap{
		"synopsis":   txtSynopsis,
		"definition": txtDefinition,
		"help":       txtHelp,
		"commands":   txtCommands,
	}

	txtHelpTemplate = template.Must(template.New("txt_template").
			Funcs(txtFunc).
			Parse(`
{{- if .Description -}}
<comment>Description:</comment>
  {{ .Description }}

{{ end -}}
<comment>Usage:</comment>
  {{ .Name }} {{ synopsis .Definition }}
{{- definition .Definition }}
{{- help . }}
		`))

	txtListTempkate = template.Must(template.New("txt_list").
			Funcs(txtFunc).
			Parse(`<comment>Usage:</comment>
  command [options] [arguments]
{{- definition .Definition }}
{{- commands .Commands -}}
	`))
)

type txt struct{}

func (t *txt) Command(ctx context.Context, out output.Output, cmd Command) error {
	var tpl bytes.Buffer

	if err := txtHelpTemplate.Execute(&tpl, cmd); err != nil {
		return err
	}

	out.Println(ctx, tpl.String())

	return nil
}

func (t *txt) Commands(ctx context.Context, out output.Output, cmds Commands) error {
	var buf bytes.Buffer

	if err := txtListTempkate.Execute(&buf, cmds); err != nil {
		return err
	}

	out.Println(ctx, buf.String())

	return nil
}

func txtDefaultArray(v value.Value, f flag.Flag) string {
	st := v.Strings()

	switch {
	case f.IsInt():
		for _, i := range v.Ints() {
			st = append(st, strconv.Itoa(i))
		}
	case f.IsInt64():
		for _, i := range v.Int64s() {
			st = append(st, strconv.FormatInt(i, 10))
		}
	case f.IsUint():
		for _, u := range v.Uints() {
			st = append(st, strconv.FormatUint(uint64(u), 10))
		}
	case f.IsUint64():
		for _, u := range v.Uint64s() {
			st = append(st, strconv.FormatUint(u, 10))
		}
	case f.IsFloat64():
		for _, f := range v.Float64s() {
			st = append(st, strconv.FormatFloat(f, 'g', -1, 64))
		}
	case f.IsDuration():
		for _, d := range v.Durations() {
			st = append(st, d.String())
		}
	case f.IsTime():
		for _, d := range v.Times() {
			st = append(st, d.Format(time.RFC3339))
		}
	}

	return strings.Join(st, ",")
}

func txtDefault(v value.Value, f flag.Flag) []byte {
	var buf bytes.Buffer

	buf.WriteString("<comment> [default: ")

	switch {
	case f.IsArray():
		buf.WriteString(txtDefaultArray(v, f))
	case f.IsInt():
		buf.WriteString(strconv.Itoa(v.Int()))
	case f.IsInt64():
		buf.WriteString(strconv.FormatInt(v.Int64(), 10))
	case f.IsUint():
		buf.WriteString(strconv.FormatUint(uint64(v.Uint()), 10))
	case f.IsUint64():
		buf.WriteString(strconv.FormatUint(v.Uint64(), 10))
	case f.IsFloat64():
		buf.WriteString(strconv.FormatFloat(v.Float64(), 'g', -1, 64))
	case f.IsDuration():
		buf.WriteString(v.Duration().String())
	case f.IsTime():
		buf.WriteString(v.Time().Format(time.RFC3339))
	case f.IsAny():
		buf.WriteString(fmt.Sprint(v.Any()))
	default:
		buf.WriteString(v.String())
	}

	buf.WriteString("]</comment>")

	return buf.Bytes()
}

func txtCommands(cmds []NSCommand) string {
	max := commandsTotalWidth(cmds)
	showNS := len(cmds) > 1

	var buf bytes.Buffer

	buf.WriteString("\n<comment>Available commands")

	if len(cmds) == 1 && cmds[0].Name != "" {
		buf.WriteString("for the \"")
		buf.WriteString(cmds[0].Name)
		buf.WriteString(`" namespace`)
	}

	buf.WriteString(":</comment>\n")

	for _, ns := range cmds {
		if ns.Name != "" && showNS {
			buf.WriteString("<comment>")
			buf.WriteString(ns.Name)
			buf.WriteString("</comment>\n")
		}

		for _, cmd := range ns.Commands {
			buf.WriteString("  <info>")
			buf.WriteString(cmd.Name)
			buf.WriteString("</info>")
			buf.WriteString(strings.Repeat(" ", max-len(cmd.Name)+defaultSpace))
			buf.WriteString(cmd.Description)
			buf.WriteString("\n")
		}
	}

	return buf.String()
}

func txtHelp(cmd Command) string {
	if cmd.Help == "" {
		return ""
	}

	tpl := template.Must(template.New("help").Parse(cmd.Help))

	var buf bytes.Buffer

	buf.WriteString("\n<comment>Help:</comment>")
	_ = tpl.Execute(&buf, cmd)

	return buf.String()
}

func txtDefinitionOption(maxLen int, def *input.Definition) string {
	buf := bytes.Buffer{}
	opts := def.Options()

	buf.WriteString("\n\n<comment>Options:</comment>\n")

	for _, name := range opts {
		opt, _ := def.Option(name)

		var op bytes.Buffer

		op.WriteString("  <info>")

		if opt.HasShort() {
			op.WriteString("-")
			op.WriteString(opt.Short)
			op.WriteString(", ")
		} else {
			op.WriteString("    ")
		}

		op.WriteString("--")
		op.WriteString(opt.Name)

		if !opt.IsBool() {
			if !opt.IsRequired() {
				op.WriteString("[")
			}

			op.WriteString("=")
			op.WriteString(strings.ToUpper(opt.Name))

			if !opt.IsRequired() {
				op.WriteString("]")
			}
		}

		op.WriteString("</info>")
		buf.Write(op.Bytes())
		buf.WriteString(strings.Repeat(" ", maxLen+17-op.Len()))
		buf.WriteString(opt.Description)

		if opt.HasDefault() {
			buf.Write(txtDefault(opt.Default, opt.Flag))
		}

		if opt.IsArray() {
			buf.WriteString("<comment> (multiple values allowed)</comment>")
		}

		buf.WriteString("\n")
	}

	return buf.String()
}

func txtDefinition(def *input.Definition) string {
	max := totalWidth(def)

	var buf bytes.Buffer

	if args := def.Arguments(); len(args) > 0 {
		buf.WriteString("\n\n<comment>Arguments:</comment>\n")

		for pos := range args {
			var ab bytes.Buffer

			arg, _ := def.Argument(pos)

			ab.WriteString("  <info>")
			ab.WriteString(arg.Name)
			ab.WriteString("</info>")
			ab.WriteString(strings.Repeat(" ", max+infoLen+defaultSpace-ab.Len()))

			buf.Write(ab.Bytes())
			buf.WriteString(arg.Description)

			if arg.HasDefault() {
				buf.Write(txtDefault(arg.Default, arg.Flag))
			}
		}
	}

	if opts := def.Options(); len(opts) > 0 {
		buf.WriteString(txtDefinitionOption(max, def))
	}

	return buf.String()
}

func txtSynopsis(def *input.Definition) string {
	var buf bytes.Buffer

	if len(def.Options()) > 0 {
		buf.WriteString("[options] ")
	}

	if buf.Len() > 0 && len(def.Arguments()) > 0 {
		buf.WriteString("[--]")
	}

	var opt int

	for pos := range def.Arguments() {
		buf.WriteString(" ")

		arg, _ := def.Argument(pos)

		if !arg.IsRequired() {
			buf.WriteString("[")
			opt++
		}

		buf.WriteString("<")
		buf.WriteString(arg.Name)
		buf.WriteString(">")

		if arg.IsArray() {
			buf.WriteString("...")
		}
	}

	buf.WriteString(strings.Repeat("]", opt))

	return buf.String()
}

func commandsTotalWidth(cmds []NSCommand) int {
	var max int

	for _, ns := range cmds {
		for _, cmd := range ns.Commands {
			if len(cmd.Name) > max {
				max = len(cmd.Name)
			}
		}
	}

	return max
}

func totalWidth(def *input.Definition) int {
	var max int

	for pos := range def.Arguments() {
		arg, _ := def.Argument(pos)
		l := len(arg.Name)

		if l > max {
			max = l
		}
	}

	for _, name := range def.Options() {
		opt, _ := def.Option(name)
		l := len(opt.Name) + 6

		if !opt.IsBool() {
			l = l*2 + 1
		}

		if opt.HasDefault() {
			l += 2
		}

		if l > max {
			max = l
		}
	}

	return max
}
