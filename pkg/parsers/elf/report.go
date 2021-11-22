package elf

import "github.com/liamg/extrude/pkg/report"

func (m *Metadata) CreateReport() (report.Report, error) {
	rep := report.New()

	overview := report.NewSection("Overview")

	overview.AddKeyValue("File", m.File.Path)
	overview.AddKeyValue("Format", m.File.Format.String())

	if m.ELF != nil {
		overview.AddKeyValue("Platform", m.ELF.Machine.String())
		overview.AddKeyValue("OS/ABI", m.ELF.OSABI.String())
	}

	overview.AddKeyValue("Compiler Name", m.Compiler.Name)
	overview.AddKeyValue("Compiler Version", m.Compiler.Version)
	overview.AddKeyValue("Source Language", m.Compiler.Language)

	rep.AddSection(overview)

	table := report.NewTable()

	table.AddRow("Fortified Source", boolToResult(m.Hardening.FortifySourceFunctions))
	table.AddRow("Stack Protection", boolToResult(m.Hardening.StackProtected))

	rep.SetResultsTable(table)

	return rep, nil
}

func boolToResult(in bool) string {
	if in {
		return "PASS"
	}
	return "FAIL"
}
