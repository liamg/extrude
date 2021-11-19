package parser

import (
	"io"

	"github.com/liamg/extrude/pkg/format"
	"github.com/liamg/extrude/pkg/parsers/elf"
	"github.com/liamg/extrude/pkg/report"
)

type Parser interface {
	Parse(r io.ReaderAt, path string, format format.Format) (report.Reporter, error)
}

var parsers = make(map[format.Format]Parser)

func init() {
	parsers[format.ELF] = elf.New()
}
