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

func txtDefaultArray(val input.Value, flag input.Flag) string {
	st := val.Strings()

	switch {
	case flag.IsInt():
		for _, i := range val.Ints() {
			st = append(st, strconv.Itoa(i))
		}
	case flag.IsInt64():
		for _, i := range val.Int64s() {
			st = append(st, strconv.FormatInt(i, 10))
		}
	case flag.IsUint():
		for _, u := range val.Uints() {
			st = append(st, strconv.FormatUint(uint64(u), 10))
		}
	case flag.IsUint64():
		for _, u := range val.Uint64s() {
			st = append(st, strconv.FormatUint(u, 10))
		}
	case flag.IsFloat64():
		for _, f := range val.Float64s() {
			st = append(st, strconv.FormatFloat(f, 'g', -1, 64))
		}
	case flag.IsDuration():
		for _, d := range val.Durations() {
			st = append(st, d.String())
		}
	case flag.IsTime():
		for _, d := range val.Times() {
			st = append(st, d.Format(time.RFC3339))
		}
	}

	return strings.Join(st, ",")
}

func txtDefault(val input.Value, flag input.Flag) []byte {
	var buf bytes.Buffer

	buf.WriteString("<comment> [default: ")

	switch {
	case flag.IsArray():
		buf.WriteString(txtDefaultArray(val, flag))
	case flag.IsInt():
		buf.WriteString(strconv.Itoa(val.Int()))
	case flag.IsInt64():
		buf.WriteString(strconv.FormatInt(val.Int64(), 10))
	case flag.IsUint():
		buf.WriteString(strconv.FormatUint(uint64(val.Uint()), 10))
	case flag.IsUint64():
		buf.WriteString(strconv.FormatUint(val.Uint64(), 10))
	case flag.IsFloat64():
		buf.WriteString(strconv.FormatFloat(val.Float64(), 'g', -1, 64))
	case flag.IsDuration():
		buf.WriteString(val.Duration().String())
	case flag.IsTime():
		buf.WriteString(val.Time().Format(time.RFC3339))
	case flag.IsAny():
		buf.WriteString(fmt.Sprint(val.Any()))
	default:
		buf.WriteString(val.String())
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
