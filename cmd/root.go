package cmd

import (
	"fmt"
	"github.com/sullrich84/preflight/app/build"
	"github.com/sullrich84/preflight/preflight"
	"github.com/sullrich84/preflight/util/terminal"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	target  string
	origin  []string
	methods []string
	header  []string
)

func init() {
	// Disable timestamps on logs
	log.SetFlags(0)

	defaultHeader := []string{"Content-Type"}

	defaultOrigin := []string{
		"http://localhost:3000",
		"http://localhost:8080",
	}

	defaultMethods := []string{
		http.MethodConnect,
		http.MethodDelete,
		http.MethodGet,
		http.MethodHead,
		http.MethodOptions,
		http.MethodPatch,
		http.MethodPost,
		http.MethodPut,
		http.MethodTrace,
	}

	rootCmd.Flags().StringVarP(&target, "target", "T", "https://api.sandbox.wettkampfdb.de", "Target of the CORS preflight")
	rootCmd.Flags().StringSliceVarP(&origin, "origin", "O", defaultOrigin, "Origin of the CORS preflight")
	rootCmd.Flags().StringSliceVarP(&methods, "methods", "M", defaultMethods, "Methods to check in preflight")
	rootCmd.Flags().StringSliceVarP(&header, "header", "H", defaultHeader, "Populates the Access-Control-Request-Headers for preflight")

	rootCmd.VersionTemplate()
	rootCmd.Version = fmt.Sprintf("%s (%s)", build.Version, build.Commit)
}

var rootCmd = &cobra.Command{
	Use:   "preflight",
	Short: "PreFlight is a CORS testing tool",
	Run:   run,
}

func run(_ *cobra.Command, _ []string) {
	log.Println()
	log.Printf(" Target: %s", target)
	log.Printf(" Origin: %s", strings.Join(origin, ", "))
	log.Printf(" Header: %s", strings.Join(header, ", "))
	log.Println()

	terminal.PrettyPrintResultsHeader(len(origin))
	preflight.Preflight(preflight.Recipe{
		Target:  target,
		Origins: origin,
		Methods: methods,
		Headers: header,
	}, terminal.PrettyPrintResults)
	terminal.PrettyPrintResultsFooter(origin)

	log.Println()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
