package terminal

import (
	"fmt"
	"github.com/ahmetalpbalkan/go-cursor"
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

	results := make(map[string]map[string]string, mLen)
	for _, method := range methods {
		results[method] = make(map[string]string, oLen)
		for _, origin := range origins {
			results[method][origin] = "CORS"
		}
	}

	return &PrettyPrinter{
		results:     results,
		tabRendered: false,
	}
}

func (prettyPrinter *PrettyPrinter) PrintHeadline(target string, origins []string, methods []string) {
	fmt.Println()
	fmt.Printf(" Target: %s\n", target)
	fmt.Printf(" Origin: %s\n", strings.Join(origins, ", "))
	fmt.Printf(" Method: %s\n", strings.Join(methods, ", "))
	fmt.Println()
}

func (prettyPrinter *PrettyPrinter) PrintResultTable() {
	if prettyPrinter.tabRendered {
		// Move cursor to beginning of previous table
		fmt.Printf(cursor.MoveUp(len(prettyPrinter.results)))
		fmt.Printf("\r")
	} else {
		prettyPrinter.tabRendered = true
	}

	for method, origins := range prettyPrinter.results {
		fmt.Printf(" %-10s", method)
		for _, result := range origins {
			fmt.Printf(" %-4s", result)
		}
		fmt.Printf("\n")
	}
}
