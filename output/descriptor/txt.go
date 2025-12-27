package descriptor

import (
	"bytes"
	"context"
	"fmt"
	"strconv"
	"strings"
	"text/template"
	"time"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/param"
	"gitoa.ru/go-4devs/config/value"
	"gitoa.ru/go-4devs/console/output"
)

const (
	defaultSpace  = 2
	infoLen       = 13
	dashDelimiter = "-"
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

	txtListTemplate = template.Must(template.New("txt_list").
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

	err := txtHelpTemplate.Execute(&tpl, cmd)
	if err != nil {
		return fmt.Errorf("execute txt help tpl:%w", err)
	}

	out.Println(ctx, tpl.String())

	return nil
}

func (t *txt) Commands(ctx context.Context, out output.Output, cmds Commands) error {
	var buf bytes.Buffer

	err := txtListTemplate.Execute(&buf, cmds)
	if err != nil {
		return fmt.Errorf("execute txt list tpl:%w", err)
	}

	out.Println(ctx, buf.String())

	return nil
}

func txtDefaultArray(val config.Value) string {
	var st any

	err := val.Unmarshal(&st)
	if err != nil {
		return ""
	}

	return fmt.Sprintf("%v", st)
}

//nolint:cyclop
func txtDefault(val config.Value, vr config.Variable) []byte {
	var buf bytes.Buffer

	buf.WriteString("<comment> [default: ")

	dataType := param.Type(vr)
	if option.IsSlice(vr) {
		buf.WriteString(txtDefaultArray(val))
	} else {
		switch dataType.(type) {
		case int:
			buf.WriteString(strconv.Itoa(val.Int()))
		case int64:
			buf.WriteString(strconv.FormatInt(val.Int64(), 10))
		case uint:
			buf.WriteString(strconv.FormatUint(uint64(val.Uint()), 10))
		case uint64:
			buf.WriteString(strconv.FormatUint(val.Uint64(), 10))
		case float64:
			buf.WriteString(strconv.FormatFloat(val.Float64(), 'g', -1, 64))
		case time.Duration:
			buf.WriteString(val.Duration().String())
		case time.Time:
			buf.WriteString(val.Time().Format(time.RFC3339))
		case string:
			buf.WriteString(val.String())
		default:
			buf.WriteString(fmt.Sprint(val.Any()))
		}
	}

	buf.WriteString("]</comment>")

	return buf.Bytes()
}

func txtCommands(cmds []NSCommand) string {
	width := commandsTotalWidth(cmds)
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
			buf.WriteString(strings.Repeat(" ", width-len(cmd.Name)+defaultSpace))
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

func txtDefinitionOption(maxLen int, opts ...config.Variable) string {
	buf := bytes.Buffer{}
	buf.WriteString("\n\n<comment>Options:</comment>\n")

	for _, opt := range opts {
		if option.IsHidden(opt) {
			continue
		}

		var op bytes.Buffer

		op.WriteString("  <info>")

		if short, ok := option.ParamShort(opt); ok {
			op.WriteString("-")
			op.WriteString(short)
			op.WriteString(", ")
		} else {
			op.WriteString("    ")
		}

		op.WriteString("--")
		op.WriteString(strings.Join(opt.Key(), dashDelimiter))

		if !option.IsBool(opt) {
			if !option.IsRequired(opt) {
				op.WriteString("[")
			}

			op.WriteString("=")
			op.WriteString(strings.ToUpper(strings.Join(opt.Key(), dashDelimiter)))

			if !option.IsRequired(opt) {
				op.WriteString("]")
			}
		}

		op.WriteString("</info>")
		buf.Write(op.Bytes())
		buf.WriteString(strings.Repeat(" ", maxLen+17-op.Len()))
		buf.WriteString(option.DataDescription(opt))

		if data, ok := option.DataDefaut(opt); ok {
			buf.Write(txtDefault(value.New(data), opt))
		}

		if option.IsSlice(opt) {
			buf.WriteString("<comment> (multiple values allowed)</comment>")
		}

		buf.WriteString("\n")
	}

	return buf.String()
}

func txtDefinition(def Definition) string {
	width := totalWidth(def)

	var buf bytes.Buffer

	if args := def.Arguments(); len(args) > 0 {
		buf.WriteString("\n\n<comment>Arguments:</comment>\n")

		for _, arg := range args {
			var ab bytes.Buffer

			ab.WriteString("  <info>")
			ab.WriteString(strings.Join(arg.Key(), dashDelimiter))
			ab.WriteString("</info>")
			ab.WriteString(strings.Repeat(" ", width+infoLen+defaultSpace-ab.Len()))

			buf.Write(ab.Bytes())
			buf.WriteString(option.DataDescription(arg))

			if data, ok := option.DataDefaut(arg); ok {
				buf.Write(txtDefault(value.New(data), arg))
			}
		}
	}

	if opts := def.Options(); len(opts) > 0 {
		buf.WriteString(txtDefinitionOption(width, opts...))
	}

	return buf.String()
}

func txtSynopsis(def Definition) string {
	var buf bytes.Buffer

	if len(def.Options()) > 0 {
		buf.WriteString("[options] ")
	}

	args := def.Arguments()

	if buf.Len() > 0 && len(args) > 0 {
		buf.WriteString("[--]")
	}

	var opt int

	for _, arg := range args {
		buf.WriteString(" ")

		if !option.IsRequired(arg) {
			buf.WriteString("[")

			opt++
		}

		buf.WriteString("<")
		buf.WriteString(strings.Join(arg.Key(), dashDelimiter))
		buf.WriteString(">")

		if option.IsSlice(arg) {
			buf.WriteString("...")
		}
	}

	buf.WriteString(strings.Repeat("]", opt))

	return buf.String()
}

func commandsTotalWidth(cmds []NSCommand) int {
	var width int

	for _, ns := range cmds {
		for _, cmd := range ns.Commands {
			if len(cmd.Name) > width {
				width = len(cmd.Name)
			}
		}
	}

	return width
}

//nolint:mnd
func totalWidth(def Definition) int {
	var width int

	for _, arg := range def.Arguments() {
		if l := len(strings.Join(arg.Key(), dashDelimiter)); l > width {
			width = l
		}
	}

	for _, opt := range def.Options() {
		current := len(strings.Join(opt.Key(), dashDelimiter)) + 6

		if !option.IsBool(opt) {
			current = current*2 + 1
		}

		if _, ok := option.DataDefaut(opt); ok {
			current += 2
		}

		if current > width {
			width = current
		}
	}

	return width
}
