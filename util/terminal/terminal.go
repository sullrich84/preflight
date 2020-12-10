package terminal

import (
	"fmt"
	"github.com/logrusorgru/aurora/v3"
	"log"
	"strings"
)

func PrettyPrintBool(value bool) string {
	if value {
		return aurora.Green("PASS").Bold().String()
	} else {
		return aurora.Red("DENY").Bold().String()
	}
}

func PrettyPrintResults(method string, results []bool) {
	var prettyResults []string
	for _, result := range results {
		prettyResults = append(prettyResults, PrettyPrintBool(result))
	}
	log.Printf(" %-10s %s", method, strings.Join(prettyResults, " "))
}

func PrettyPrintResultsHeader(origins []string) {
	var headlines []string
	for i := 0; i < len(origins); i++ {
		headlines = append(headlines, fmt.Sprintf("%-4d", i))
	}

	log.Printf("%12s%s", "", strings.Join(headlines, " "))
}

func PrettyPrintResultsFooter(_ []string) {

}
