package elf

import (
	"github.com/liamg/extrude/pkg/parsers/elf/compiler"
	"github.com/liamg/extrude/pkg/parsers/elf/hardening"
)

func (m *Metadata) analyse() error {

	m.CompilerInfo = compiler.Identify(m.ELF)
	m.Hardening = hardening.IdentifyAttributes(m.ELF)

	m.checkDisclosure()

	return nil
}
