package main

import (
	"fmt"
	"os"

	"github.com/liamg/extrude/internal/app/extrude"
	"github.com/liamg/extrude/pkg/output"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "extrude",
	Short: "Analyse an executable",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		report, err := extrude.Analyse(args[0])
		if err != nil {
			fail("Error: %s\n", err)
		}
		if err := output.Terminal(report); err != nil {
			fail("Failed to write report: %s\n", err)
		}
	},
}

func fail(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}

func main() {
	rootCmd.Execute()
}
