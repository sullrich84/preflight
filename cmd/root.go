package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
}

var rootCmd = &cobra.Command{
	Use:   "preflight",
	Short: "PreFlight is a CORS testing tool",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

// Execute the preflight command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
