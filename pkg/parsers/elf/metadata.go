package elf

import (
	"debug/elf"

	"github.com/liamg/extrude/pkg/format"
)

type Metadata struct {
	File struct {
		Path   string
		Name   string
		Format format.Format
	}
	ELF       *elf.File
	Compiler  Compiler
	Hardening Hardening
}
