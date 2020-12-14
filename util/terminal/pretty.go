package terminal

import (
	"fmt"
	"log"
	"strings"
)

type PrettyPrinter struct {
	origins []string
	methods []string
	results map[string]map[string]string
}

func NewPrettyPrinter(origins []string, methods []string) *PrettyPrinter {
	oLen := len(origins)
	mLen := len(methods)

	results := make(map[string]map[string]string, mLen)
	for _, method := range methods {
		results[method] = make(map[string]string, oLen)
		for _, origin := range origins {
			results[method][origin] = "FOBA"
		}
	}

	return &PrettyPrinter{origins, methods, results}
}

func (prettyPrinter *PrettyPrinter) PrintHeadline(target string, origins []string, methods []string) {
	log.Println()
	log.Printf(" Target: %s", target)
	log.Printf(" Origin: %s", strings.Join(origins, ", "))
	log.Printf(" Method: %s", strings.Join(methods, ", "))
	log.Println()
}

func (prettyPrinter *PrettyPrinter) PrintWindow() {
	var columnWidth = make(map[string]int, len(prettyPrinter.origins))
	for _, origin := range prettyPrinter.origins {
		columnWidth[origin] = len(origin)
	}

	log.Printf(" %-10s %s", " ", strings.Join(prettyPrinter.origins, " "))

	for _, method := range prettyPrinter.methods {
		var pOrigins []string
		for _, origin := range prettyPrinter.origins {
			res := prettyPrinter.results[method][origin]
			cWidth := columnWidth[origin]
			format := fmt.Sprintf("%%-%ds", cWidth)
			pOrigins = append(pOrigins, fmt.Sprintf(format, res))
		}

		log.Printf(" %-10s %s", method, strings.Join(pOrigins, " "))
	}
}
