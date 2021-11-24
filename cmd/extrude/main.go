package main

import (
	"fmt"
	"os"

	"github.com/liamg/extrude/pkg/output"
	"github.com/liamg/extrude/pkg/parser"
	"github.com/liamg/extrude/pkg/report"
	"github.com/spf13/cobra"
)

var options output.Options
var failOnWarning bool

var rootCmd = &cobra.Command{
	Use:   "extrude [file]",
	Short: "Analyse an executable",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		rep, err := parser.ParseFile(args[0])
		if err != nil {
			fail("Error: %s\n", err)
		}
		if err := output.Terminal(rep, &options); err != nil {
			fail("Failed to write report: %s\n", err)
		}
		switch rep.Status() {
		case report.Fail:
			os.Exit(1)
		case report.Warning:
			if failOnWarning {
				os.Exit(1)
			}
		}
	},
}

func fail(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format, args...)
	os.Exit(1)
}

func main() {
	rootCmd.Flags().BoolVarP(&options.IncludePassingTests, "all", "a", false, "Show details of all tests, not just those which failed.")
	rootCmd.Flags().BoolVarP(&failOnWarning, "fail-on-warning", "w", false, "Exit with a non-zero status even if only warnings are discovered.")
	rootCmd.Execute()
}
