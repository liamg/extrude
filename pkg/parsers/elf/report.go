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

	if len(m.Notes) > 0 {
		notes := report.NewSection("Other Findings")
		for _, note := range m.Notes {
			notes.AddKeyValue(note.Heading, note.Content)
		}
		rep.AddSection(notes)
	}

	table := report.NewTable()

	table.AddTest(
		"Source Fortification",
		boolToResult(m.Hardening.FortifySourceFunctions),
		`This is a security feature which applies to GLIBC functions vulnerable to buffer overflow attacks. It overrides the use of such functions with a safe variation and is enabled by default on most Linux platforms. If GLIBC functions are used within the binary, this test will fail if none are fortified.`,
	)
	table.AddTest("Stack Protection", boolToResult(m.Hardening.StackProtected), `The basic idea behind stack protection is to push a "canary" (a randomly chosen integer) on the stack just after the function return pointer has been pushed. The canary value is then checked before the function returns; if it has changed, the program will abort. Generally, stack buffer overflow (aka "stack smashing") attacks will have to change the value of the canary as they write beyond the end of the buffer before they can get to the return pointer. Since the value of the canary is unknown to the attacker, it cannot be replaced by the attack. Thus, the stack protection allows the program to abort when that happens rather than return to wherever the attacker wanted it to go.`)

	rep.SetResultsTable(table)

	return rep, nil
}

func boolToResult(in bool) report.Result {
	if in {
		return report.Pass
	}
	return report.Fail
}
