package elf

import (
	"debug/elf"
	"fmt"

	"github.com/liamg/extrude/pkg/report"
)

func (m *Metadata) CreateReport() (report.Report, error) {
	rep := report.New()

	overview := report.NewSection("Overview")

	overview.AddKeyValue("File", m.File.Path)
	overview.AddKeyValue("Format", m.File.Format.String())

	overview.AddKeyValue("Platform", m.ELF.Machine.String())
	class := "32-bit"
	if m.ELF.Class == elf.ELFCLASS64 {
		class = "64-bit"
	}
	overview.AddKeyValue("Class", class)
	overview.AddKeyValue("Type", m.ELF.Type.String())
	overview.AddKeyValue("Compiler Name", m.CompilerInfo.Compiler.String())
	if m.CompilerInfo.Version != "" {
		overview.AddKeyValue("Compiler Version", m.CompilerInfo.Version)
	}
	overview.AddKeyValue("Source Language", m.CompilerInfo.Language.String())

	rep.AddSection(overview)

	security := report.NewSection("Security Features")

	security.AddTest(
		"Position Independent Executable",
		boolToResult(m.Hardening.PositionIndependent),
		`A PIE binary and all of its dependencies are loaded into random locations within virtual memory each time the application is executed. This makes Return Oriented Programming (ROP) attacks much more difficult to execute reliably.`,
	)

	security.AddTest(
		"Read-Only Relocations",
		boolToResult(m.Hardening.ReadOnlyRelocations),
		`Hardens ELF programs against loader memory area overwrites by having the loader mark any areas of the relocation table as read-only for any symbols resolved at load-time ("read-only relocations"). This reduces the area of possible GOT-overwrite-style memory corruption attacks. `,
	)

	security.AddTest(
		"Immediate Binding",
		boolToResult(m.Hardening.ImmediateBinding),
		`The runtime linker will resolve all relocations before starting program execution, meaning memory corruption attacks are much less likely.`,
	)

	fortifiedFuncs := m.Hardening.FortifySourceFunctions
	fortified := fortifiedFuncs.Total == 0 || fortifiedFuncs.Fortified > 0
	fortificationTitle := "n/a"
	if fortifiedFuncs.Total > 0 {
		fortificationTitle = fmt.Sprintf("%d/%d", fortifiedFuncs.Fortified, fortifiedFuncs.Total)
	}

	security.AddTest(
		fmt.Sprintf("Fortified Source Functions (%s)", fortificationTitle),
		boolToResult(fortified),
		`This is a security feature which applies to GLIBC functions vulnerable to buffer overflow attacks. It overrides the use of such functions with a safe variation and is enabled by default on most Linux platforms. If GLIBC functions are used within the binary, this test will fail if none are fortified.`,
	)

	security.AddTest(
		"Stack Canary",
		boolToResult(m.Hardening.StackProtected),
		`A "canary" value is pushed onto the stack immediately after the function return pointer. The canary value is then checked before the function returns; if it has changed, the program will abort. This makes buffer overflow attacks much more difficult to carry out.`,
	)

	security.AddTest(
		"Non-Executable Stack",
		boolToResult(m.Hardening.NonExecutableStackHeader),
		`Preventing the stack from being executable means that malicious code injected onto the stack cannot be run.`,
	)

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
