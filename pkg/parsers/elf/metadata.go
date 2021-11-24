package elf

import (
	"debug/elf"

	"github.com/liamg/extrude/pkg/parsers/elf/hardening"

	"github.com/liamg/extrude/pkg/parsers/elf/compiler"

	"github.com/liamg/extrude/pkg/format"
)

type Metadata struct {
	File struct {
		Path   string
		Name   string
		Format format.Format
	}
	ELF          *elf.File
	CompilerInfo compiler.Info
	Hardening    hardening.Attributes
	Notes        []Note
}

type Note struct {
	Heading string
	Content string
}
