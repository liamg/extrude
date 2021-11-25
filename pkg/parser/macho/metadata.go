package macho

import (
	"debug/macho"

	"github.com/liamg/extrude/pkg/format"
	"github.com/liamg/extrude/pkg/parser/macho/hardening"
)

type Metadata struct {
	File struct {
		Path   string
		Name   string
		Format format.Format
	}
	Hardening hardening.Attributes
	thin      *macho.File
	fat       *macho.FatFile
	Notes     []Note
}

type Note struct {
	Heading string
	Content string
}
