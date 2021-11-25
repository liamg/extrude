package macho

import "github.com/liamg/extrude/pkg/parser/macho/hardening"

func (m *Metadata) analyse() error {

	if m.fat != nil {
		var coreAttr hardening.Attributes
		for _, arch := range m.fat.Arches {
			attr := hardening.IdentifyAttributes(arch.File)
			coreAttr = coreAttr.Merge(attr)
		}
		m.Hardening = coreAttr
	} else {
		m.Hardening = hardening.IdentifyAttributes(m.thin)
	}

	return nil
}
