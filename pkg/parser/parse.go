package parser

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/liamg/extrude/pkg/format"
	"github.com/liamg/extrude/pkg/report"
)

func ParseFile(path string) (report.Report, error) {

	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	f, err := os.Open(absPath)
	if err != nil {
		return nil, err
	}
	defer func() { _ = f.Close() }()

	fileFormat, err := format.Sniff(f)
	if err != nil {
		return nil, fmt.Errorf("failed to derive file format: %w", err)
	}

	parser, ok := parsers[fileFormat]
	if !ok {
		return nil, fmt.Errorf("parsing not supported for %s files", fileFormat.Long())
	}

	reporter, err := parser.Parse(f, path, fileFormat)
	if err != nil {
		return nil, err
	}

	return reporter.CreateReport()
}
