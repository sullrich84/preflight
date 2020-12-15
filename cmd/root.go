package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/sullrich84/preflight/app/build"
	"github.com/sullrich84/preflight/preflight"
	"github.com/sullrich84/preflight/terminal"
	"log"
	"net/http"
	"os"
)

var (
	target  string
	origins []string
	methods []string
	headers []string
)

// defaultHeader will be used when no header args has been provided.
var defaultHeader = []string{
	"Content-Type",
}

// defaultOrigin will be used when no origin args has been provided.
var defaultOrigin = []string{
	"http://localhost:3000",
	"http://localhost:8080",
}

// defaultMethods will be used when no method args has been provided.
var defaultMethods = []string{
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

// rootCmd is the default command interpreter, cobra will invoke on startup.
var rootCmd = &cobra.Command{
	Use:   "preflight",
	Short: "PreFlight is a CORS testing tool",
	Run:   run,
}

func init() {
	// Disable timestamps on logs
	log.SetFlags(0)

	rootCmd.Flags().StringVarP(
		&target, "target", "T", "https://api.github.com",
		"Target of the CORS preflight")

	rootCmd.Flags().StringSliceVarP(
		&origins, "origins", "O", defaultOrigin,
		"Origin of the CORS preflight")

	rootCmd.Flags().StringSliceVarP(
		&methods, "methods", "M", defaultMethods,
		"Methods to check in preflight")

	rootCmd.Flags().StringSliceVarP(
		&headers, "headers", "H", defaultHeader,
		"Populates the Access-Control-Request-Headers for preflight")

	rootCmd.VersionTemplate()
	rootCmd.Version = fmt.Sprintf("%s (%s)", build.Version, build.Commit)
}

func run(_ *cobra.Command, _ []string) {
	prettyPrinter := terminal.NewPrettyPrinter(target, origins, methods, headers)
	prettyPrinter.Render()

	for _, origin := range origins {
		for _, method := range methods {
			flight, _ := preflight.NewPreFlight(target, origin, method, headers)
			prettyPrinter.Update(origin, method, flight.PreFly())
		}
	}
}

// Execute will run the default command interpreter.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
