package extrude

import (
	"github.com/liamg/extrude/pkg/parser"
	"github.com/liamg/extrude/pkg/report"
)

func Analyse(path string) (report.Report, error) {
	return parser.ParseFile(path)
}
