package macho

import (
	"strings"

	"github.com/liamg/extrude/pkg/report"
)

func (m *Metadata) CreateReport() (report.Report, error) {
	rep := report.New()

	overview := report.NewSection("Overview")

	overview.AddKeyValue("File", m.File.Path)
	overview.AddKeyValue("Format", m.File.Format.String())

	if m.fat != nil {
		overview.AddKeyValue("Universal", "Yes (Fat)")
		var arches []string
		for _, arch := range m.fat.Arches {
			arches = append(arches, arch.Cpu.String())
		}
		overview.AddKeyValue("Architectures", strings.Join(arches, ", "))
	} else {
		overview.AddKeyValue("Univeral", "No (Thin)")
		overview.AddKeyValue("Type", m.thin.Type.String())
		overview.AddKeyValue("Architecture", m.thin.Cpu.String())
	}

	rep.AddSection(overview)

	security := report.NewSection("Security Features")

	security.AddTest(
		"Position Independent Executable",
		boolToResult(m.Hardening.PositionIndependentExecutable),
		`A PIE binary and all of its dependencies are loaded into random locations within virtual memory each time the application is executed. This makes Return Oriented Programming (ROP) attacks much more difficult to execute reliably.`,
	)

	security.AddTest(
		"Stack Canary",
		boolToResult(m.Hardening.StackProtected),
		`A "canary" value is pushed onto the stack immediately after the function return pointer. The canary value is then checked before the function returns; if it has changed, the program will abort. This makes buffer overflow attacks much more difficult to carry out.`,
	)

	security.AddTest(
		"Non-Executable Stack",
		boolToResult(m.Hardening.StackExecutionNotAllowed),
		`Preventing the stack from being executable means that malicious code injected onto the stack cannot be run.`,
	)

	security.AddTest(
		"Non-Executable Heap",
		boolToResult(m.Hardening.HeapExecutionNotAllowed),
		`Preventing the heap from being executable means that malicious code written to the heap cannot be run.`,
	)

	security.AddTest(
		"Automatic Reference Counting",
		boolToResult(m.Hardening.AutomaticReferenceCounting),
		`ARC is a runtime memory safety mechanism which keeps track of objects and frees them once they are no longer referenced.`,
	)

	/* In progress...
	security.AddTest(
		"Encryption",
		boolToResult(m.Hardening.Encrypted),
		``,
	)
	*/

	rep.AddSection(security)

	if len(m.Notes) > 0 {
		notes := report.NewSection("Other Findings")
		for _, note := range m.Notes {
			notes.AddTest(note.Heading, report.Warning, note.Content)
		}
		rep.AddSection(notes)
	}

	return rep, nil
}

func boolToResult(in bool) report.Result {
	if in {
		return report.Pass
	}
	return report.Fail
}
