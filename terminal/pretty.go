package terminal

import (
	"fmt"
	"github.com/ahmetalpbalkan/go-cursor"
	"github.com/logrusorgru/aurora/v3"
	"sort"
	"strings"
)

type PrettyPrinter struct {
	origins     []string
	methods     []string
	results     map[string]map[string]string
	tabRendered bool
}

func NewPrettyPrinter(origins []string, methods []string) *PrettyPrinter {
	oLen := len(origins)
	mLen := len(methods)

	// Hide the ansi cursor
	fmt.Printf(cursor.Hide())

	// Sort methods
	sort.Strings(methods)

	results := make(map[string]map[string]string, mLen)
	for _, method := range methods {
		results[method] = make(map[string]string, oLen)
		for _, origin := range origins {
			results[method][origin] = renderPending()
		}
	}

	return &PrettyPrinter{
		origins:     origins,
		methods:     methods,
		results:     results,
		tabRendered: false,
	}
}

func (prettyPrinter *PrettyPrinter) PrintHeadline(target string, origins []string, header []string) {
	fmt.Println()
	fmt.Printf(" Target: %s\n", target)
	fmt.Printf(" Origin: %s\n", strings.Join(origins, ", "))
	fmt.Printf(" Header: %s\n", strings.Join(header, ", "))
	fmt.Println()
}

func (prettyPrinter *PrettyPrinter) PrintResultTable() {
	if prettyPrinter.tabRendered {
		// Move cursor to beginning of previous table
		fmt.Printf(cursor.MoveUp(len(prettyPrinter.methods)))
		fmt.Printf("\r")
	} else {
		prettyPrinter.tabRendered = true
	}

	for _, method := range prettyPrinter.methods {
		fmt.Printf(" %-10s", method)
		for _, origin := range prettyPrinter.origins {
			fmt.Printf(" %-4s", prettyPrinter.results[method][origin])
		}
		fmt.Printf("\n")
	}
}

func (prettyPrinter *PrettyPrinter) Update(origin string, method string, succeeded bool) {
	if succeeded {
		prettyPrinter.results[method][origin] = renderPass()
	} else {
		prettyPrinter.results[method][origin] = renderFail()
	}
	prettyPrinter.PrintResultTable()
}

func renderPass() string {
	return aurora.BrightCyan("PASS").Bold().String()
}

func renderFail() string {
	return aurora.Magenta("FAIL").Bold().String()
}

func renderPending() string {
	return aurora.Gray(7, "WAIT").String()
}
