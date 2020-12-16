package terminal

import (
	"fmt"
	"github.com/ahmetalpbalkan/go-cursor"
	"github.com/logrusorgru/aurora/v3"
	"github.com/sullrich84/preflight/app/build"
	"sort"
	"strings"
)

// pass is the themed output for successful responses
var pass = aurora.BrightCyan("PASS").Bold().String()

// fail is the themed output for failed responses
var fail = aurora.Magenta("FAIL").Bold().String()

// pending is the themed response for ongoing requests
var pending = aurora.Gray(7, "WAIT").String()

type PrettyPrinter struct {
	target      string
	origins     []string
	methods     []string
	header      []string
	results     map[string]map[string]string
	tabRendered bool
}

// NewPrettyPrinter will initialize a new NewPrettyPrinter that will be used to
// prompt application states to the ansi terminal.
func NewPrettyPrinter(target string, origins []string, methods []string, header []string) *PrettyPrinter {
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
			results[method][origin] = pending
		}
	}

	return &PrettyPrinter{
		target:      target,
		origins:     origins,
		methods:     methods,
		header:      header,
		results:     results,
		tabRendered: false,
	}
}

// Render renders the terminal output. Output will show a brief argument
// overview and the result table of all CORS preflights.
func (prettyPrinter *PrettyPrinter) Render() {
	prettyPrinter.printHeadline()
	prettyPrinter.printResultTable()
}

// Summary renders the outcome of the preflights to the terminal.
func (prettyPrinter *PrettyPrinter) Summary() {
	fmt.Println()
}

// printHeadline prints a brief argument overview.
func (prettyPrinter *PrettyPrinter) printHeadline() {
	fmt.Println()
	fmt.Printf(" %s %s\n", aurora.Bold("PreFlight"), build.Version)
	fmt.Println()
	fmt.Printf(" Target: %s\n", prettyPrinter.target)
	fmt.Printf(" Origin: %s\n", strings.Join(prettyPrinter.origins, ", "))
	fmt.Printf(" Header: %s\n", strings.Join(prettyPrinter.header, ", "))
	fmt.Println()
}

// printResultTable prints the result table in pending state.
func (prettyPrinter *PrettyPrinter) printResultTable() {
	for _, method := range prettyPrinter.methods {
		fmt.Printf(" %-10s", method)
		for _, origin := range prettyPrinter.origins {
			fmt.Printf(" %-4s", prettyPrinter.results[method][origin])
		}
		fmt.Println()
	}
}

// resetCursor sets the cursor to the start coordinates of the previously
// printed result table.
func (prettyPrinter *PrettyPrinter) resetCursor() {
	fmt.Printf(cursor.MoveUp(len(prettyPrinter.methods)))
	fmt.Printf("\r")
}

// Update updates the CORS test cell in the result table.
func (prettyPrinter *PrettyPrinter) Update(origin string, method string, success bool) {
	if success {
		prettyPrinter.results[method][origin] = pass
	} else {
		prettyPrinter.results[method][origin] = fail
	}

	prettyPrinter.resetCursor()
	prettyPrinter.printResultTable()
}
