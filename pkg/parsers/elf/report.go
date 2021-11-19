package elf

import "github.com/liamg/extrude/pkg/report"

func (m *Metadata) CreateReport() (report.Report, error) {
	rep := report.New()

	overview := report.NewSection("Overview")

	overview.AddKeyValue("File", m.File.Path)
	overview.AddKeyValue("Format", m.File.Format.String())

	if m.ELF != nil {
		overview.AddKeyValue("Architecture", m.ELF.Machine.String())
		overview.AddKeyValue("OS/ABI", m.ELF.OSABI.String())
	}

	overview.AddKeyValue("Compiler", m.Compiler.Name)
	overview.AddKeyValue("Compiler Version", m.Compiler.Version)

	rep.AddSection(overview)

	return rep, nil
}
