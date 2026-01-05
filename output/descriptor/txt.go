package descriptor

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"text/template"

	"gitoa.ru/go-4devs/config"
	"gitoa.ru/go-4devs/config/definition/option"
	"gitoa.ru/go-4devs/config/param"
	"gitoa.ru/go-4devs/config/provider/arg"
	"gitoa.ru/go-4devs/console/output"
)

const (
	defaultSpace  = 2
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
  {{ .Name }} {{ synopsis .Options }}
{{ definition .Options }}
{{- help . }}
		`))

	txtListTemplate = template.Must(template.New("txt_list").
			Funcs(txtFunc).
			Parse(`<comment>Usage:</comment>
  command [options] [arguments]
{{ definition .Options }}
{{- commands .Commands -}}
	`))
)

func TxtStyle() param.Option {
	return arg.WithStyle(
		arg.Style{
			Start: "<comment>",
			End:   "</comment>",
		},
		arg.Style{
			Start: "<info>",
			End:   "</info>",
		},
	)
}

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

	var buf bytes.Buffer

	buf.WriteString("\n<comment>Help:</comment>")
	buf.WriteString(cmd.Help)

	return buf.String()
}

func txtDefinition(options config.Options) string {
	var buf bytes.Buffer

	err := arg.NewDump().Reference(&buf, options)
	if err != nil {
		return err.Error()
	}

	return buf.String()
}

func txtSynopsis(options config.Options) string {
	def := arg.NewViews(options, nil)

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
		buf.WriteString(arg.Name(dashDelimiter))
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
