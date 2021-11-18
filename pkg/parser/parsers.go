package parser

import (
	"github.com/liamg/extrude/pkg/format"
	"github.com/liamg/extrude/pkg/parsers/elf"
)

var parsers = make(map[format.Format]ParseFunc)

func init() {
	parsers[format.ELF] = elf.Parse
}
