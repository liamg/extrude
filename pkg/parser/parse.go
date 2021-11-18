package parser

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/liamg/extrude/pkg/format"
	"github.com/liamg/extrude/pkg/report"
)

func ParseFile(path string) (report.Report, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	fileFormat, err := format.Sniff(f)
	if err != nil {
		return nil, fmt.Errorf("failed to derive file format: %w", err)
	}

	parseFunc, ok := parsers[fileFormat]
	if !ok {
		return nil, fmt.Errorf("parsing not supported for %s files", fileFormat.Short())
	}

	if _, err := f.Seek(0, 0); err != nil {
		return nil, err
	}

	return parseFunc(f, filepath.Base(path), fileFormat)
}
