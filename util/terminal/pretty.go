package terminal

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type PrettyPrinter struct {
	templates templates
}

type templates struct {
	headline *template.Template
	window   *template.Template
}

func NewPrettyPrinter() *PrettyPrinter {
	return &PrettyPrinter{
		templates: templates{
			headline: initHeadline(),
			window:   initWindow(),
		},
	}
}

func initHeadline() *template.Template {
	tmpl := template.New("headline")

	tmpl.Funcs(template.FuncMap{"StringsJoin": strings.Join})

	headlineTemplate, headTmplErr := tmpl.Parse(headline)
	if headTmplErr != nil {
		log.Fatal(headTmplErr)
	}

	return headlineTemplate
}

type Headline struct {
	Target string
	Origin []string
	Header []string
}

func (prettyPrinter *PrettyPrinter) PrintHeadline(headline *Headline) {
	err := prettyPrinter.templates.headline.Execute(os.Stdout, headline)
	if err != nil {
		log.Fatal(err)
	}
}

var headline = `
 Target: {{ .Target }}
 Origin: {{ StringsJoin .Origin ", " }}
 Header: {{ StringsJoin .Header ", " }}
`

func initWindow() *template.Template {
	tmpl := template.New("window")

	tmpl.Funcs(template.FuncMap{"PrintOutcome": func(origins []string, method string) string {
		return strings.Join(origins, ", ")
	}})

	windowTemplate, winTmplErr := tmpl.Parse(window)
	if winTmplErr != nil {
		log.Fatal(winTmplErr)
	}

	return windowTemplate
}

type Window struct {
	Origins []string
	Methods []string
}

func (prettyPrinter *PrettyPrinter) PrintWindow(window *Window) {
	err := prettyPrinter.templates.window.Execute(os.Stdout, window)
	if err != nil {
		log.Fatal(err)
	}
}

var window = `
{{ range $mIdx, $m := .Methods }} {{ $m }} {{ PrintOutcome $.Origins $m }}
{{end}}
`
